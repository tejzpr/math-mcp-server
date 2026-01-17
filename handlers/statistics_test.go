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

func TestSumHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"sum [1,2,3,4,5]", map[string]any{"numbers": []any{1.0, 2.0, 3.0, 4.0, 5.0}}, "15", false},
		{"sum [10]", map[string]any{"numbers": []any{10.0}}, "10", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := sumHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestProductHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"product [1,2,3,4]", map[string]any{"numbers": []any{1.0, 2.0, 3.0, 4.0}}, "24", false},
		{"product [5]", map[string]any{"numbers": []any{5.0}}, "5", false},
		{"product with zero", map[string]any{"numbers": []any{5.0, 0.0, 3.0}}, "0", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := productHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestMeanHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"mean [1,2,3,4,5]", map[string]any{"numbers": []any{1.0, 2.0, 3.0, 4.0, 5.0}}, "3", false},
		{"mean [10,20]", map[string]any{"numbers": []any{10.0, 20.0}}, "15", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := meanHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestMedianHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"median odd [1,2,3,4,5]", map[string]any{"numbers": []any{1.0, 2.0, 3.0, 4.0, 5.0}}, "3", false},
		{"median even [1,2,3,4]", map[string]any{"numbers": []any{1.0, 2.0, 3.0, 4.0}}, "2.5", false},
		{"median unsorted", map[string]any{"numbers": []any{5.0, 1.0, 3.0}}, "3", false},
		{"single element", map[string]any{"numbers": []any{42.0}}, "42", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := medianHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestModeHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"single mode", map[string]any{"numbers": []any{1.0, 2.0, 2.0, 3.0}}, "2", false},
		{"multiple modes", map[string]any{"numbers": []any{1.0, 1.0, 2.0, 2.0}}, "[1, 2]", false},
		{"all same", map[string]any{"numbers": []any{5.0, 5.0, 5.0}}, "5", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := modeHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestVarianceHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"variance", map[string]any{"numbers": []any{2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0}}, "4", false},
		{"all same", map[string]any{"numbers": []any{5.0, 5.0, 5.0}}, "0", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := varianceHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestStdDevHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"std_dev", map[string]any{"numbers": []any{2.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 9.0}}, "2", false},
		{"all same", map[string]any{"numbers": []any{5.0, 5.0, 5.0}}, "0", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := stdDevHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestRangeStatHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"range [1,5,10]", map[string]any{"numbers": []any{1.0, 5.0, 10.0}}, "9", false},
		{"range [3,3,3]", map[string]any{"numbers": []any{3.0, 3.0, 3.0}}, "0", false},
		{"empty array", map[string]any{"numbers": []any{}}, "empty", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := rangeStatHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
