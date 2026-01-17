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

// registerSpecial registers special mathematical function tools.
func (r *Registry) registerSpecial() {
	cat := config.CategorySpecial

	// Gamma
	r.addTool(
		mcp.NewTool("gamma",
			mcp.WithDescription("Gamma function of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		gammaHandler,
		cat,
	)

	// Lgamma
	r.addTool(
		mcp.NewTool("lgamma",
			mcp.WithDescription("Natural logarithm of the absolute value of Gamma(x)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		lgammaHandler,
		cat,
	)

	// Erf
	r.addTool(
		mcp.NewTool("erf",
			mcp.WithDescription("Error function of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		erfHandler,
		cat,
	)

	// Erfc
	r.addTool(
		mcp.NewTool("erfc",
			mcp.WithDescription("Complementary error function of x (1 - erf(x))"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		erfcHandler,
		cat,
	)

	// Erfinv
	r.addTool(
		mcp.NewTool("erfinv",
			mcp.WithDescription("Inverse error function of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value in range (-1, 1)")),
		),
		erfinvHandler,
		cat,
	)

	// Erfcinv
	r.addTool(
		mcp.NewTool("erfcinv",
			mcp.WithDescription("Inverse complementary error function of x"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value in range (0, 2)")),
		),
		erfcinvHandler,
		cat,
	)

	// J0 - Bessel function first kind order 0
	r.addTool(
		mcp.NewTool("j0",
			mcp.WithDescription("Bessel function of the first kind, order 0"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		j0Handler,
		cat,
	)

	// J1 - Bessel function first kind order 1
	r.addTool(
		mcp.NewTool("j1",
			mcp.WithDescription("Bessel function of the first kind, order 1"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value")),
		),
		j1Handler,
		cat,
	)

	// Y0 - Bessel function second kind order 0
	r.addTool(
		mcp.NewTool("y0",
			mcp.WithDescription("Bessel function of the second kind, order 0"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value (must be positive)")),
		),
		y0Handler,
		cat,
	)

	// Y1 - Bessel function second kind order 1
	r.addTool(
		mcp.NewTool("y1",
			mcp.WithDescription("Bessel function of the second kind, order 1"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Value (must be positive)")),
		),
		y1Handler,
		cat,
	)
}

func gammaHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Gamma(x)
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return mcp.NewToolResultError("gamma undefined for this input"), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func lgammaHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result, sign := math.Lgamma(x)
	return mcp.NewToolResultText(fmt.Sprintf("lgamma: %g, sign: %d", result, sign)), nil
}

func erfHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Erf(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func erfcHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Erfc(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func erfinvHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= -1 || x >= 1 {
		return mcp.NewToolResultError("erfinv: input must be in range (-1, 1)"), nil
	}
	result := math.Erfinv(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func erfcinvHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= 0 || x >= 2 {
		return mcp.NewToolResultError("erfcinv: input must be in range (0, 2)"), nil
	}
	result := math.Erfcinv(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func j0Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.J0(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func j1Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.J1(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func y0Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= 0 {
		return mcp.NewToolResultError("y0: input must be positive"), nil
	}
	result := math.Y0(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func y1Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if x <= 0 {
		return mcp.NewToolResultError("y1: input must be positive"), nil
	}
	result := math.Y1(x)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
