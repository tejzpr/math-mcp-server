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

// registerHyperbolic registers hyperbolic function tools.
func (r *Registry) registerHyperbolic() {
	cat := config.CategoryHyperbolic

	// Sinh
	r.addTool(
		mcp.NewTool("sinh",
			mcp.WithDescription("Hyperbolic sine of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		sinhHandler,
		cat,
	)

	// Cosh
	r.addTool(
		mcp.NewTool("cosh",
			mcp.WithDescription("Hyperbolic cosine of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		coshHandler,
		cat,
	)

	// Tanh
	r.addTool(
		mcp.NewTool("tanh",
			mcp.WithDescription("Hyperbolic tangent of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		tanhHandler,
		cat,
	)

	// Asinh
	r.addTool(
		mcp.NewTool("asinh",
			mcp.WithDescription("Inverse hyperbolic sine of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		asinhHandler,
		cat,
	)

	// Acosh
	r.addTool(
		mcp.NewTool("acosh",
			mcp.WithDescription("Inverse hyperbolic cosine of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value (must be >= 1)")),
		),
		acoshHandler,
		cat,
	)

	// Atanh
	r.addTool(
		mcp.NewTool("atanh",
			mcp.WithDescription("Inverse hyperbolic tangent of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value in range (-1, 1)")),
		),
		atanhHandler,
		cat,
	)
}

func sinhHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Sinh(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func coshHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Cosh(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func tanhHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Tanh(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func asinhHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Asinh(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func acoshHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x < 1 {
		return mcp.NewToolResultError("acosh: input must be >= 1"), nil
	}
	result := math.Acosh(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func atanhHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= -1 || x >= 1 {
		return mcp.NewToolResultError("atanh: input must be in range (-1, 1)"), nil
	}
	result := math.Atanh(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
