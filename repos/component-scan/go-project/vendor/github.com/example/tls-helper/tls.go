// Package tlshelper wraps crypto/tls for simplified dial operations.
package tlshelper

import (
	"crypto/tls"
	"net"
)

// Dial opens a TLS connection to addr.
func Dial(addr string) (net.Conn, error) {
	return tls.Dial("tcp", addr, &tls.Config{MinVersion: tls.VersionTLS12})
}
