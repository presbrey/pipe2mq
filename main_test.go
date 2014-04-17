package main

import (
	"testing"
)

func TestIncoming(t *testing.T) {
	incoming <- newRequest("1")
	incoming <- newRequest("2")
	incoming <- newRequest("3")
}
