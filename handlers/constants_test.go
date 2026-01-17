// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstantsHandler(t *testing.T) {
	req := mcp.ReadResourceRequest{}
	result, err := constantsHandler(context.Background(), req)

	require.NoError(t, err)
	require.Len(t, result, 1)

	contents := result[0].(mcp.TextResourceContents)
	assert.Equal(t, "math://constants", contents.URI)
	assert.Equal(t, "application/json", contents.MIMEType)

	// Check that JSON contains expected constants
	assert.Contains(t, contents.Text, "\"pi\"")
	assert.Contains(t, contents.Text, "\"e\"")
	assert.Contains(t, contents.Text, "\"phi\"")
	assert.Contains(t, contents.Text, "\"sqrt2\"")
	assert.Contains(t, contents.Text, "\"ln2\"")
	assert.Contains(t, contents.Text, "\"log2E\"")
	assert.Contains(t, contents.Text, "\"ln10\"")
	assert.Contains(t, contents.Text, "\"log10E\"")
	assert.Contains(t, contents.Text, "3.14159") // pi value
	assert.Contains(t, contents.Text, "2.71828") // e value
	assert.Contains(t, contents.Text, "1.61803") // phi value
}

func TestPhiConstant(t *testing.T) {
	// Golden ratio should be approximately (1 + sqrt(5)) / 2
	assert.InDelta(t, 1.618033988749895, Phi, 0.0000001)
}

func TestSqrtPhiConstant(t *testing.T) {
	// sqrt(phi) should be approximately 1.272
	assert.InDelta(t, 1.272019649514069, SqrtPhi, 0.0000001)
}
