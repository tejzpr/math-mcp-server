// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"
	"math"

	"github.com/tejzpr/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
)

// registerLogarithm registers logarithmic tools.
func (r *Registry) registerLogarithm() {
	cat := config.CategoryLogarithm

	// Log (natural)
	r.addTool(
		mcp.NewTool("log",
			mcp.WithDescription("Natural logarithm (ln) of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number (must be positive)")),
		),
		logHandler,
		cat,
	)

	// Log10
	r.addTool(
		mcp.NewTool("log10",
			mcp.WithDescription("Base-10 logarithm of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number (must be positive)")),
		),
		log10Handler,
		cat,
	)

	// Log2
	r.addTool(
		mcp.NewTool("log2",
			mcp.WithDescription("Base-2 logarithm of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number (must be positive)")),
		),
		log2Handler,
		cat,
	)

	// Log1p
	r.addTool(
		mcp.NewTool("log1p",
			mcp.WithDescription("Natural logarithm of (1 + x), accurate for small x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number (must be > -1)")),
		),
		log1pHandler,
		cat,
	)

	// Logb
	r.addTool(
		mcp.NewTool("logb",
			mcp.WithDescription("Binary exponent of x (unbiased exponent)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		logbHandler,
		cat,
	)
}

func logHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= 0 {
		return mcp.NewToolResultError("logarithm undefined for non-positive numbers"), nil
	}
	result := math.Log(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func log10Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= 0 {
		return mcp.NewToolResultError("logarithm undefined for non-positive numbers"), nil
	}
	result := math.Log10(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func log2Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= 0 {
		return mcp.NewToolResultError("logarithm undefined for non-positive numbers"), nil
	}
	result := math.Log2(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func log1pHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= -1 {
		return mcp.NewToolResultError("log1p undefined for x <= -1"), nil
	}
	result := math.Log1p(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func logbHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Logb(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
