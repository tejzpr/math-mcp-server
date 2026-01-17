# Math MCP Server

A comprehensive mathematical operations MCP (Model Context Protocol) server written in Go. This server exposes 70+ mathematical tools that can be used by AI assistants and MCP-compatible clients.

[![CI](https://github.com/tejzpr/math-mcp-server/actions/workflows/ci.yml/badge.svg)](https://github.com/tejzpr/math-mcp-server/actions/workflows/ci.yml)
[![Docker](https://github.com/tejzpr/math-mcp-server/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/tejzpr/math-mcp-server/actions/workflows/docker-publish.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tejzpr/math-mcp-server)](https://goreportcard.com/report/github.com/tejzpr/math-mcp-server)
[![Docker Hub](https://img.shields.io/docker/v/tejzpr/math-mcp-server?label=Docker%20Hub)](https://hub.docker.com/r/tejzpr/math-mcp-server)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](https://mozilla.org/MPL/2.0/)

## Features

- **70+ Mathematical Tools** organized into 15 categories
- **Environment-based configuration** for enabling/disabling tool categories
- **Dual transport support**: stdio (default) and HTTP
- **MCP Resources**: Mathematical constants exposed as a resource
- **Comprehensive test coverage**

## Quick Start

### Prerequisites

- **Go 1.23 or later** - [Download and install Go](https://go.dev/doc/install)

### Installation

#### Option 1: Using `go run` (Recommended for MCP clients)

No installation needed! Configure your MCP client to run directly:

```json
{
    "mcpServers": {
        "math": {
            "command": "go",
            "args": ["run", "github.com/tejzpr/math-mcp-server@latest"],
            "env": {
                "MATH_CATEGORIES": "all"
            }
        }
    }
}
```

#### Option 2: Using `go install`

```bash
go install github.com/tejzpr/math-mcp-server@latest
```

Then configure your MCP client:

```json
{
    "mcpServers": {
        "math": {
            "command": "math-mcp-server",
            "env": {
                "MATH_CATEGORIES": "all"
            }
        }
    }
}
```

#### Option 3: Build from source

```bash
git clone https://github.com/tejzpr/math-mcp-server.git
cd math-mcp-server
go build -o math-mcp-server .
```

#### Option 4: Docker

**Build the image:**
```bash
docker build -t math-mcp-server .
```

**Run (HTTP transport - default):**
```bash
docker run -d --rm \
  -e MATH_CATEGORIES=all \
  -p 8080:8080 \
  math-mcp-server
```

The server will be available at `http://localhost:8080/mcp`.

**Run with stdio (for MCP clients):**
```bash
docker run -i --rm \
  -e TRANSPORT=stdio \
  -e MATH_CATEGORIES=all \
  math-mcp-server
```

**MCP client configuration with Docker (stdio):**
```json
{
    "mcpServers": {
        "math": {
            "command": "docker",
            "args": ["run", "-i", "--rm", "-e", "TRANSPORT=stdio", "math-mcp-server"],
            "env": {
                "MATH_CATEGORIES": "all"
            }
        }
    }
}
```

## Configuration

### Environment Variables

| Variable | Description | Default | Values |
|----------|-------------|---------|--------|
| `MATH_CATEGORIES` | Comma-separated list of categories to enable | `all` | `all` or category names |
| `TRANSPORT` | Transport protocol | `stdio` | `stdio`, `http` |

### Command Line Flags

| Flag | Description |
|------|-------------|
| `-t`, `--transport` | Override transport (overrides `TRANSPORT` env var) |

### Example Configurations

**Enable all categories (default):**
```json
{
    "mcpServers": {
        "math": {
            "command": "go",
            "args": ["run", "github.com/tejzpr/math-mcp-server@latest"],
            "env": {
                "MATH_CATEGORIES": "all"
            }
        }
    }
}
```

**Enable only specific categories:**
```json
{
    "mcpServers": {
        "math": {
            "command": "go",
            "args": ["run", "github.com/tejzpr/math-mcp-server@latest"],
            "env": {
                "MATH_CATEGORIES": "arithmetic,trig,statistics"
            }
        }
    }
}
```

**Use HTTP transport:**
```json
{
    "mcpServers": {
        "math": {
            "command": "go",
            "args": ["run", "github.com/tejzpr/math-mcp-server@latest"],
            "env": {
                "TRANSPORT": "http"
            }
        }
    }
}
```

## Available Categories

| Category | ID | Tools |
|----------|-----|-------|
| **Arithmetic** | `arithmetic` | `add`, `subtract`, `multiply`, `divide`, `mod`, `remainder`, `abs` |
| **Power & Roots** | `power` | `pow`, `pow10`, `sqrt`, `cbrt`, `exp`, `exp2`, `expm1`, `hypot` |
| **Logarithmic** | `logarithm` | `log`, `log10`, `log2`, `log1p`, `logb` |
| **Trigonometric** | `trig` | `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `atan2`, `sincos` |
| **Hyperbolic** | `hyperbolic` | `sinh`, `cosh`, `tanh`, `asinh`, `acosh`, `atanh` |
| **Rounding** | `rounding` | `ceil`, `floor`, `round`, `round_to_even`, `trunc` |
| **Comparison** | `comparison` | `max`, `min`, `dim`, `copysign` |
| **Special Functions** | `special` | `gamma`, `lgamma`, `erf`, `erfc`, `erfinv`, `erfcinv`, `j0`, `j1`, `y0`, `y1` |
| **Float Utilities** | `float_utils` | `frexp`, `ldexp`, `modf`, `ilogb`, `nextafter`, `fma`, `signbit`, `is_nan`, `is_inf` |
| **Conversions** | `conversion` | `degrees_to_radians`, `radians_to_degrees` |
| **Number Theory** | `number_theory` | `gcd`, `lcm`, `factorial`, `fibonacci`, `is_prime`, `prime_factors` |
| **Statistics** | `statistics` | `sum`, `product`, `mean`, `median`, `mode`, `variance`, `std_dev`, `range_stat` |
| **Bitwise** | `bitwise` | `bit_and`, `bit_or`, `bit_xor`, `bit_not`, `bit_left_shift`, `bit_right_shift` |
| **Complex Numbers** | `complex` | `complex_abs`, `complex_phase`, `complex_conj`, `complex_exp`, `complex_log`, `complex_sqrt`, `complex_pow`, `complex_sin`, `complex_cos`, `complex_tan`, `complex_polar`, `complex_rect` |
| **Constants** | `constants` | Resource: `math://constants` |

## Tool Reference

### Arithmetic (`arithmetic`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `add` | Add two numbers | `a`, `b` |
| `subtract` | Subtract b from a | `a`, `b` |
| `multiply` | Multiply two numbers | `a`, `b` |
| `divide` | Divide a by b | `a`, `b` |
| `mod` | Floating-point modulo | `x`, `y` |
| `remainder` | IEEE 754 remainder | `x`, `y` |
| `abs` | Absolute value | `x` |

### Power & Roots (`power`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `pow` | x raised to power y | `x`, `y` |
| `pow10` | 10 raised to power n | `n` |
| `sqrt` | Square root | `x` |
| `cbrt` | Cube root | `x` |
| `exp` | e^x | `x` |
| `exp2` | 2^x | `x` |
| `expm1` | e^x - 1 (accurate for small x) | `x` |
| `hypot` | sqrt(x² + y²) | `x`, `y` |

### Logarithmic (`logarithm`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `log` | Natural logarithm | `x` |
| `log10` | Base-10 logarithm | `x` |
| `log2` | Base-2 logarithm | `x` |
| `log1p` | ln(1 + x) (accurate for small x) | `x` |
| `logb` | Binary exponent | `x` |

### Trigonometric (`trig`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `sin` | Sine (radians) | `x` |
| `cos` | Cosine (radians) | `x` |
| `tan` | Tangent (radians) | `x` |
| `asin` | Arc sine | `x` |
| `acos` | Arc cosine | `x` |
| `atan` | Arc tangent | `x` |
| `atan2` | Arc tangent of y/x | `y`, `x` |
| `sincos` | Returns both sin and cos | `x` |

### Hyperbolic (`hyperbolic`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `sinh` | Hyperbolic sine | `x` |
| `cosh` | Hyperbolic cosine | `x` |
| `tanh` | Hyperbolic tangent | `x` |
| `asinh` | Inverse hyperbolic sine | `x` |
| `acosh` | Inverse hyperbolic cosine | `x` |
| `atanh` | Inverse hyperbolic tangent | `x` |

### Rounding (`rounding`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `ceil` | Round up | `x` |
| `floor` | Round down | `x` |
| `round` | Round to nearest (half away from zero) | `x` |
| `round_to_even` | Banker's rounding | `x` |
| `trunc` | Truncate towards zero | `x` |

### Comparison (`comparison`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `max` | Larger of x and y | `x`, `y` |
| `min` | Smaller of x and y | `x`, `y` |
| `dim` | max(x - y, 0) | `x`, `y` |
| `copysign` | x with sign of y | `x`, `y` |

### Special Functions (`special`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `gamma` | Gamma function | `x` |
| `lgamma` | Log of absolute gamma | `x` |
| `erf` | Error function | `x` |
| `erfc` | Complementary error function | `x` |
| `erfinv` | Inverse error function | `x` |
| `erfcinv` | Inverse complementary error function | `x` |
| `j0` | Bessel function (first kind, order 0) | `x` |
| `j1` | Bessel function (first kind, order 1) | `x` |
| `y0` | Bessel function (second kind, order 0) | `x` |
| `y1` | Bessel function (second kind, order 1) | `x` |

### Float Utilities (`float_utils`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `frexp` | Break into fraction and exponent | `x` |
| `ldexp` | frac × 2^exp | `frac`, `exp` |
| `modf` | Split into integer and fraction | `x` |
| `ilogb` | Binary exponent as integer | `x` |
| `nextafter` | Next representable float | `x`, `y` |
| `fma` | Fused multiply-add: x×y + z | `x`, `y`, `z` |
| `signbit` | Check if sign bit is set | `x` |
| `is_nan` | Check if NaN | `x` |
| `is_inf` | Check if infinity | `x`, `sign` (optional) |

### Conversions (`conversion`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `degrees_to_radians` | Convert degrees to radians | `degrees` |
| `radians_to_degrees` | Convert radians to degrees | `radians` |

### Number Theory (`number_theory`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `gcd` | Greatest common divisor | `a`, `b` |
| `lcm` | Least common multiple | `a`, `b` |
| `factorial` | n! | `n` |
| `fibonacci` | Nth Fibonacci number | `n` |
| `is_prime` | Check if prime | `n` |
| `prime_factors` | Prime factorization | `n` |

### Statistics (`statistics`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `sum` | Sum of array | `numbers` (array) |
| `product` | Product of array | `numbers` (array) |
| `mean` | Arithmetic mean | `numbers` (array) |
| `median` | Median value | `numbers` (array) |
| `mode` | Most frequent value(s) | `numbers` (array) |
| `variance` | Population variance | `numbers` (array) |
| `std_dev` | Population standard deviation | `numbers` (array) |
| `range_stat` | max - min | `numbers` (array) |

### Bitwise (`bitwise`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `bit_and` | Bitwise AND | `a`, `b` |
| `bit_or` | Bitwise OR | `a`, `b` |
| `bit_xor` | Bitwise XOR | `a`, `b` |
| `bit_not` | Bitwise NOT | `a` |
| `bit_left_shift` | Left shift (a << n) | `a`, `n` |
| `bit_right_shift` | Right shift (a >> n) | `a`, `n` |

### Complex Numbers (`complex`)

| Tool | Description | Parameters |
|------|-------------|------------|
| `complex_abs` | Magnitude | `real`, `imag` |
| `complex_phase` | Phase (argument) | `real`, `imag` |
| `complex_conj` | Complex conjugate | `real`, `imag` |
| `complex_exp` | Complex exponential | `real`, `imag` |
| `complex_log` | Complex logarithm | `real`, `imag` |
| `complex_sqrt` | Complex square root | `real`, `imag` |
| `complex_pow` | Complex power | `x_real`, `x_imag`, `y_real`, `y_imag` |
| `complex_sin` | Complex sine | `real`, `imag` |
| `complex_cos` | Complex cosine | `real`, `imag` |
| `complex_tan` | Complex tangent | `real`, `imag` |
| `complex_polar` | Convert to polar form | `real`, `imag` |
| `complex_rect` | Convert from polar form | `r`, `theta` |

### Constants Resource (`constants`)

The `math://constants` resource provides commonly used mathematical constants:

```json
{
  "pi": 3.141592653589793,
  "e": 2.718281828459045,
  "phi": 1.618033988749895,
  "sqrt2": 1.4142135623730951,
  "sqrtE": 1.6487212707001282,
  "sqrtPi": 1.7724538509055159,
  "sqrtPhi": 1.272019649514069,
  "ln2": 0.6931471805599453,
  "log2E": 1.4426950408889634,
  "ln10": 2.302585092994046,
  "log10E": 0.4342944819032518,
  "maxFloat64": 1.7976931348623157e+308,
  "smallestNonzeroFloat64": 5e-324
}
```

## Usage Examples

### With Cursor IDE

Add to your Cursor MCP settings (`.cursor/mcp.json` or global settings):

```json
{
    "mcpServers": {
        "math": {
            "command": "go",
            "args": ["run", "github.com/tejzpr/math-mcp-server@latest"],
            "env": {
                "MATH_CATEGORIES": "all"
            }
        }
    }
}
```

### With Claude Desktop

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json` on macOS):

```json
{
    "mcpServers": {
        "math": {
            "command": "go",
            "args": ["run", "github.com/tejzpr/math-mcp-server@latest"],
            "env": {
                "MATH_CATEGORIES": "all"
            }
        }
    }
}
```

### HTTP Mode

Start the server in HTTP mode:

```bash
TRANSPORT=http go run github.com/tejzpr/math-mcp-server@latest
```

Or with the flag:

```bash
go run github.com/tejzpr/math-mcp-server@latest -t http
```

The server will listen on `http://localhost:8080/mcp`.

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o math-mcp-server .
```

### Project Structure

```
math-mcp-server/
├── main.go                 # Entry point
├── config/
│   ├── config.go          # Configuration loading
│   └── config_test.go     # Config tests
└── handlers/
    ├── registry.go        # Tool registration
    ├── arithmetic.go      # Arithmetic tools
    ├── power.go           # Power & root tools
    ├── logarithm.go       # Logarithmic tools
    ├── trig.go            # Trigonometric tools
    ├── hyperbolic.go      # Hyperbolic tools
    ├── rounding.go        # Rounding tools
    ├── comparison.go      # Comparison tools
    ├── special.go         # Special function tools
    ├── float_utils.go     # Float utility tools
    ├── conversion.go      # Conversion tools
    ├── number_theory.go   # Number theory tools
    ├── statistics.go      # Statistics tools
    ├── bitwise.go         # Bitwise tools
    ├── complex.go         # Complex number tools
    ├── constants.go       # Constants resource
    └── *_test.go          # Tests for each category
```

## License

Mozilla Public License 2.0 - see [LICENSE](LICENSE) for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Built with [mcp-go](https://github.com/mark3labs/mcp-go) - Go implementation of the Model Context Protocol
- Leverages Go's standard `math` and `math/cmplx` packages
