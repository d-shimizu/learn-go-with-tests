package main

import "testing"

func TestRomanNumeals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
