package main

import "testing"

func Test1Plus1(t *testing.T) {
	expected := 2
	got := 1 + 1
	if got != expected {
		t.Errorf("1Plus1 = %d; want %d", got, expected)
	}
}
