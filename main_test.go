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

	got, err = strParser("abc")
	if err == nil {
		t.Error("Error, input is string but it can pass through the process")
	}

	got, err = strParser("200-80")
	if err == nil {
		t.Error("Error, reverse range can pass through the process")
	}

	t.Log(got, want)

}
