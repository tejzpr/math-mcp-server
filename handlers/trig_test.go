// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"math"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSinHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"sin(0)", map[string]any{"x": 0.0}, "0"},
		{"sin(pi/2)", map[string]any{"x": math.Pi / 2}, "1"},
		{"sin(pi)", map[string]any{"x": math.Pi}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := sinHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestCosHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"cos(0)", map[string]any{"x": 0.0}, "1"},
		{"cos(pi)", map[string]any{"x": math.Pi}, "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := cosHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestTanHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"tan(0)", map[string]any{"x": 0.0}},
		{"tan(pi/4)", map[string]any{"x": math.Pi / 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := tanHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestAsinHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"asin(0)", map[string]any{"x": 0.0}, false},
		{"asin(1)", map[string]any{"x": 1.0}, false},
		{"asin(-1)", map[string]any{"x": -1.0}, false},
		{"asin(2)", map[string]any{"x": 2.0}, true},
		{"asin(-2)", map[string]any{"x": -2.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := asinHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}

func TestAcosHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"acos(0)", map[string]any{"x": 0.0}, false},
		{"acos(1)", map[string]any{"x": 1.0}, false},
		{"acos(-1)", map[string]any{"x": -1.0}, false},
		{"acos(2)", map[string]any{"x": 2.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := acosHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}

func TestAtanHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": 1.0})
	result, err := atanHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
}

func TestAtan2Handler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"atan2(1,1)", map[string]any{"y": 1.0, "x": 1.0}},
		{"atan2(0,1)", map[string]any{"y": 0.0, "x": 1.0}},
		{"atan2(1,0)", map[string]any{"y": 1.0, "x": 0.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := atan2Handler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestSincosHandler(t *testing.T) {
	req := makeRequest(map[string]any{"x": math.Pi / 4})
	result, err := sincosHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "sin:")
	assert.Contains(t, text, "cos:")
}
