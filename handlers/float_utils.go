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

// registerFloatUtils registers floating-point utility tools.
func (r *Registry) registerFloatUtils() {
	cat := config.CategoryFloatUtils

	// Frexp
	r.addTool(
		mcp.NewTool("frexp",
			mcp.WithDescription("Break x into a normalized fraction and integral power of 2"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		frexpHandler,
		cat,
	)

	// Ldexp
	r.addTool(
		mcp.NewTool("ldexp",
			mcp.WithDescription("Compute frac * 2^exp (inverse of frexp)"),
			mcp.WithNumber("frac", mcp.Required(), mcp.Description("Fraction")),
			mcp.WithNumber("exp", mcp.Required(), mcp.Description("Exponent (integer)")),
		),
		ldexpHandler,
		cat,
	)

	// Modf
	r.addTool(
		mcp.NewTool("modf",
			mcp.WithDescription("Split x into integer and fractional parts"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		modfHandler,
		cat,
	)

	// Ilogb
	r.addTool(
		mcp.NewTool("ilogb",
			mcp.WithDescription("Binary exponent of x as an integer"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		ilogbHandler,
		cat,
	)

	// Nextafter
	r.addTool(
		mcp.NewTool("nextafter",
			mcp.WithDescription("Next representable float64 value after x towards y"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Starting value")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Direction value")),
		),
		nextafterHandler,
		cat,
	)

	// FMA
	r.addTool(
		mcp.NewTool("fma",
			mcp.WithDescription("Fused multiply-add: x*y + z with single rounding"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("First multiplicand")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("Second multiplicand")),
			mcp.WithNumber("z", mcp.Required(), mcp.Description("Addend")),
		),
		fmaHandler,
		cat,
	)

	// Signbit
	r.addTool(
		mcp.NewTool("signbit",
			mcp.WithDescription("Check if the sign bit of x is set (true for negative)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		signbitHandler,
		cat,
	)

	// IsNaN
	r.addTool(
		mcp.NewTool("is_nan",
			mcp.WithDescription("Check if x is NaN (Not a Number)"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
		),
		isNaNHandler,
		cat,
	)

	// IsInf
	r.addTool(
		mcp.NewTool("is_inf",
			mcp.WithDescription("Check if x is infinity"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("Number")),
			mcp.WithNumber("sign", mcp.Description("Sign: 1 for +Inf, -1 for -Inf, 0 for either")),
		),
		isInfHandler,
		cat,
	)
}

func frexpHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	frac, exp := math.Frexp(x)
	return mcp.NewToolResultText(fmt.Sprintf("frac: %g, exp: %d", frac, exp)), nil
}

func ldexpHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	frac, err := req.RequireFloat("frac")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	exp, err := req.RequireInt("exp")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Ldexp(frac, exp)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func modfHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	integer, frac := math.Modf(x)
	return mcp.NewToolResultText(fmt.Sprintf("integer: %g, frac: %g", integer, frac)), nil
}

func ilogbHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Ilogb(x)
	return mcp.NewToolResultText(fmt.Sprintf("%d", result)), nil
}

func nextafterHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Nextafter(x, y)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func fmaHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	z, err := req.RequireFloat("z")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.FMA(x, y, z)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func signbitHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.Signbit(x)
	return mcp.NewToolResultText(fmt.Sprintf("%t", result)), nil
}

func isNaNHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := math.IsNaN(x)
	return mcp.NewToolResultText(fmt.Sprintf("%t", result)), nil
}

func isInfHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sign := req.GetInt("sign", 0)
	result := math.IsInf(x, sign)
	return mcp.NewToolResultText(fmt.Sprintf("%t", result)), nil
}
