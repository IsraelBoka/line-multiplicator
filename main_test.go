package main

import "testing"

func TestCalcul(t *testing.T) {
	lines := []string{"3 * 2", "4 * 3"}
	got := Calcul(lines)
	want := 18
	if got != want {
		t.Errorf("Error was expecting %v but got %v", want, got)
	}
}
