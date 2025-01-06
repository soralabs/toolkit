mod custom_tool;
use custom_tool::CustomTool;
use serde_json::json;
use std::sync::Arc;
use toolkit::toolkit::Toolkit;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let custom_tool = CustomTool::new();

    let toolkit = Toolkit::builder()
        .name("my-toolkit".to_string())
        .description("My toolkit".to_string())
        .add_tool(Arc::new(custom_tool))
        .build();

    let tool = toolkit.get_tool("custom_tool").unwrap();
    let result = tool
        .execute(json!({
            "input": "test input"
        }))
        .await?;
    println!("Result: {}", result);

    Ok(())
}
