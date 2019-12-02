package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := sayHello()
	if got != "hello world" {
		t.Errorf("got %s; want hello world", got)
	}
}
