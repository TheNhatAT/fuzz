package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// ReverseBytes cause error when running fuzz test because it's reversed each byte instead of each rune
func ReverseBytes(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// ReverseRunes still cause error when input is not a valid UTF-8 string
func ReverseRunes(s string) string {
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseUtf8 reverses the string and checks if the input is valid UTF-8
func ReverseUtf8(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"

	// Run ReverseBytes
	rev := ReverseBytes(input)
	doubleRev := ReverseBytes(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	// Run ReverseRunes
	revRunes := ReverseRunes(input)
	doubleRevRunes := ReverseRunes(revRunes)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", revRunes)
	fmt.Printf("reversed again: %q\n", doubleRevRunes)

	// Run ReverseUtf8
	revUtf8, revUtf8Err := ReverseUtf8(input)
	doubleRevUtf8, doubleRevUtf8Err := ReverseUtf8(revUtf8)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, error: %v\n", revUtf8, revUtf8Err)
	fmt.Printf("reversed again: %q, error: %v\n", doubleRevUtf8, doubleRevUtf8Err)
}
