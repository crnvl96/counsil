# Counsil

A CLI tool to automate the installation and management of development tools and dependencies.

Counsil installs essential development tools using various package managers (mise, yay, uv, npm, Go) in the correct order, handling dependencies efficiently.

> You can also check the documentation at [pkg.go.dev](https://pkg.go.dev/github.com/crnvl96/counsil)

## Installation

### Using Go

```bash
go install github.com/crnvl96/counsil@latest
```

### From Releases

Download the latest binary from the [releases page](https://github.com/crnvl96/counsil/releases).

## Usage

```bash
counsil [flags]
```

### Flags

None currently. The tool runs with default settings.

### Examples

Install all tools:

```bash
counsil
```

### Output

The tool installs tools sequentially or concurrently where possible, printing progress and success/error messages for each tool.

## Building from Source

```bash
git clone https://github.com/crnvl96/counsil.git
cd counsil
go build -o counsil main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
