package internal_test

import (
	"ghact/internal"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"1 + 2 == 3", 1, 2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := internal.Add(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
