/*
Copyright 2024 The Spice.ai OSS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

use async_openai::types::CreateChatCompletionRequest;
use jsonpath_rust::JsonPath;
use llms::chat::Chat;
use serde_json::json;
use std::{
    str::FromStr,
    sync::{Arc, LazyLock},
};

use crate::{init_tracing, TEST_ARGS};

mod create;

#[derive(Clone)]
pub struct TestCase {
    pub name: &'static str,
    pub req: CreateChatCompletionRequest,

    /// Maps (id, `JSONPath` selector), where the selector is into the [`CreateChatCompletionResponse`].
    /// This is used in snapshot testing to assert certain properties of the response.
    pub json_path: Vec<(&'static str, &'static str)>,
}

/// Creates [`TestCase`] instances from request/response that JSON serialize to
/// [`CreateChatCompletionRequest`] and [`CreateChatCompletionResponse`].
#[macro_export]
macro_rules! test_case {
    ($name:expr, $req:expr, $jsonpaths:expr) => {
        TestCase {
            name: $name,
            req: serde_json::from_value($req)
                .expect(&format!("Failed to parse request in test case '{}'", $name)),
            json_path: $jsonpaths,
        }
    };
}

/// For a given mode name, a function that instantiates the model..
type ModelFn<'a> = (&'a str, Box<dyn Fn() -> Arc<Box<dyn Chat>>>);

/// A given model to test.
type ModelDef<'a> = (&'a str, Arc<Box<dyn Chat>>);
#[allow(clippy::expect_used)]
static TEST_MODELS: LazyLock<Vec<ModelDef>> = LazyLock::new(|| {
    let model_creators: [ModelFn; 4] = [
        (
            "anthropic",
            Box::new(|| create::create_anthropic(None).expect("failed to create anthropic model")),
        ),
        ("openai", Box::new(|| create::create_openai("gpt-4o-mini"))),
        (
            "hf_phi3",
            Box::new(|| {
                create::create_hf("microsoft/Phi-3-mini-4k-instruct")
                    .expect("failed to create 'microsoft/Phi-3-mini-4k-instruct' from HF")
            }),
        ),
        (
            "local_phi3",
            Box::new(|| {
                create::create_local("microsoft/Phi-3-mini-4k-instruct")
                    .expect("failed to create 'microsoft/Phi-3-mini-4k-instruct' from local system")
            }),
        ),
    ];

    model_creators
        .iter()
        .filter_map(|(name, creator)| {
            if TEST_ARGS.skip_model(name) {
                None
            } else {
                Some((*name, creator()))
            }
        })
        .collect()
});

/// A mapping of model names (in [`TEST_MODELS`]) and test names (in [`TEST_CASES`]) to skip.
static TEST_DENY_LIST: LazyLock<Vec<(&'static str, &'static str)>> =
    LazyLock::new(|| vec![("hf_phi3", "tool_use"), ("local_phi3", "tool_use")]);

static TEST_CASES: LazyLock<Vec<TestCase>> = LazyLock::new(|| {
    vec![
        test_case!(
            "basic",
            json!({
                "model": "not_needed",
                "messages": [
                    {
                        "role": "user",
                        "content": "Say Hello"
                    }
                ]
            }),
            vec![
                (
                    "message_keys",
                    "$.choices[*].message['role', 'tool_calls', 'refusal']"
                ),
                (
                    "replied_appropriately",
                    "$.choices[*].message[?(@.content ~= 'Hello')].length()"
                )
            ]
        ),
        test_case!(
            "system_prompt",
            json!({
                "model": "not_needed",
                "messages": [
                    {
                        "role": "system",
                        "content": "Quote back the exact message from the user"
                    },
                    {
                        "role": "user",
                        "content": "pong"
                    }
                ]
            }),
            vec![
                (
                    "assistant_response",
                    "$.choices[*].message[?(@.role == 'assistant' && @.content ~= 'pong')].length()"
                ),
                (
                    "replied_appropriately",
                    "$.choices[*].message[?(@.content ~= 'pong')].length()"
                )
            ]
        ),
        test_case!(
            "tool_use",
            json!({
                "model": "not_needed",
                "messages": [
                    {
                      "role": "user",
                      "content": "What'\''s the weather like in Boston today?"
                    }
                ],
                "tool_choice": {"type": "function", "function": {"name": "get_current_weather"}},
                "tools": [
                  {
                    "type": "function",
                    "function": {
                      "name": "get_current_weather",
                      "description": "Get the current weather in a given location, in Celsius",
                      "parameters": {
                        "type": "object",
                        "properties": {
                          "location": {
                            "type": "string",
                            "description": "The city and state, e.g. San Francisco, CA"
                          },
                          "unit": {
                            "type": "string",
                            "enum": ["celsius", "fahrenheit"]
                          }
                        },
                        "required": ["location"]
                      }
                    }
                  }
                ]
            }),
            vec![
                ("finish_reason", "$.choices[0].finish_reason"),
                (
                    "tool_choice",
                    "$.choices[0].message.tool_calls[0].function.name"
                ),
                (
                    "valid_function_args",
                    "$.choices[0].message.tool_calls[0].function.arguments"
                )
            ]
        ),
    ]
});

#[allow(clippy::expect_used, clippy::expect_fun_call)]
async fn run_single_test(test_name: &str, model_name: &str) -> Result<(), anyhow::Error> {
    let _ = dotenvy::from_filename(".env").expect("failed to load .env file");
    init_tracing(None);

    if TEST_DENY_LIST
        .iter()
        .any(|(m, t)| *m == model_name && *t == test_name)
    {
        return Ok(());
    }

    let test = TEST_CASES
        .iter()
        .find(|t| t.name == test_name)
        .expect("test case not found");

    if TEST_ARGS.skip_model(model_name) {
        tracing::debug!("Skipping test {model_name}/{test_name}");
        return Ok(());
    }

    let (_, model) = TEST_MODELS
        .iter()
        .find(|(name, _)| *name == model_name)
        .unwrap_or_else(|| panic!("model {model_name} not found"));

    tracing::info!("Running test {test_name}/{model_name} with {:?}", test.req);

    let actual_resp = model
        .chat_request(test.req.clone())
        .await
        .unwrap_or_else(|_| panic!("For test {test_name}/{model_name}, chat_request failed"));
    tracing::trace!("Response for {test_name}/{model_name}: {actual_resp:?}");

    let resp_value =
        serde_json::to_value(&actual_resp).expect("failed to serialize response to JSON");

    for (id, json_ptr) in &test.json_path {
        let resp_ptr = JsonPath::from_str(json_ptr)
            .expect("invalid JSONPath selector")
            .find(&resp_value);
        insta::assert_snapshot!(
            format!("{test_name}_{model_name}_{id}"),
            serde_json::to_string_pretty(&resp_ptr).expect("Failed to serialize snapshot")
        );
    }
    Ok(())
}

// Macro to create test module and functions
#[macro_export]
macro_rules! generate_model_tests {
    () => {
        macro_rules! test_model_case {
            ($model_name_expr:expr, $test_case_expr:expr) => {
                paste::paste! {
                    #[tokio::test]
                    async fn [<test_ $model_name_expr _ $test_case_expr>]() {
                        run_single_test(stringify!($test_case_expr), stringify!($model_name_expr)).await
                            .expect("test failed");
                    }
                }
            };
        }

        test_model_case!(anthropic, basic);
        test_model_case!(openai, basic);
        test_model_case!(hf_phi3, basic);
        test_model_case!(local_phi3, basic);

        test_model_case!(anthropic, system_prompt);
        test_model_case!(openai, system_prompt);
        test_model_case!(hf_phi3, system_prompt);
        test_model_case!(local_phi3, system_prompt);

        test_model_case!(anthropic, tool_use);
        test_model_case!(openai, tool_use);
        test_model_case!(hf_phi3, tool_use);
        test_model_case!(local_phi3, tool_use);
    };
}
#[cfg(test)]
mod tests {
    use super::*;
    generate_model_tests!();
}
