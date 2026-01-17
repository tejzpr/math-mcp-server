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

func TestBitAndHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"5 & 3", map[string]any{"a": 5.0, "b": 3.0}, "1"},
		{"15 & 8", map[string]any{"a": 15.0, "b": 8.0}, "8"},
		{"0 & 5", map[string]any{"a": 0.0, "b": 5.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := bitAndHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestBitOrHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"5 | 3", map[string]any{"a": 5.0, "b": 3.0}, "7"},
		{"8 | 4", map[string]any{"a": 8.0, "b": 4.0}, "12"},
		{"0 | 5", map[string]any{"a": 0.0, "b": 5.0}, "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := bitOrHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestBitXorHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"5 ^ 3", map[string]any{"a": 5.0, "b": 3.0}, "6"},
		{"15 ^ 15", map[string]any{"a": 15.0, "b": 15.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := bitXorHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestBitNotHandler(t *testing.T) {
	req := makeRequest(map[string]any{"a": 0.0})
	result, err := bitNotHandler(context.Background(), req)

	require.NoError(t, err)
	assert.False(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "-1")
}

func TestBitLeftShiftHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"1 << 3", map[string]any{"a": 1.0, "n": 3.0}, "8", false},
		{"5 << 1", map[string]any{"a": 5.0, "n": 1.0}, "10", false},
		{"negative shift", map[string]any{"a": 5.0, "n": -1.0}, "non-negative", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := bitLeftShiftHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestBitRightShiftHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"8 >> 3", map[string]any{"a": 8.0, "n": 3.0}, "1", false},
		{"10 >> 1", map[string]any{"a": 10.0, "n": 1.0}, "5", false},
		{"negative shift", map[string]any{"a": 5.0, "n": -1.0}, "non-negative", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := bitRightShiftHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
