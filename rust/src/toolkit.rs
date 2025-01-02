use crate::Tool;
use std::collections::HashMap;
use std::sync::Arc;

#[derive(Clone)]
pub struct Toolkit {
    name: String,
    description: String,
    tools: HashMap<String, Arc<dyn Tool>>,
}

impl Toolkit {
    pub fn builder(name: String) -> ToolkitBuilder {
        ToolkitBuilder {
            name,
            description: None,
            tools: HashMap::new(),
        }
    }

    pub fn get_name(&self) -> String {
        self.name.clone()
    }

    pub fn get_description(&self) -> String {
        self.description.clone()
    }

    pub fn get_tools(&self) -> Vec<Arc<dyn Tool>> {
        self.tools.values().cloned().collect()
    }

    pub fn register_tool(&mut self, tool: Arc<dyn Tool>) -> Result<(), String> {
        if self.tools.contains_key(&tool.get_name()) {
            return Err(format!("tool {} already registered", tool.get_name()));
        }
        self.tools.insert(tool.get_name(), tool);
        Ok(())
    }

    pub fn get_tool(&self, name: &str) -> Option<Arc<dyn Tool>> {
        self.tools.get(name).cloned()
    }
}

pub struct ToolkitBuilder {
    name: String,
    description: Option<String>,
    tools: HashMap<String, Arc<dyn Tool>>,
}

impl ToolkitBuilder {
    pub fn description(mut self, description: String) -> Self {
        self.description = Some(description);
        self
    }

    pub fn add_tool(mut self, tool: Arc<dyn Tool>) -> Self {
        self.tools.insert(tool.get_name(), tool);
        self
    }

    pub fn build(self) -> Toolkit {
        Toolkit {
            name: self.name,
            description: self.description.unwrap_or_default(),
            tools: self.tools,
        }
    }
}
