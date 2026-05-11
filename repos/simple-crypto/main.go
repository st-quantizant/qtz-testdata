// Package main is a test fixture containing intentional weak crypto patterns.
// This file is NOT production code – it is used only by qtz-discovery-cli tests.
package main

import (
	"crypto/md5"  //nolint – intentional for testing detection
	"crypto/sha256"
	"fmt"
)

// weakHash demonstrates weak algorithm (MD5) detection.
func weakHash(data []byte) []byte {
	h := md5.New() //nolint
	h.Write(data)
	return h.Sum(nil)
}

// strongHash demonstrates a quantum-safer algorithm detection.
func strongHash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func main() {
	msg := []byte("hello")
	fmt.Printf("MD5:    %x\n", weakHash(msg))
	fmt.Printf("SHA256: %x\n", strongHash(msg))
}
