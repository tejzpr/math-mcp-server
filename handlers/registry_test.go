// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"testing"

	"github.com/sagacient/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/server"
	"github.com/stretchr/testify/assert"
)

func TestNewRegistry(t *testing.T) {
	registry := NewRegistry()

	// Should have tools registered
	assert.NotNil(t, registry)
	assert.Greater(t, len(registry.tools), 0)
}

func TestRegistryToolCategories(t *testing.T) {
	registry := NewRegistry()

	// Count tools per category
	categoryCounts := make(map[config.Category]int)
	for _, td := range registry.tools {
		categoryCounts[td.Category]++
	}

	// Verify each category has tools (except constants which is a resource)
	assert.Greater(t, categoryCounts[config.CategoryArithmetic], 0, "arithmetic should have tools")
	assert.Greater(t, categoryCounts[config.CategoryPower], 0, "power should have tools")
	assert.Greater(t, categoryCounts[config.CategoryLogarithm], 0, "logarithm should have tools")
	assert.Greater(t, categoryCounts[config.CategoryTrig], 0, "trig should have tools")
	assert.Greater(t, categoryCounts[config.CategoryHyperbolic], 0, "hyperbolic should have tools")
	assert.Greater(t, categoryCounts[config.CategoryRounding], 0, "rounding should have tools")
	assert.Greater(t, categoryCounts[config.CategoryComparison], 0, "comparison should have tools")
	assert.Greater(t, categoryCounts[config.CategorySpecial], 0, "special should have tools")
	assert.Greater(t, categoryCounts[config.CategoryFloatUtils], 0, "float_utils should have tools")
	assert.Greater(t, categoryCounts[config.CategoryConversion], 0, "conversion should have tools")
	assert.Greater(t, categoryCounts[config.CategoryNumberTheory], 0, "number_theory should have tools")
	assert.Greater(t, categoryCounts[config.CategoryStatistics], 0, "statistics should have tools")
	assert.Greater(t, categoryCounts[config.CategoryBitwise], 0, "bitwise should have tools")
	assert.Greater(t, categoryCounts[config.CategoryComplex], 0, "complex should have tools")
}

func TestRegisterToolsWithAllEnabled(t *testing.T) {
	registry := NewRegistry()

	cfg := &config.Config{
		Categories: make(map[config.Category]bool),
	}
	for _, cat := range config.AllCategories() {
		cfg.Categories[cat] = true
	}

	mcpServer := server.NewMCPServer("test", "1.0.0", server.WithToolCapabilities(true))
	registry.RegisterTools(mcpServer, cfg)

	// Server should have tools registered (we can't easily count them, but no panic is good)
}

func TestRegisterToolsWithOnlyArithmetic(t *testing.T) {
	registry := NewRegistry()

	cfg := &config.Config{
		Categories: map[config.Category]bool{
			config.CategoryArithmetic: true,
		},
	}

	mcpServer := server.NewMCPServer("test", "1.0.0", server.WithToolCapabilities(true))
	registry.RegisterTools(mcpServer, cfg)

	// Should work without error
}

func TestRegisterToolsWithNoCategories(t *testing.T) {
	registry := NewRegistry()

	cfg := &config.Config{
		Categories: make(map[config.Category]bool),
	}

	mcpServer := server.NewMCPServer("test", "1.0.0", server.WithToolCapabilities(true))
	registry.RegisterTools(mcpServer, cfg)

	// Should work without error even with no categories
}

func TestRegisterConstants(t *testing.T) {
	cfg := &config.Config{
		Categories: map[config.Category]bool{
			config.CategoryConstants: true,
		},
	}

	mcpServer := server.NewMCPServer("test", "1.0.0", server.WithResourceCapabilities(true, false))
	RegisterConstants(mcpServer, cfg)

	// Should work without error
}

func TestRegisterConstantsDisabled(t *testing.T) {
	cfg := &config.Config{
		Categories: map[config.Category]bool{},
	}

	mcpServer := server.NewMCPServer("test", "1.0.0", server.WithResourceCapabilities(true, false))
	RegisterConstants(mcpServer, cfg)

	// Should work without error (just not register anything)
}
