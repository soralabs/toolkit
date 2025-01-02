package toolkit

// ToolkitOption defines a function type for configuring a Toolkit
type ToolkitOption func(*Toolkit)

func WithToolkitName(name string) ToolkitOption {
	return func(t *Toolkit) {
		t.name = name
	}
}

func WithToolkitDescription(description string) ToolkitOption {
	return func(t *Toolkit) {
		t.description = description
	}
}

func WithTools(tools ...Tool) ToolkitOption {
	return func(t *Toolkit) {
		for _, tool := range tools {
			t.tools[tool.GetName()] = tool
		}
	}
}
