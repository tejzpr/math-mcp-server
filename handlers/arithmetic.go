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

// registerArithmetic registers basic arithmetic tools.
func (r *Registry) registerArithmetic() {
	cat := config.CategoryArithmetic

	// Add
	r.addTool(
		mcp.NewTool("add",
			mcp.WithDescription("Add two numbers"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
		),
		addHandler,
		cat,
	)

	// Subtract
	r.addTool(
		mcp.NewTool("subtract",
			mcp.WithDescription("Subtract two numbers (a - b)"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
		),
		subtractHandler,
		cat,
	)

	// Multiply
	r.addTool(
		mcp.NewTool("multiply",
			mcp.WithDescription("Multiply two numbers"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First number")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second number")),
		),
		multiplyHandler,
		cat,
	)

	// Divide
	r.addTool(
		mcp.NewTool("divide",
			mcp.WithDescription("Divide two numbers (a / b)"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("Dividend")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Divisor")),
		),
		divideHandler,
		cat,
	)

	// Mod
	r.addTool(
		mcp.NewTool("mod",
			mcp.WithDescription("Floating-point modulo (x mod y)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Dividend")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Divisor")),
		),
		modHandler,
		cat,
	)

	// Remainder
	r.addTool(
		mcp.NewTool("remainder",
			mcp.WithDescription("IEEE 754 floating-point remainder of x/y"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Dividend")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Divisor")),
		),
		remainderHandler,
		cat,
	)

	// Abs
	r.addTool(
		mcp.NewTool("abs",
			mcp.WithDescription("Absolute value of a number"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		absHandler,
		cat,
	)
}

func addHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireFloat("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireFloat("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := a + b
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func subtractHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireFloat("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireFloat("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := a - b
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func multiplyHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireFloat("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireFloat("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := a * b
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func divideHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireFloat("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireFloat("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if b == 0 {
		return mcp.NewToolResultError("division by zero"), nil
	}
	result := a / b
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func modHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if y == 0 {
		return mcp.NewToolResultError("modulo by zero"), nil
	}
	result := math.Mod(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func remainderHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if y == 0 {
		return mcp.NewToolResultError("remainder by zero"), nil
	}
	result := math.Remainder(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func absHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Abs(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
