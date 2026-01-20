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

// registerConversion registers angle conversion tools.
func (r *Registry) registerConversion() {
	cat := config.CategoryConversion

	// Degrees to Radians
	r.addTool(
		mcp.NewTool("degrees_to_radians",
			mcp.WithDescription("Convert degrees to radians"),
			mcp.WithNumber("degrees", mcp.Required(), mcp.Description("Angle in degrees")),
		),
		degreesToRadiansHandler,
		cat,
	)

	// Radians to Degrees
	r.addTool(
		mcp.NewTool("radians_to_degrees",
			mcp.WithDescription("Convert radians to degrees"),
			mcp.WithNumber("radians", mcp.Required(), mcp.Description("Angle in radians")),
		),
		radiansToDegreesHandler,
		cat,
	)
}

func degreesToRadiansHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	degrees, err := req.RequireFloat("degrees")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := degrees * math.Pi / 180
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}

func radiansToDegreesHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	radians, err := req.RequireFloat("radians")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := radians * 180 / math.Pi
	return mcp.NewToolResultText(fmt.Sprintf("%g", result)), nil
}
