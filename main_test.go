package main

import (
	"slices"
	"testing"
)

func TestStrParser(t *testing.T) {
	got, err := strParser("1-3")
	want := []int{1, 2, 3}
	if slices.Compare(got, want) != 0 {

		t.Errorf("got: %d, but want: %d", got, want)

	}

	t.Log(got, want, err)

}
