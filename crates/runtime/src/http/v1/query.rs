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
use std::sync::Arc;

use axum::{
    body::Bytes,
    http::StatusCode,
    response::{IntoResponse, Response},
    Extension,
};
use axum_extra::TypedHeader;
use headers_accept::Accept;

use crate::datafusion::DataFusion;

use super::{sql_to_http_response, ArrowFormat};

pub(crate) async fn post(
    Extension(df): Extension<Arc<DataFusion>>,
    accept: Option<TypedHeader<Accept>>,
    body: Bytes,
) -> Response {
    let query = match String::from_utf8(body.to_vec()) {
        Ok(query) => query,
        Err(e) => {
            tracing::debug!("Error reading query: {e}");
            return (StatusCode::BAD_REQUEST, e.to_string()).into_response();
        }
    };

    sql_to_http_response(df, &query, ArrowFormat::from_accept_header(&accept)).await
}
