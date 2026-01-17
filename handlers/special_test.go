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

func TestGammaHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"gamma(1)", map[string]any{"x": 1.0}, "1"},
		{"gamma(5)", map[string]any{"x": 5.0}, "24"}, // 4!
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := gammaHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestLgammaHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 5.0})
	result, err := lgammaHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "lgamma:")
	assert.Contains(t, text, "sign:")
}

func TestErfHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"erf(0)", map[string]any{"x": 0.0}},
		{"erf(1)", map[string]any{"x": 1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := erfHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestErfcHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 0.0})
	result, err := erfcHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "1")
}

func TestErfinvHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"erfinv(0)", map[string]any{"x": 0.0}, false},
		{"erfinv(0.5)", map[string]any{"x": 0.5}, false},
		{"erfinv(1)", map[string]any{"x": 1.0}, true},
		{"erfinv(-1)", map[string]any{"x": -1.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := erfinvHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}

func TestErfcinvHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"erfcinv(1)", map[string]any{"x": 1.0}, false},
		{"erfcinv(0)", map[string]any{"x": 0.0}, true},
		{"erfcinv(2)", map[string]any{"x": 2.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := erfcinvHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}

func TestJ0Handler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 0.0})
	result, err := j0Handler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "1")
}

func TestJ1Handler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 0.0})
	result, err := j1Handler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "0")
}

func TestY0Handler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"y0(1)", map[string]any{"x": 1.0}, false},
		{"y0(0)", map[string]any{"x": 0.0}, true},
		{"y0(-1)", map[string]any{"x": -1.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := y0Handler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}

func TestY1Handler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"y1(1)", map[string]any{"x": 1.0}, false},
		{"y1(0)", map[string]any{"x": 0.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := y1Handler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}
