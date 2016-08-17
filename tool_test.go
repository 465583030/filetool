package main

import "testing"

func TestParseSize(t *testing.T) {
	tests := []struct {
		s    string
		want int64
	}{
		{"0K", 0},
		{"123", 123},
		{"77K", 77 * 1024},
		{"77M", 77 * 1024 * 1024},
		{"77G", 77 * 1024 * 1024 * 1024},
		{"77T", 77 * 1024 * 1024 * 1024 * 1024},
	}

	for _, te := range tests {
		act, err := parseSize(te.s)
		if err != nil {
			t.Fatal(err)
		}
		if act != te.want {
			t.Fatalf("parse %s, got %d, want %d", te.s, act, te.want)
		}
	}
}
