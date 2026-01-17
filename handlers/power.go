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

// registerPower registers power and root tools.
func (r *Registry) registerPower() {
	cat := config.CategoryPower

	// Pow
	r.addTool(
		mcp.NewTool("pow",
			mcp.WithDescription("Raise x to the power y (x^y)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Base")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Exponent")),
		),
		powHandler,
		cat,
	)

	// Pow10
	r.addTool(
		mcp.NewTool("pow10",
			mcp.WithDescription("10 raised to the power n (10^n)"),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Exponent (integer)")),
		),
		pow10Handler,
		cat,
	)

	// Sqrt
	r.addTool(
		mcp.NewTool("sqrt",
			mcp.WithDescription("Square root of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number (must be non-negative)")),
		),
		sqrtHandler,
		cat,
	)

	// Cbrt
	r.addTool(
		mcp.NewTool("cbrt",
			mcp.WithDescription("Cube root of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		cbrtHandler,
		cat,
	)

	// Exp
	r.addTool(
		mcp.NewTool("exp",
			mcp.WithDescription("e raised to the power x (e^x)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Exponent")),
		),
		expHandler,
		cat,
	)

	// Exp2
	r.addTool(
		mcp.NewTool("exp2",
			mcp.WithDescription("2 raised to the power x (2^x)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Exponent")),
		),
		exp2Handler,
		cat,
	)

	// Expm1
	r.addTool(
		mcp.NewTool("expm1",
			mcp.WithDescription("e^x - 1, accurate for small x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Exponent")),
		),
		expm1Handler,
		cat,
	)

	// Hypot
	r.addTool(
		mcp.NewTool("hypot",
			mcp.WithDescription("sqrt(x^2 + y^2) without overflow"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("First value")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Second value")),
		),
		hypotHandler,
		cat,
	)
}

func powHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Pow(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func pow10Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Pow10(n)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func sqrtHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x < 0 {
		return mcp.NewToolResultError("cannot compute square root of negative number"), nil
	}
	result := math.Sqrt(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func cbrtHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Cbrt(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func expHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Exp(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func exp2Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Exp2(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func expm1Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Expm1(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func hypotHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Hypot(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
