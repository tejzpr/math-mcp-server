// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"
	"math/cmplx"

	"github.com/tejzpr/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
)

// registerComplex registers complex number tools.
func (r *Registry) registerComplex() {
	cat := config.CategoryComplex

	// Abs
	r.addTool(
		mcp.NewTool("complex_abs",
			mcp.WithDescription("Absolute value (magnitude) of complex number"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexAbsHandler,
		cat,
	)

	// Phase
	r.addTool(
		mcp.NewTool("complex_phase",
			mcp.WithDescription("Phase (argument) of complex number in radians"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexPhaseHandler,
		cat,
	)

	// Conj
	r.addTool(
		mcp.NewTool("complex_conj",
			mcp.WithDescription("Complex conjugate"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexConjHandler,
		cat,
	)

	// Exp
	r.addTool(
		mcp.NewTool("complex_exp",
			mcp.WithDescription("Complex exponential e^z"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexExpHandler,
		cat,
	)

	// Log
	r.addTool(
		mcp.NewTool("complex_log",
			mcp.WithDescription("Complex natural logarithm"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexLogHandler,
		cat,
	)

	// Sqrt
	r.addTool(
		mcp.NewTool("complex_sqrt",
			mcp.WithDescription("Complex square root"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexSqrtHandler,
		cat,
	)

	// Pow
	r.addTool(
		mcp.NewTool("complex_pow",
			mcp.WithDescription("Complex power x^y"),
			mcp.WithNumber("x_real", mcp.Required(), mcp.Description("Base real part")),
			mcp.WithNumber("x_imag", mcp.Required(), mcp.Description("Base imaginary part")),
			mcp.WithNumber("y_real", mcp.Required(), mcp.Description("Exponent real part")),
			mcp.WithNumber("y_imag", mcp.Required(), mcp.Description("Exponent imaginary part")),
		),
		complexPowHandler,
		cat,
	)

	// Sin
	r.addTool(
		mcp.NewTool("complex_sin",
			mcp.WithDescription("Complex sine"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexSinHandler,
		cat,
	)

	// Cos
	r.addTool(
		mcp.NewTool("complex_cos",
			mcp.WithDescription("Complex cosine"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexCosHandler,
		cat,
	)

	// Tan
	r.addTool(
		mcp.NewTool("complex_tan",
			mcp.WithDescription("Complex tangent"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexTanHandler,
		cat,
	)

	// Polar
	r.addTool(
		mcp.NewTool("complex_polar",
			mcp.WithDescription("Convert complex number to polar form (r, theta)"),
			mcp.WithNumber("real", mcp.Required(), mcp.Description("Real part")),
			mcp.WithNumber("imag", mcp.Required(), mcp.Description("Imaginary part")),
		),
		complexPolarHandler,
		cat,
	)

	// Rect
	r.addTool(
		mcp.NewTool("complex_rect",
			mcp.WithDescription("Convert polar form to complex number"),
			mcp.WithNumber("r", mcp.Required(), mcp.Description("Magnitude")),
			mcp.WithNumber("theta", mcp.Required(), mcp.Description("Phase in radians")),
		),
		complexRectHandler,
		cat,
	)
}

func getComplex(req mcp.CallToolRequest, realKey, imagKey string) (complex128, error) {
	r, err := req.RequireFloat(realKey)
	if err != nil {
		return 0, err
	}
	i, err := req.RequireFloat(imagKey)
	if err != nil {
		return 0, err
	}
	return complex(r, i), nil
}

func formatComplex(c complex128) string {
	r, i := real(c), imag(c)
	if i >= 0 {
		return fmt.Sprintf("%g + %gi", r, i)
	}
	return fmt.Sprintf("%g - %gi", r, -i)
}

func complexAbsHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Abs(z)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func complexPhaseHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Phase(z)
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func complexConjHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Conj(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexExpHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Exp(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexLogHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Log(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexSqrtHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Sqrt(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexPowHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	x, err := getComplex(req, "x_real", "x_imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	y, err := getComplex(req, "y_real", "y_imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Pow(x, y)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexSinHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Sin(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexCosHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Cos(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexTanHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Tan(z)
	return mcp.NewToolResultText(formatComplex(result)), nil
}

func complexPolarHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	z, err := getComplex(req, "real", "imag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	r, theta := cmplx.Polar(z)
	return mcp.NewToolResultText(fmt.Sprintf("r: %g, theta: %g", r, theta)), nil
}

func complexRectHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	r, err := req.RequireFloat("r")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	theta, err := req.RequireFloat("theta")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := cmplx.Rect(r, theta)
	return mcp.NewToolResultText(formatComplex(result)), nil
}
