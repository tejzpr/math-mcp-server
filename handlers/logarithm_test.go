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

func TestLogHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"ln(1)", map[string]any{"x": 1.0}, "0", false},
		{"ln(e)", map[string]any{"x": 2.718281828}, "1", false},
		{"ln(0)", map[string]any{"x": 0.0}, "non-positive", true},
		{"ln(negative)", map[string]any{"x": -1.0}, "non-positive", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := logHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestLog10Handler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"log10(1)", map[string]any{"x": 1.0}, "0", false},
		{"log10(10)", map[string]any{"x": 10.0}, "1", false},
		{"log10(100)", map[string]any{"x": 100.0}, "2", false},
		{"log10(0)", map[string]any{"x": 0.0}, "non-positive", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := log10Handler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestLog2Handler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"log2(1)", map[string]any{"x": 1.0}, "0", false},
		{"log2(2)", map[string]any{"x": 2.0}, "1", false},
		{"log2(8)", map[string]any{"x": 8.0}, "3", false},
		{"log2(0)", map[string]any{"x": 0.0}, "non-positive", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := log2Handler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestLog1pHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"log1p(0)", map[string]any{"x": 0.0}, "0", false},
		{"log1p(-1)", map[string]any{"x": -1.0}, "log1p undefined", true},
		{"log1p(-2)", map[string]any{"x": -2.0}, "log1p undefined", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := log1pHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestLogbHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"logb(1)", map[string]any{"x": 1.0}, "0"},
		{"logb(2)", map[string]any{"x": 2.0}, "1"},
		{"logb(4)", map[string]any{"x": 4.0}, "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := logbHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
