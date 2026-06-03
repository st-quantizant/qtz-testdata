// Package logger provides simple structured logging.
package logger

import (
	"fmt"
	"time"
)

// Log writes a timestamped message to stdout.
func Log(msg string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), msg)
}
