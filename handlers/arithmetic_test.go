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

func makeRequest(args map[string]any) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: args,
		},
	}
}

func TestAddHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"positive numbers", map[string]any{"a": 2.0, "b": 3.0}, "5", false},
		{"negative numbers", map[string]any{"a": -2.0, "b": -3.0}, "-5", false},
		{"mixed signs", map[string]any{"a": -5.0, "b": 3.0}, "-2", false},
		{"zeros", map[string]any{"a": 0.0, "b": 0.0}, "0", false},
		{"decimals", map[string]any{"a": 1.5, "b": 2.5}, "4", false},
		{"missing a", map[string]any{"b": 2.0}, "", true},
		{"missing b", map[string]any{"a": 2.0}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := addHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
				assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
			}
		})
	}
}

func TestSubtractHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"positive result", map[string]any{"a": 5.0, "b": 3.0}, "2"},
		{"negative result", map[string]any{"a": 3.0, "b": 5.0}, "-2"},
		{"zero result", map[string]any{"a": 5.0, "b": 5.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := subtractHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestMultiplyHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"positive numbers", map[string]any{"a": 3.0, "b": 4.0}, "12"},
		{"negative result", map[string]any{"a": -3.0, "b": 4.0}, "-12"},
		{"zero", map[string]any{"a": 5.0, "b": 0.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := multiplyHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestDivideHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"normal division", map[string]any{"a": 10.0, "b": 2.0}, "5", false},
		{"division with remainder", map[string]any{"a": 7.0, "b": 2.0}, "3.5", false},
		{"division by zero", map[string]any{"a": 5.0, "b": 0.0}, "division by zero", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := divideHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestModHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"positive mod", map[string]any{"x": 7.0, "y": 3.0}, "1", false},
		{"no remainder", map[string]any{"x": 6.0, "y": 3.0}, "0", false},
		{"mod by zero", map[string]any{"x": 5.0, "y": 0.0}, "modulo by zero", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := modHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestRemainderHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"positive", map[string]any{"x": 5.0, "y": 3.0}, "-1", false}, // IEEE remainder
		{"remainder by zero", map[string]any{"x": 5.0, "y": 0.0}, "remainder by zero", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := remainderHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestAbsHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"positive", map[string]any{"x": 5.0}, "5"},
		{"negative", map[string]any{"x": -5.0}, "5"},
		{"zero", map[string]any{"x": 0.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := absHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
