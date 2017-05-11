package main

import "testing"

func TestPotatoesCase1(t *testing.T) {
	response := Potatoes(99, 100, 98)
	if response != 50 {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "50")
	}
}

func TestPotatoesCase2(t *testing.T) {
	response := Potatoes(82, 127, 80)
	if response != 114 {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "114")
	}
}
