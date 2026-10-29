package main

import (
	"slices"
	"testing"
)

func TestStrParser(t *testing.T) {
	got, err := strParser("1-3")
	if err != nil {
		t.Errorf("Error, strParser couldn't process correct input")
	}
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

	got, err = strParser("1")
	if err == nil {
		t.Error("Error, the process is processing a single input")
	}

	got, err = strParser("20-100000")
	if err == nil {
		t.Error("Error, the process is processing out of port range")
	}
	t.Log(got, want)

}
