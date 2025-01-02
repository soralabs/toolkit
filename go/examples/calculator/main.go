package main

import (
	"log"

	"github.com/soralabs/toolkit"
)

func main() {
	// Create a new toolkit
	tk := toolkit.NewToolkit("math-toolkit", toolkit.WithToolkitDescription("A toolkit for mathematical operations"))

	// Create and register the addition tool
	addition := &AdditionTool{}
	if err := tk.RegisterTool(addition); err != nil {
		log.Fatalf("Failed to register addition tool: %v", err)
	}
}
