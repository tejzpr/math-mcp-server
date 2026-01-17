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

func TestCeilHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"ceil(1.1)", map[string]any{"x": 1.1}, "2"},
		{"ceil(1.9)", map[string]any{"x": 1.9}, "2"},
		{"ceil(-1.1)", map[string]any{"x": -1.1}, "-1"},
		{"ceil(2.0)", map[string]any{"x": 2.0}, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := ceilHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestFloorHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"floor(1.1)", map[string]any{"x": 1.1}, "1"},
		{"floor(1.9)", map[string]any{"x": 1.9}, "1"},
		{"floor(-1.1)", map[string]any{"x": -1.1}, "-2"},
		{"floor(2.0)", map[string]any{"x": 2.0}, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := floorHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestRoundHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"round(1.4)", map[string]any{"x": 1.4}, "1"},
		{"round(1.5)", map[string]any{"x": 1.5}, "2"},
		{"round(1.6)", map[string]any{"x": 1.6}, "2"},
		{"round(-1.5)", map[string]any{"x": -1.5}, "-2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := roundHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestRoundToEvenHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"round_to_even(0.5)", map[string]any{"x": 0.5}, "0"},
		{"round_to_even(1.5)", map[string]any{"x": 1.5}, "2"},
		{"round_to_even(2.5)", map[string]any{"x": 2.5}, "2"},
		{"round_to_even(3.5)", map[string]any{"x": 3.5}, "4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := roundToEvenHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestTruncHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"trunc(1.9)", map[string]any{"x": 1.9}, "1"},
		{"trunc(-1.9)", map[string]any{"x": -1.9}, "-1"},
		{"trunc(2.0)", map[string]any{"x": 2.0}, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := truncHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
