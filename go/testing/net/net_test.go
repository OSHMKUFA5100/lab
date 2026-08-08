package net

import (
	"net"
	"testing"
)

func TestNet(t *testing.T) {
	s := net.JoinHostPort("127.0.0.1", "8080")

	t.Log(s)
	t.Log(s)
	addrs, err := net.LookupAddr("198.18.0.18")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addrs)
}
