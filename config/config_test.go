// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_AllCategories(t *testing.T) {
	// Clear env
	os.Unsetenv("MATH_CATEGORIES")
	os.Unsetenv("TRANSPORT")

	cfg := LoadConfig()

	assert.Equal(t, "stdio", cfg.Transport)
	assert.Len(t, cfg.Categories, len(AllCategories()))

	for _, cat := range AllCategories() {
		assert.True(t, cfg.IsEnabled(cat), "category %s should be enabled", cat)
	}
}

func TestLoadConfig_AllKeyword(t *testing.T) {
	os.Setenv("MATH_CATEGORIES", "all")
	defer os.Unsetenv("MATH_CATEGORIES")

	cfg := LoadConfig()

	assert.Len(t, cfg.Categories, len(AllCategories()))
}

func TestLoadConfig_AllKeywordCaseInsensitive(t *testing.T) {
	os.Setenv("MATH_CATEGORIES", "ALL")
	defer os.Unsetenv("MATH_CATEGORIES")

	cfg := LoadConfig()

	assert.Len(t, cfg.Categories, len(AllCategories()))
}

func TestLoadConfig_SpecificCategories(t *testing.T) {
	os.Setenv("MATH_CATEGORIES", "arithmetic,trig,statistics")
	defer os.Unsetenv("MATH_CATEGORIES")

	cfg := LoadConfig()

	assert.True(t, cfg.IsEnabled(CategoryArithmetic))
	assert.True(t, cfg.IsEnabled(CategoryTrig))
	assert.True(t, cfg.IsEnabled(CategoryStatistics))
	assert.False(t, cfg.IsEnabled(CategoryPower))
	assert.False(t, cfg.IsEnabled(CategoryComplex))
}

func TestLoadConfig_CategoriesWithSpaces(t *testing.T) {
	os.Setenv("MATH_CATEGORIES", " arithmetic , trig , statistics ")
	defer os.Unsetenv("MATH_CATEGORIES")

	cfg := LoadConfig()

	assert.True(t, cfg.IsEnabled(CategoryArithmetic))
	assert.True(t, cfg.IsEnabled(CategoryTrig))
	assert.True(t, cfg.IsEnabled(CategoryStatistics))
}

func TestLoadConfig_InvalidCategories(t *testing.T) {
	os.Setenv("MATH_CATEGORIES", "arithmetic,invalid,nonexistent")
	defer os.Unsetenv("MATH_CATEGORIES")

	cfg := LoadConfig()

	assert.True(t, cfg.IsEnabled(CategoryArithmetic))
	assert.False(t, cfg.IsEnabled(Category("invalid")))
	assert.Len(t, cfg.EnabledCategories(), 1)
}

func TestLoadConfig_HTTPTransport(t *testing.T) {
	os.Setenv("TRANSPORT", "http")
	defer os.Unsetenv("MATH_TRANSPORT")

	cfg := LoadConfig()

	assert.Equal(t, "http", cfg.Transport)
}

func TestLoadConfig_StdioTransport(t *testing.T) {
	os.Setenv("TRANSPORT", "stdio")
	defer os.Unsetenv("MATH_TRANSPORT")

	cfg := LoadConfig()

	assert.Equal(t, "stdio", cfg.Transport)
}

func TestLoadConfig_InvalidTransportDefaultsToStdio(t *testing.T) {
	os.Setenv("TRANSPORT", "websocket")
	defer os.Unsetenv("MATH_TRANSPORT")

	cfg := LoadConfig()

	assert.Equal(t, "stdio", cfg.Transport)
}

func TestEnabledCategories(t *testing.T) {
	os.Setenv("MATH_CATEGORIES", "arithmetic,power")
	defer os.Unsetenv("MATH_CATEGORIES")

	cfg := LoadConfig()
	enabled := cfg.EnabledCategories()

	require.Len(t, enabled, 2)
	// Categories should be in order as defined in AllCategories
	assert.Equal(t, CategoryArithmetic, enabled[0])
	assert.Equal(t, CategoryPower, enabled[1])
}

func TestAllCategories(t *testing.T) {
	categories := AllCategories()

	assert.Len(t, categories, 15)
	assert.Contains(t, categories, CategoryArithmetic)
	assert.Contains(t, categories, CategoryPower)
	assert.Contains(t, categories, CategoryLogarithm)
	assert.Contains(t, categories, CategoryTrig)
	assert.Contains(t, categories, CategoryHyperbolic)
	assert.Contains(t, categories, CategoryRounding)
	assert.Contains(t, categories, CategoryComparison)
	assert.Contains(t, categories, CategorySpecial)
	assert.Contains(t, categories, CategoryFloatUtils)
	assert.Contains(t, categories, CategoryConversion)
	assert.Contains(t, categories, CategoryNumberTheory)
	assert.Contains(t, categories, CategoryStatistics)
	assert.Contains(t, categories, CategoryBitwise)
	assert.Contains(t, categories, CategoryComplex)
	assert.Contains(t, categories, CategoryConstants)
}
