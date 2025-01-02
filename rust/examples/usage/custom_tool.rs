use async_trait::async_trait;
use serde_json::{json, Value};
use std::error::Error;
use toolkit::{Schema, Tool};

pub struct CustomTool;

impl CustomTool {
    pub fn new() -> Self {
        Self
    }
}

#[async_trait]
impl Tool for CustomTool {
    fn get_name(&self) -> String {
        "custom_tool".to_string()
    }

    fn get_description(&self) -> String {
        "A custom tool that does something specific".to_string()
    }

    fn get_schema(&self) -> Schema {
        Schema {
            parameters: json!({
                "type": "object",
                "properties": {
                    "input": {
                        "type": "string",
                        "description": "The input to process"
                    }
                },
                "required": ["input"]
            }),
            returns: json!({
                "type": "object",
                "properties": {
                    "result": {
                        "type": "string",
                        "description": "The processing result"
                    }
                }
            }),
        }
    }

    async fn execute(&self, _params: Value) -> Result<Value, Box<dyn Error>> {
        Ok(json!({"result": "success"}))
    }
}
