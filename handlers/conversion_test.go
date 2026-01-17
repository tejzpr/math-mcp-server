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

func TestDegreesToRadiansHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantVal float64
	}{
		{"0 degrees", map[string]any{"degrees": 0.0}, 0},
		{"90 degrees", map[string]any{"degrees": 90.0}, math.Pi / 2},
		{"180 degrees", map[string]any{"degrees": 180.0}, math.Pi},
		{"360 degrees", map[string]any{"degrees": 360.0}, 2 * math.Pi},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := degreesToRadiansHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestRadiansToDegreesHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"0 radians", map[string]any{"radians": 0.0}, "0"},
		{"pi/2 radians", map[string]any{"radians": math.Pi / 2}, "90"},
		{"pi radians", map[string]any{"radians": math.Pi}, "180"},
		{"2pi radians", map[string]any{"radians": 2 * math.Pi}, "360"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := radiansToDegreesHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
