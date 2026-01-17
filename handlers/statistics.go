// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/tejzpr/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
)

// registerStatistics registers statistical tools.
func (r *Registry) registerStatistics() {
	cat := config.CategoryStatistics

	// Sum
	r.addTool(
		mcp.NewTool("sum",
			mcp.WithDescription("Sum of all numbers in the array"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		sumHandler,
		cat,
	)

	// Product
	r.addTool(
		mcp.NewTool("product",
			mcp.WithDescription("Product of all numbers in the array"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		productHandler,
		cat,
	)

	// Mean
	r.addTool(
		mcp.NewTool("mean",
			mcp.WithDescription("Arithmetic mean (average) of numbers"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		meanHandler,
		cat,
	)

	// Median
	r.addTool(
		mcp.NewTool("median",
			mcp.WithDescription("Median value of numbers"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		medianHandler,
		cat,
	)

	// Mode
	r.addTool(
		mcp.NewTool("mode",
			mcp.WithDescription("Most frequent value(s) in numbers"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		modeHandler,
		cat,
	)

	// Variance
	r.addTool(
		mcp.NewTool("variance",
			mcp.WithDescription("Population variance of numbers"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		varianceHandler,
		cat,
	)

	// StdDev
	r.addTool(
		mcp.NewTool("std_dev",
			mcp.WithDescription("Population standard deviation of numbers"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		stdDevHandler,
		cat,
	)

	// Range
	r.addTool(
		mcp.NewTool("range_stat",
			mcp.WithDescription("Range (max - min) of numbers"),
			mcp.WithArray("numbers", mcp.Required(), mcp.Description("Array of numbers"), mcp.WithNumberItems()),
		),
		rangeStatHandler,
		cat,
	)
}

func sumHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	return mcp.NewToolResultText(fmt.Sprintf("%g", sum)), nil
}

func productHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}
	product := 1.0
	for _, n := range numbers {
		product *= n
	}
	return mcp.NewToolResultText(fmt.Sprintf("%g", product)), nil
}

func meanHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))
	return mcp.NewToolResultText(fmt.Sprintf("%g", mean)), nil
}

func medianHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}

	// Make a copy to avoid modifying the original
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	var median float64
	n := len(sorted)
	if n%2 == 0 {
		median = (sorted[n/2-1] + sorted[n/2]) / 2
	} else {
		median = sorted[n/2]
	}
	return mcp.NewToolResultText(fmt.Sprintf("%g", median)), nil
}

func modeHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}

	// Count frequencies
	freq := make(map[float64]int)
	maxFreq := 0
	for _, n := range numbers {
		freq[n]++
		if freq[n] > maxFreq {
			maxFreq = freq[n]
		}
	}

	// Find all modes
	var modes []float64
	for n, f := range freq {
		if f == maxFreq {
			modes = append(modes, n)
		}
	}
	sort.Float64s(modes)

	// Format output
	if len(modes) == 1 {
		return mcp.NewToolResultText(fmt.Sprintf("%g", modes[0])), nil
	}

	strs := make([]string, len(modes))
	for i, m := range modes {
		strs[i] = fmt.Sprintf("%g", m)
	}
	return mcp.NewToolResultText(fmt.Sprintf("[%s]", joinStrings(strs, ", "))), nil
}

func varianceHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}

	// Calculate mean
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate variance
	var variance float64
	for _, n := range numbers {
		diff := n - mean
		variance += diff * diff
	}
	variance /= float64(len(numbers))

	return mcp.NewToolResultText(fmt.Sprintf("%g", variance)), nil
}

func stdDevHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}

	// Calculate mean
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate variance
	var variance float64
	for _, n := range numbers {
		diff := n - mean
		variance += diff * diff
	}
	variance /= float64(len(numbers))

	stdDev := math.Sqrt(variance)
	return mcp.NewToolResultText(fmt.Sprintf("%g", stdDev)), nil
}

func rangeStatHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	numbers, err := req.RequireFloatSlice("numbers")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if len(numbers) == 0 {
		return mcp.NewToolResultError("array must not be empty"), nil
	}

	minVal := numbers[0]
	maxVal := numbers[0]
	for _, n := range numbers[1:] {
		if n < minVal {
			minVal = n
		}
		if n > maxVal {
			maxVal = n
		}
	}
	return mcp.NewToolResultText(fmt.Sprintf("%g", maxVal-minVal)), nil
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result += sep + s
	}
	return result
}
