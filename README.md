# Toolkit

<div align="center">
  <img src="./img/sora_readme_banner.png" alt="Sora Banner" width="100%" />
</div>

A dual-language package (Go/Rust) for building and managing LLM function calling tools and toolkits. Built specifically for Zen.

## Features

- Abstract Tool interface for implementing LLM-compatible functions
- Default tool implementation with required metadata
- Toolkit management for grouping related tools
- Functional options pattern for configuration (Go)
- Builder pattern for configuration (Rust)
- JSON schema support for function parameters and returns

## Installation

### Go
```bash
go get github.com/soralabs/toolkit/go
```

### Rust
```toml
[dependencies]
toolkit = { git = "https://github.com/soralabs/toolkit", subdirectory = "rust" }
```

## Usage

Both implementations provide similar functionality with language-specific idioms:

- Go uses the functional options pattern for configuration
- Rust uses the builder pattern for configuration
- Both support async execution of tools
- Both use JSON schemas for parameter and return type definitions

For examples, see the language-specific directories:
- Go examples in [go/examples](go/examples)
- Rust examples in [rust/examples](rust/examples)
