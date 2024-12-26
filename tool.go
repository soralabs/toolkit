package toolkit

import (
	"context"
	"encoding/json"
	"fmt"
)

// DefaultTool provides a default implementation of the Tool interface
type DefaultTool struct {
	name        string
	description string
	schema      Schema
}

func NewTool(name string, opts ...ToolOption) (*DefaultTool, error) {
	tool := &DefaultTool{
		name: name,
	}

	for _, opt := range opts {
		opt(tool)
	}

	// Validate all required fields are set
	if tool.name == "" {
		return nil, fmt.Errorf("tool name is required")
	}
	if tool.description == "" {
		return nil, fmt.Errorf("tool description is required")
	}
	if tool.schema.Parameters == nil {
		return nil, fmt.Errorf("tool parameters schema is required")
	}

	return tool, nil
}

func (t *DefaultTool) GetName() string {
	return t.name
}

func (t *DefaultTool) GetDescription() string {
	return t.description
}

func (t *DefaultTool) GetSchema() Schema {
	return t.schema
}

// Execute must be implemented by concrete tools
func (t *DefaultTool) Execute(ctx context.Context, params json.RawMessage) (json.RawMessage, error) {
	return nil, fmt.Errorf("Execute not implemented for tool %s", t.name)
}
