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

// registerRounding registers rounding tools.
func (r *Registry) registerRounding() {
	cat := config.CategoryRounding

	// Ceil
	r.addTool(
		mcp.NewTool("ceil",
			mcp.WithDescription("Round x up to the nearest integer"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		ceilHandler,
		cat,
	)

	// Floor
	r.addTool(
		mcp.NewTool("floor",
			mcp.WithDescription("Round x down to the nearest integer"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		floorHandler,
		cat,
	)

	// Round
	r.addTool(
		mcp.NewTool("round",
			mcp.WithDescription("Round x to the nearest integer (half away from zero)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		roundHandler,
		cat,
	)

	// RoundToEven
	r.addTool(
		mcp.NewTool("round_to_even",
			mcp.WithDescription("Round x to the nearest even integer (banker's rounding)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		roundToEvenHandler,
		cat,
	)

	// Trunc
	r.addTool(
		mcp.NewTool("trunc",
			mcp.WithDescription("Truncate x to integer (round towards zero)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		truncHandler,
		cat,
	)
}

func ceilHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Ceil(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func floorHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Floor(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func roundHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Round(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func roundToEvenHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.RoundToEven(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func truncHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Trunc(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
