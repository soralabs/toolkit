package toolkit

import "fmt"

// Toolkit is a collection of tools
type Toolkit struct {
	name        string
	description string
	tools       map[string]Tool // map of tool name to tool
}

func NewToolkit(opts ...ToolkitOption) *Toolkit {
	toolkit := &Toolkit{
		tools: make(map[string]Tool),
	}

	for _, opt := range opts {
		opt(toolkit)
	}

	return toolkit
}

func (t *Toolkit) GetName() string        { return t.name }
func (t *Toolkit) GetDescription() string { return t.description }

func (t *Toolkit) GetTools() []Tool {
	tools := make([]Tool, 0, len(t.tools))
	for _, tool := range t.tools {
		tools = append(tools, tool)
	}
	return tools
}

func (t *Toolkit) RegisterTool(tool Tool) error {
	if _, exists := t.tools[tool.GetName()]; exists {
		return fmt.Errorf("tool %s already registered", tool.GetName())
	}
	t.tools[tool.GetName()] = tool
	return nil
}

func (t *Toolkit) GetTool(name string) (Tool, error) {
	tool, exists := t.tools[name]
	if !exists {
		return nil, fmt.Errorf("tool %s not found", name)
	}
	return tool, nil
}
