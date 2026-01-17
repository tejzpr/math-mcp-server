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

func TestMaxHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"max(2,3)", map[string]any{"x": 2.0, "y": 3.0}, "3"},
		{"max(5,2)", map[string]any{"x": 5.0, "y": 2.0}, "5"},
		{"max(-1,-2)", map[string]any{"x": -1.0, "y": -2.0}, "-1"},
		{"max(equal)", map[string]any{"x": 5.0, "y": 5.0}, "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := maxHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestMinHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"min(2,3)", map[string]any{"x": 2.0, "y": 3.0}, "2"},
		{"min(5,2)", map[string]any{"x": 5.0, "y": 2.0}, "2"},
		{"min(-1,-2)", map[string]any{"x": -1.0, "y": -2.0}, "-2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := minHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestDimHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"dim(5,3)", map[string]any{"x": 5.0, "y": 3.0}, "2"},
		{"dim(3,5)", map[string]any{"x": 3.0, "y": 5.0}, "0"},
		{"dim(5,5)", map[string]any{"x": 5.0, "y": 5.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := dimHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestCopysignHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"copysign(5,-1)", map[string]any{"x": 5.0, "y": -1.0}, "-5"},
		{"copysign(-5,1)", map[string]any{"x": -5.0, "y": 1.0}, "5"},
		{"copysign(5,1)", map[string]any{"x": 5.0, "y": 1.0}, "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := copysignHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
