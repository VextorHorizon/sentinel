package main

import (
	"testing"
)

func TestStrParser(t *testing.T) {
	got, err := strParser("1-3")
	t.Log(got, err)
}
