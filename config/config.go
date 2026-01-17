// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

// Package config provides environment-based configuration for the math MCP server.
package config

import (
	"os"
	"strings"
)

// Category represents a category of math operations.
type Category string

// Available categories.
const (
	CategoryArithmetic   Category = "arithmetic"
	CategoryPower        Category = "power"
	CategoryLogarithm    Category = "logarithm"
	CategoryTrig         Category = "trig"
	CategoryHyperbolic   Category = "hyperbolic"
	CategoryRounding     Category = "rounding"
	CategoryComparison   Category = "comparison"
	CategorySpecial      Category = "special"
	CategoryFloatUtils   Category = "float_utils"
	CategoryConversion   Category = "conversion"
	CategoryNumberTheory Category = "number_theory"
	CategoryStatistics   Category = "statistics"
	CategoryBitwise      Category = "bitwise"
	CategoryComplex      Category = "complex"
	CategoryConstants    Category = "constants"
)

// AllCategories returns a slice of all available categories.
func AllCategories() []Category {
	return []Category{
		CategoryArithmetic,
		CategoryPower,
		CategoryLogarithm,
		CategoryTrig,
		CategoryHyperbolic,
		CategoryRounding,
		CategoryComparison,
		CategorySpecial,
		CategoryFloatUtils,
		CategoryConversion,
		CategoryNumberTheory,
		CategoryStatistics,
		CategoryBitwise,
		CategoryComplex,
		CategoryConstants,
	}
}

// Config holds the server configuration.
type Config struct {
	// Categories maps category names to whether they are enabled.
	Categories map[Category]bool
	// Transport specifies the transport type ("stdio" or "http").
	Transport string
}

// LoadConfig loads configuration from environment variables.
// MATH_CATEGORIES: comma-separated list of categories or "all" (default).
// TRANSPORT: "stdio" (default) or "http".
func LoadConfig() *Config {
	cfg := &Config{
		Categories: make(map[Category]bool),
		Transport:  "stdio",
	}

	// Parse transport
	transport := os.Getenv("TRANSPORT")
	if transport == "http" {
		cfg.Transport = "http"
	}

	// Parse categories
	categories := os.Getenv("MATH_CATEGORIES")
	if categories == "" || strings.ToLower(categories) == "all" {
		// Enable all categories
		for _, cat := range AllCategories() {
			cfg.Categories[cat] = true
		}
		return cfg
	}

	// Parse comma-separated list
	cfg.Categories = parseCategories(categories)
	return cfg
}

// parseCategories parses a comma-separated list of category names.
func parseCategories(input string) map[Category]bool {
	result := make(map[Category]bool)
	parts := strings.Split(input, ",")

	validCategories := make(map[Category]bool)
	for _, cat := range AllCategories() {
		validCategories[cat] = true
	}

	for _, part := range parts {
		cat := Category(strings.TrimSpace(strings.ToLower(part)))
		if validCategories[cat] {
			result[cat] = true
		}
	}

	return result
}

// IsEnabled checks if a category is enabled.
func (c *Config) IsEnabled(cat Category) bool {
	return c.Categories[cat]
}

// EnabledCategories returns a slice of enabled categories.
func (c *Config) EnabledCategories() []Category {
	var enabled []Category
	for _, cat := range AllCategories() {
		if c.Categories[cat] {
			enabled = append(enabled, cat)
		}
	}
	return enabled
}
