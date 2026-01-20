// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"

	"github.com/sagacient/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
)

// registerBitwise registers bitwise operation tools.
func (r *Registry) registerBitwise() {
	cat := config.CategoryBitwise

	// AND
	r.addTool(
		mcp.NewTool("bit_and",
			mcp.WithDescription("Bitwise AND of a and b"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First integer")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second integer")),
		),
		bitAndHandler,
		cat,
	)

	// OR
	r.addTool(
		mcp.NewTool("bit_or",
			mcp.WithDescription("Bitwise OR of a and b"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First integer")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second integer")),
		),
		bitOrHandler,
		cat,
	)

	// XOR
	r.addTool(
		mcp.NewTool("bit_xor",
			mcp.WithDescription("Bitwise XOR of a and b"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First integer")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second integer")),
		),
		bitXorHandler,
		cat,
	)

	// NOT
	r.addTool(
		mcp.NewTool("bit_not",
			mcp.WithDescription("Bitwise NOT of a (complement)"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("Integer")),
		),
		bitNotHandler,
		cat,
	)

	// Left Shift
	r.addTool(
		mcp.NewTool("bit_left_shift",
			mcp.WithDescription("Left shift a by n bits (a << n)"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("Integer to shift")),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Number of bits to shift")),
		),
		bitLeftShiftHandler,
		cat,
	)

	// Right Shift
	r.addTool(
		mcp.NewTool("bit_right_shift",
			mcp.WithDescription("Right shift a by n bits (a >> n)"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("Integer to shift")),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Number of bits to shift")),
		),
		bitRightShiftHandler,
		cat,
	)
}

func bitAndHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireInt("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := a & b
	return mcp.NewToolResultText(fmt.Sprintf("%d (binary: %b)", result, result)), nil
}

func bitOrHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireInt("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := a | b
	return mcp.NewToolResultText(fmt.Sprintf("%d (binary: %b)", result, result)), nil
}

func bitXorHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireInt("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := a ^ b
	return mcp.NewToolResultText(fmt.Sprintf("%d (binary: %b)", result, result)), nil
}

func bitNotHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := ^a
	return mcp.NewToolResultText(fmt.Sprintf("%d", result)), nil
}

func bitLeftShiftHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if n < 0 {
		return mcp.NewToolResultError("shift count must be non-negative"), nil
	}
	result := a << uint(n)
	return mcp.NewToolResultText(fmt.Sprintf("%d (binary: %b)", result, result)), nil
}

func bitRightShiftHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if n < 0 {
		return mcp.NewToolResultError("shift count must be non-negative"), nil
	}
	result := a >> uint(n)
	return mcp.NewToolResultText(fmt.Sprintf("%d (binary: %b)", result, result)), nil
}
