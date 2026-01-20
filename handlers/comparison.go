// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"
	"math"

	"github.com/sagacient/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
)

// registerComparison registers comparison and selection tools.
func (r *Registry) registerComparison() {
	cat := config.CategoryComparison

	// Max
	r.addTool(
		mcp.NewTool("max",
			mcp.WithDescription("Return the larger of x and y"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("First number")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Second number")),
		),
		maxHandler,
		cat,
	)

	// Min
	r.addTool(
		mcp.NewTool("min",
			mcp.WithDescription("Return the smaller of x and y"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("First number")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Second number")),
		),
		minHandler,
		cat,
	)

	// Dim
	r.addTool(
		mcp.NewTool("dim",
			mcp.WithDescription("Return max(x - y, 0)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("First number")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Second number")),
		),
		dimHandler,
		cat,
	)

	// Copysign
	r.addTool(
		mcp.NewTool("copysign",
			mcp.WithDescription("Return x with the sign of y"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Magnitude")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Sign source")),
		),
		copysignHandler,
		cat,
	)
}

func maxHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Max(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func minHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Min(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func dimHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Dim(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func copysignHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Copysign(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
