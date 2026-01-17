// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSinhHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"sinh(0)", map[string]any{"x": 0.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := sinhHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestCoshHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"cosh(0)", map[string]any{"x": 0.0}, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := coshHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestTanhHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"tanh(0)", map[string]any{"x": 0.0}},
		{"tanh(1)", map[string]any{"x": 1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := tanhHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestAsinhHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
	}{
		{"asinh(0)", map[string]any{"x": 0.0}},
		{"asinh(1)", map[string]any{"x": 1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := asinhHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
		})
	}
}

func TestAcoshHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"acosh(1)", map[string]any{"x": 1.0}, false},
		{"acosh(2)", map[string]any{"x": 2.0}, false},
		{"acosh(0.5)", map[string]any{"x": 0.5}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := acoshHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}

func TestAtanhHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		wantErr bool
	}{
		{"atanh(0)", map[string]any{"x": 0.0}, false},
		{"atanh(0.5)", map[string]any{"x": 0.5}, false},
		{"atanh(1)", map[string]any{"x": 1.0}, true},
		{"atanh(-1)", map[string]any{"x": -1.0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := atanhHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			} else {
				assert.False(t, result.IsError)
			}
		})
	}
}
