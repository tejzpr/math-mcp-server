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

func TestGcdHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"gcd(12,8)", map[string]any{"a": 12.0, "b": 8.0}, "4"},
		{"gcd(17,13)", map[string]any{"a": 17.0, "b": 13.0}, "1"},
		{"gcd(100,25)", map[string]any{"a": 100.0, "b": 25.0}, "25"},
		{"gcd(0,5)", map[string]any{"a": 0.0, "b": 5.0}, "5"},
		{"gcd(-12,8)", map[string]any{"a": -12.0, "b": 8.0}, "4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := gcdHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestLcmHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"lcm(4,6)", map[string]any{"a": 4.0, "b": 6.0}, "12"},
		{"lcm(3,5)", map[string]any{"a": 3.0, "b": 5.0}, "15"},
		{"lcm(0,5)", map[string]any{"a": 0.0, "b": 5.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := lcmHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestFactorialHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"0!", map[string]any{"n": 0.0}, "1", false},
		{"1!", map[string]any{"n": 1.0}, "1", false},
		{"5!", map[string]any{"n": 5.0}, "120", false},
		{"10!", map[string]any{"n": 10.0}, "3628800", false},
		{"negative", map[string]any{"n": -1.0}, "negative", true},
		{"too large", map[string]any{"n": 200.0}, "too large", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := factorialHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestFibonacciHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"fib(0)", map[string]any{"n": 0.0}, "0", false},
		{"fib(1)", map[string]any{"n": 1.0}, "1", false},
		{"fib(10)", map[string]any{"n": 10.0}, "55", false},
		{"fib(20)", map[string]any{"n": 20.0}, "6765", false},
		{"negative", map[string]any{"n": -1.0}, "negative", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := fibonacciHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestIsPrimeHandler(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want string
	}{
		{"is_prime(2)", map[string]any{"n": 2.0}, "true"},
		{"is_prime(3)", map[string]any{"n": 3.0}, "true"},
		{"is_prime(4)", map[string]any{"n": 4.0}, "false"},
		{"is_prime(17)", map[string]any{"n": 17.0}, "true"},
		{"is_prime(1)", map[string]any{"n": 1.0}, "false"},
		{"is_prime(0)", map[string]any{"n": 0.0}, "false"},
		{"is_prime(-5)", map[string]any{"n": -5.0}, "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := isPrimeHandler(context.Background(), req)

			require.NoError(t, err)
			assert.False(t, result.IsError)
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}

func TestPrimeFactorsHandler(t *testing.T) {
	tests := []struct {
		name    string
		args    map[string]any
		want    string
		wantErr bool
	}{
		{"factors(12)", map[string]any{"n": 12.0}, "2 × 2 × 3", false},
		{"factors(17)", map[string]any{"n": 17.0}, "17", false},
		{"factors(100)", map[string]any{"n": 100.0}, "2 × 2 × 5 × 5", false},
		{"factors(0)", map[string]any{"n": 0.0}, "positive", true},
		{"factors(-5)", map[string]any{"n": -5.0}, "positive", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := makeRequest(tt.args)
			result, err := primeFactorsHandler(context.Background(), req)

			require.NoError(t, err)
			if tt.wantErr {
				assert.True(t, result.IsError)
			}
			assert.Contains(t, result.Content[0].(mcp.TextContent).Text, tt.want)
		})
	}
}
