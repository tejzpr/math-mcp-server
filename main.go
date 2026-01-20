// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

// Package main provides the entry point for the math MCP server.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sagacient/math-mcp-server/config"
	"github.com/sagacient/math-mcp-server/handlers"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Parse command line flags
	var transport string
	flag.StringVar(&transport, "transport", "", "Transport type (stdio or http)")
	flag.StringVar(&transport, "t", "", "Transport type (stdio or http) (shorthand)")
	flag.Parse()

	// Load configuration
	cfg := config.LoadConfig()

	// Override transport from flag if provided
	if transport != "" {
		cfg.Transport = transport
	}

	// Log enabled categories
	enabled := cfg.EnabledCategories()
	if len(enabled) == 0 {
		log.Fatal("No categories enabled. Set MATH_CATEGORIES environment variable.")
	}

	categoryNames := make([]string, len(enabled))
	for i, cat := range enabled {
		categoryNames[i] = string(cat)
	}
	log.Printf("Enabled categories: %s", strings.Join(categoryNames, ", "))

	// Create MCP server
	mcpServer := server.NewMCPServer(
		"math-mcp-server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, false),
		server.WithRecovery(),
	)

	// Register tools based on configuration
	registry := handlers.NewRegistry()
	registry.RegisterTools(mcpServer, cfg)

	// Register constants resource if enabled
	handlers.RegisterConstants(mcpServer, cfg)

	// Start server based on transport
	switch cfg.Transport {
	case "http":
		httpServer := server.NewStreamableHTTPServer(mcpServer)
		port := os.Getenv("MATH_PORT")
		if port == "" {
			port = "8080"
		}
		log.Printf("Starting HTTP server on :%s/mcp", port)
		if err := httpServer.Start(":" + port); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	default:
		fmt.Fprintln(os.Stderr, "Starting stdio server...")
		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatalf("Stdio server error: %v", err)
		}
	}
}
