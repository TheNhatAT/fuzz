# Fuzz Testing in Go

This project demonstrates the use of fuzz testing in Go to test string reversal functions. It includes three different implementations of string reversal and tests their correctness and behavior under various inputs, including edge cases and invalid UTF-8 strings.

## Project Structure

- **`main.go`**: Contains the main logic for string reversal functions and a demonstration of their usage.
- **`reverse_test.go`**: Includes unit tests and fuzz tests for the string reversal functions.
- **`go.mod`**: Defines the Go module and its dependencies.
- **`testdata/`**: Stores fuzzing corpus data generated during fuzz testing.
- **`.gitignore`**: Specifies files and directories to be ignored by Git.

## String Reversal Functions

1. **`ReverseBytes`**: Reverses the bytes of a string. This function does not handle multi-byte characters correctly.
2. **`ReverseRunes`**: Reverses the runes of a string. This function works for valid UTF-8 strings but may fail for invalid UTF-8 inputs.
3. **`ReverseUtf8`**: Reverses the runes of a string and validates the input as UTF-8. Returns an error if the input is not valid UTF-8.

## How to Run

### Prerequisites

- Go 1.22 or later installed on your system.

### Running the Application

To run the application and see the string reversal functions in action:

```bash
go run main.go
```

### Running Tests

To run the unit tests and fuzz tests:

```bash
go test ./...
```

### Running Fuzz Tests

To specifically run fuzz tests:

```bash
go test -fuzz=FuzzReverseBytes ./...
go test -fuzz=FuzzReverseRunes ./...
go test -fuzz=FuzzReverseUtf8 ./...
```

## Fuzz Testing

Fuzz testing is used to automatically generate test cases to uncover edge cases and bugs. The project includes fuzz tests for all three string reversal functions:

- **`FuzzReverseBytes`**: Tests `ReverseBytes` for correctness and UTF-8 validity.
- **`FuzzReverseRunes`**: Tests `ReverseRunes` for correctness and not have UTF-8 validation.
- **`FuzzReverseUtf8`**: Tests `ReverseUtf8` for correctness and error handling.

### Fuzz Corpus

The `testdata/` directory contains fuzzing corpus data generated during fuzz testing. This data helps reproduce specific edge cases discovered during fuzzing.

## Known Issues

- `ReverseBytes` does not handle multi-byte characters correctly.
- `ReverseRunes` may fail for invalid UTF-8 inputs.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

