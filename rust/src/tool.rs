use async_trait::async_trait;
use serde_json::Value;

#[derive(Clone, Debug)]
pub struct Schema {
    pub parameters: Value,
    pub returns: Value,
}

#[async_trait]
pub trait Tool: Send + Sync {
    fn get_name(&self) -> String;
    fn get_description(&self) -> String;
    fn get_schema(&self) -> Schema;
    async fn execute(&self, params: Value) -> Result<Value, Box<dyn std::error::Error>>;
}
