package main

import "testing"

func Test(t *testing.T) {

	got := ""
	expected := ""

	if got != expected {
		t.Errorf("got %s; want %s", got, expected)
	}
}
