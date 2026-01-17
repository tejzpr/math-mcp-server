// SPDX-License-Identifier: MPL-2.0
// Copyright 2026 Tejus Pratap <tejzpr@gmail.com>
//
// See CONTRIBUTORS.md for full contributor list.

package handlers

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/tejzpr/math-mcp-server/config"

	"github.com/mark3labs/mcp-go/mcp"
)

// registerNumberTheory registers number theory tools.
func (r *Registry) registerNumberTheory() {
	cat := config.CategoryNumberTheory

	// GCD
	r.addTool(
		mcp.NewTool("gcd",
			mcp.WithDescription("Greatest common divisor of a and b"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First integer")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second integer")),
		),
		gcdHandler,
		cat,
	)

	// LCM
	r.addTool(
		mcp.NewTool("lcm",
			mcp.WithDescription("Least common multiple of a and b"),
			mcp.WithNumber("a", mcp.Required(), mcp.Description("First integer")),
			mcp.WithNumber("b", mcp.Required(), mcp.Description("Second integer")),
		),
		lcmHandler,
		cat,
	)

	// Factorial
	r.addTool(
		mcp.NewTool("factorial",
			mcp.WithDescription("Factorial of n (n!)"),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Non-negative integer")),
		),
		factorialHandler,
		cat,
	)

	// Fibonacci
	r.addTool(
		mcp.NewTool("fibonacci",
			mcp.WithDescription("Nth Fibonacci number (0-indexed: fib(0)=0, fib(1)=1)"),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Non-negative integer index")),
		),
		fibonacciHandler,
		cat,
	)

	// IsPrime
	r.addTool(
		mcp.NewTool("is_prime",
			mcp.WithDescription("Check if n is a prime number"),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Integer to check")),
		),
		isPrimeHandler,
		cat,
	)

	// PrimeFactors
	r.addTool(
		mcp.NewTool("prime_factors",
			mcp.WithDescription("Prime factorization of n"),
			mcp.WithNumber("n", mcp.Required(), mcp.Description("Positive integer")),
		),
		primeFactorsHandler,
		cat,
	)
}

func gcdHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireInt("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := gcd(abs64(int64(a)), abs64(int64(b)))
	return mcp.NewToolResultText(fmt.Sprintf("%d", result)), nil
}

func lcmHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, err := req.RequireInt("a")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	b, err := req.RequireInt("b")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if a == 0 || b == 0 {
		return mcp.NewToolResultText("0"), nil
	}
	absA := abs64(int64(a))
	absB := abs64(int64(b))
	result := absA / gcd(absA, absB) * absB
	return mcp.NewToolResultText(fmt.Sprintf("%d", result)), nil
}

func factorialHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if n < 0 {
		return mcp.NewToolResultError("factorial undefined for negative numbers"), nil
	}
	if n > 170 {
		return mcp.NewToolResultError("factorial too large (max n=170 for float64)"), nil
	}
	result := factorial(n)
	return mcp.NewToolResultText(result.String()), nil
}

func fibonacciHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if n < 0 {
		return mcp.NewToolResultError("fibonacci undefined for negative indices"), nil
	}
	if n > 1000 {
		return mcp.NewToolResultError("fibonacci index too large (max n=1000)"), nil
	}
	result := fibonacci(n)
	return mcp.NewToolResultText(result.String()), nil
}

func isPrimeHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result := isPrime(int64(n))
	return mcp.NewToolResultText(fmt.Sprintf("%t", result)), nil
}

func primeFactorsHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	n, err := req.RequireInt("n")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if n <= 0 {
		return mcp.NewToolResultError("prime factorization requires a positive integer"), nil
	}
	factors := primeFactors(int64(n))
	strs := make([]string, len(factors))
	for i, f := range factors {
		strs[i] = fmt.Sprintf("%d", f)
	}
	return mcp.NewToolResultText(strings.Join(strs, " × ")), nil
}

// Helper functions

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func fibonacci(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int64) []int64 {
	var factors []int64

	// Handle 2s
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Handle odd factors
	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is still > 1, it's a prime factor
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
