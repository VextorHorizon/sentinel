package main

import (
	"testing"
)

func TestStrParser(t *testing.T) {
	got, err := strParser("1")
	t.Log(got, err)

}
