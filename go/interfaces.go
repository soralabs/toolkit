package toolkit

import (
	"context"
	"encoding/json"
)

type Schema struct {
	Parameters json.RawMessage `json:"parameters"`
	Returns    json.RawMessage `json:"returns"`
}

type Tool interface {
	GetName() string
	GetDescription() string
	GetSchema() Schema
	Execute(ctx context.Context, params json.RawMessage) (json.RawMessage, error)
}
