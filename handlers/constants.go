// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Mathematical constants
const (
	// Phi is the golden ratio (1 + sqrt(5)) / 2
	Phi = 1.6180339887498948482
	// SqrtPhi is the square root of the golden ratio
	SqrtPhi = 1.2720196495140689643
)

// registerConstantsResource registers the mathematical constants resource.
func registerConstantsResource(s *server.MCPServer) {
	resource := mcp.NewResource(
		"math://constants",
		"Mathematical Constants",
		mcp.WithResourceDescription("Common mathematical constants"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(resource, constantsHandler)
}

func constantsHandler(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	constants := fmt.Sprintf(`{
  "pi": %v,
  "e": %v,
  "phi": %v,
  "sqrt2": %v,
  "sqrtE": %v,
  "sqrtPi": %v,
  "sqrtPhi": %v,
  "ln2": %v,
  "log2E": %v,
  "ln10": %v,
  "log10E": %v,
  "maxFloat64": %v,
  "smallestNonzeroFloat64": %v
}`,
		math.Pi,
		math.E,
		Phi,
		math.Sqrt2,
		math.SqrtE,
		math.SqrtPi,
		SqrtPhi,
		math.Ln2,
		math.Log2E,
		math.Ln10,
		math.Log10E,
		math.MaxFloat64,
		math.SmallestNonzeroFloat64,
	)

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      "math://constants",
			MIMEType: "application/json",
			Text:     constants,
		},
	}, nil
}
