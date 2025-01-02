package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/soralabs/toolkit"
)

type Calculation struct {
	A    float64   `json:"a"`
	B    float64   `json:"b"`
	Sum  float64   `json:"sum"`
	Time time.Time `json:"timestamp"`
}

type AdditionTool struct {
	history []Calculation
	mu      sync.Mutex // for thread safety
}

func (t *AdditionTool) GetName() string {
	return "addition"
}

func (t *AdditionTool) GetDescription() string {
	return "A simple addition tool that adds two numbers"
}

func (t *AdditionTool) GetSchema() toolkit.Schema {
	return toolkit.Schema{
		Parameters: json.RawMessage(`{
			"type": "object",
			"required": ["a", "b"],
			"properties": {
				"a": {
					"type": "number",
					"description": "First number to add"
				},
				"b": {
					"type": "number",
					"description": "Second number to add"
				}
			}
		}`),
		Returns: json.RawMessage(`{
			"type": "object",
			"properties": {
				"sum": {
					"type": "number",
					"description": "The sum of a and b"
				},
				"calculation_id": {
					"type": "integer",
					"description": "ID of the calculation in history"
				}
			}
		}`),
	}
}

func (t *AdditionTool) Execute(ctx context.Context, params json.RawMessage) (json.RawMessage, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Parse input parameters
	var input struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}
	if err := json.Unmarshal(params, &input); err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	// Calculate and store result
	calc := Calculation{
		A:    input.A,
		B:    input.B,
		Sum:  input.A + input.B,
		Time: time.Now(),
	}
	t.history = append(t.history, calc)

	// Return result as JSON
	result := struct {
		Sum           float64 `json:"sum"`
		CalculationID int     `json:"calculation_id"`
	}{
		Sum:           calc.Sum,
		CalculationID: len(t.history) - 1,
	}

	response, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	return response, nil
}

func (t *AdditionTool) GetHistory() []Calculation {
	t.mu.Lock()
	defer t.mu.Unlock()
	return append([]Calculation{}, t.history...) // Return a copy to prevent data races
}

func (t *AdditionTool) GetCalculation(id int) (Calculation, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if id < 0 || id >= len(t.history) {
		return Calculation{}, fmt.Errorf("calculation ID %d not found", id)
	}
	return t.history[id], nil
}
