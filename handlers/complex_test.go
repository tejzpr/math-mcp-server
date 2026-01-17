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

func TestComplexAbsHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"abs(3+4i)", map[string]any{"real": 3.0, "imag": 4.0}, "5"},
		{"abs(0+0i)", map[string]any{"real": 0.0, "imag": 0.0}, "0"},
		{"abs(1+0i)", map[string]any{"real": 1.0, "imag": 0.0}, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := complexAbsHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestComplexPhaseHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"phase(1+0i)", map[string]any{"real": 1.0, "imag": 0.0}},
		{"phase(0+1i)", map[string]any{"real": 0.0, "imag": 1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := complexPhaseHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestComplexConjHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"conj(3+4i)", map[string]any{"real": 3.0, "imag": 4.0}, "3 - 4i"},
		{"conj(3-4i)", map[string]any{"real": 3.0, "imag": -4.0}, "3 + 4i"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := complexConjHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestComplexExpHandler(t *testing.T) {
	req := makeRequest(map[string]any{"real": 0.0, "imag": 0.0})
	result, err := complexExpHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "1")
}

func TestComplexLogHandler(t *testing.T) {
	req := makeRequest(map[string]any{"real": 1.0, "imag": 0.0})
	result, err := complexLogHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
}

func TestComplexSqrtHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"sqrt(4+0i)", map[string]any{"real": 4.0, "imag": 0.0}},
		{"sqrt(-1+0i)", map[string]any{"real": -1.0, "imag": 0.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := complexSqrtHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestComplexPowHandler(t *testing.T) {
	req := makeRequest(map[string]any{
		"x_real": 2.0, "x_imag": 0.0,
		"y_real": 2.0, "y_imag": 0.0,
	})
	result, err := complexPowHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "4")
}

func TestComplexSinHandler(t *testing.T) {
	req := makeRequest(map[string]any{"real": 0.0, "imag": 0.0})
	result, err := complexSinHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
}

func TestComplexCosHandler(t *testing.T) {
	req := makeRequest(map[string]any{"real": 0.0, "imag": 0.0})
	result, err := complexCosHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "1")
}

func TestComplexTanHandler(t *testing.T) {
	req := makeRequest(map[string]any{"real": 0.0, "imag": 0.0})
	result, err := complexTanHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
}

func TestComplexPolarHandler(t *testing.T) {
	req := makeRequest(map[string]any{"real": 3.0, "imag": 4.0})
	result, err := complexPolarHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "r:")
	assert.Contains(t, text, "theta:")
	assert.Contains(t, text, "5") // magnitude should be 5
}

func TestComplexRectHandler(t *testing.T) {
	req := makeRequest(map[string]any{"r": 5.0, "theta": 0.0})
	result, err := complexRectHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "5")
}
