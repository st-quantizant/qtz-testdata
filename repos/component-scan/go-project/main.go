package main

import (
	"github.com/example/logger"
	"github.com/example/tls-helper"
)

func main() {
	logger.Log("starting")
	tlshelper.Dial("example.com:443")
}
