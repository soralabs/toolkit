package toolkit

// ToolOption defines a function type for configuring a DefaultTool
type ToolOption func(*DefaultTool)

func WithDescription(description string) ToolOption {
	return func(t *DefaultTool) {
		t.description = description
	}
}

func WithSchema(schema Schema) ToolOption {
	return func(t *DefaultTool) {
		t.schema = schema
	}
}
