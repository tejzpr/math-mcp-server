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

// registerTrig registers trigonometric tools.
func (r *Registry) registerTrig() {
	cat := config.CategoryTrig

	// Sin
	r.addTool(
		mcp.NewTool("sin",
			mcp.WithDescription("Sine of x (x in radians)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Angle in radians")),
		),
		sinHandler,
		cat,
	)

	// Cos
	r.addTool(
		mcp.NewTool("cos",
			mcp.WithDescription("Cosine of x (x in radians)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Angle in radians")),
		),
		cosHandler,
		cat,
	)

	// Tan
	r.addTool(
		mcp.NewTool("tan",
			mcp.WithDescription("Tangent of x (x in radians)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Angle in radians")),
		),
		tanHandler,
		cat,
	)

	// Asin
	r.addTool(
		mcp.NewTool("asin",
			mcp.WithDescription("Arc sine (inverse sine) of x, returns radians"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value in range [-1, 1]")),
		),
		asinHandler,
		cat,
	)

	// Acos
	r.addTool(
		mcp.NewTool("acos",
			mcp.WithDescription("Arc cosine (inverse cosine) of x, returns radians"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value in range [-1, 1]")),
		),
		acosHandler,
		cat,
	)

	// Atan
	r.addTool(
		mcp.NewTool("atan",
			mcp.WithDescription("Arc tangent (inverse tangent) of x, returns radians"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		atanHandler,
		cat,
	)

	// Atan2
	r.addTool(
		mcp.NewTool("atan2",
			mcp.WithDescription("Arc tangent of y/x, using signs to determine quadrant"),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Y coordinate")),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("X coordinate")),
		),
		atan2Handler,
		cat,
	)

	// Sincos
	r.addTool(
		mcp.NewTool("sincos",
			mcp.WithDescription("Returns both sine and cosine of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Angle in radians")),
		),
		sincosHandler,
		cat,
	)
}

func sinHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Sin(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func cosHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Cos(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func tanHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Tan(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func asinHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x < -1 || x > 1 {
		return mcp.NewToolResultError("asin: input must be in range [-1, 1]"), nil
	}
	result := math.Asin(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func acosHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x < -1 || x > 1 {
		return mcp.NewToolResultError("acos: input must be in range [-1, 1]"), nil
	}
	result := math.Acos(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func atanHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Atan(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func atan2Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Atan2(y, x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func sincosHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sin, cos := math.Sincos(x)
	return mcp.NewToolResultText(fmt.Sprintf("sin: %g, cos: %g", sin, cos)), nil
}
