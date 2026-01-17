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

func TestPowHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"2^3", map[string]any{"x": 2.0, "y": 3.0}, "8"},
		{"2^0", map[string]any{"x": 2.0, "y": 0.0}, "1"},
		{"2^-1", map[string]any{"x": 2.0, "y": -1.0}, "0.5"},
		{"negative base", map[string]any{"x": -2.0, "y": 2.0}, "4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := powHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestPow10Handler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"10^2", map[string]any{"n": 2.0}, "100"},
		{"10^0", map[string]any{"n": 0.0}, "1"},
		{"10^-1", map[string]any{"n": -1.0}, "0.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := pow10Handler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestSqrtHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"sqrt(4)", map[string]any{"x": 4.0}, "2", false},
		{"sqrt(2)", map[string]any{"x": 2.0}, "1.41421", false},
		{"sqrt(0)", map[string]any{"x": 0.0}, "0", false},
		{"sqrt(negative)", map[string]any{"x": -1.0}, "cannot compute square root of negative", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := sqrtHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestCbrtHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"cbrt(8)", map[string]any{"x": 8.0}, "2"},
		{"cbrt(-8)", map[string]any{"x": -8.0}, "-2"},
		{"cbrt(0)", map[string]any{"x": 0.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := cbrtHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestExpHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"e^0", map[string]any{"x": 0.0}, "1"},
		{"e^1", map[string]any{"x": 1.0}, "2.71828"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := expHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestExp2Handler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"2^0", map[string]any{"x": 0.0}, "1"},
		{"2^3", map[string]any{"x": 3.0}, "8"},
		{"2^-1", map[string]any{"x": -1.0}, "0.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := exp2Handler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestExpm1Handler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 0.0})
	result, err := expm1Handler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "0")
}

func TestHypotHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"3-4-5 triangle", map[string]any{"x": 3.0, "y": 4.0}, "5"},
		{"zeros", map[string]any{"x": 0.0, "y": 0.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := hypotHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
