// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFrexpHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 8.0})
	result, err := frexpHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "frac:")
	assert.Contains(t, text, "exp:")
}

func TestLdexpHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"ldexp(0.5, 4)", map[string]any{"frac": 0.5, "exp": 4.0}, "8"},
		{"ldexp(1, 0)", map[string]any{"frac": 1.0, "exp": 0.0}, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := ldexpHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestModfHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 3.14})
	result, err := modfHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "integer:")
	assert.Contains(t, text, "frac:")
}

func TestIlogbHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"ilogb(8)", map[string]any{"x": 8.0}, "3"},
		{"ilogb(1)", map[string]any{"x": 1.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := ilogbHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestNextafterHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 1.0, "y": 2.0})
	result, err := nextafterHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
}

func TestFmaHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"fma(2,3,4)", map[string]any{"x": 2.0, "y": 3.0, "z": 4.0}, "10"},
		{"fma(0,5,3)", map[string]any{"x": 0.0, "y": 5.0, "z": 3.0}, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := fmaHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestSignbitHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"signbit(5)", map[string]any{"x": 5.0}, "false"},
		{"signbit(-5)", map[string]any{"x": -5.0}, "true"},
		{"signbit(0)", map[string]any{"x": 0.0}, "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := signbitHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestIsNaNHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 5.0})
	result, err := isNaNHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "false")
}

func TestIsInfHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 5.0})
	result, err := isInfHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "false")
}
