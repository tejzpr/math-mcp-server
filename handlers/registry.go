// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

// Package handlers provides MCP tool handlers for mathematical operations.
package handlers

import (
	"github.com/sagacient/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ToolDefinition holds a tool and its handler.
type ToolDefinition struct {
	Tool     mcp.Tool
	Handler  server.ToolHandlerFunc
	Category config.Category
}

// Registry holds all tool definitions organized by category.
type Registry struct {
	tools []ToolDefinition
}

// NewRegistry creates a new tool registry with all available tools.
func NewRegistry() *Registry {
	r := &Registry{
		tools: make([]ToolDefinition, 0),
	}

	// Register all categories
	r.registerArithmetic()
	r.registerPower()
	r.registerLogarithm()
	r.registerTrig()
	r.registerHyperbolic()
	r.registerRounding()
	r.registerComparison()
	r.registerSpecial()
	r.registerFloatUtils()
	r.registerConversion()
	r.registerNumberTheory()
	r.registerStatistics()
	r.registerBitwise()
	r.registerComplex()

	return r
}

// RegisterTools registers all enabled tools with the MCP server.
func (r *Registry) RegisterTools(s *server.MCPServer, cfg *config.Config) {
	for _, td := range r.tools {
		if cfg.IsEnabled(td.Category) {
			s.AddTool(td.Tool, td.Handler)
		}
	}
}

// RegisterConstants registers the constants resource if enabled.
func RegisterConstants(s *server.MCPServer, cfg *config.Config) {
	if cfg.IsEnabled(config.CategoryConstants) {
		registerConstantsResource(s)
	}
}

// addTool adds a tool definition to the registry.
func (r *Registry) addTool(tool mcp.Tool, handler server.ToolHandlerFunc, category config.Category) {
	r.tools = append(r.tools, ToolDefinition{
		Tool:     tool,
		Handler:  handler,
		Category: category,
	})
}
