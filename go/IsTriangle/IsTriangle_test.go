package main

import "testing"

func TestIsTriangleFalse(t *testing.T) {
	response := IsTriangle(5, 1, 2)
	if response == true {
		t.Errorf("Response was incorrect, got: %v, want %v.", response, false)
	}
}

func TestIsTriangleTrue(t *testing.T) {
	response := IsTriangle(4, 2, 3)
	if response == false {
		t.Errorf("Response was incorrect, got: %v, want %v.", response, true)
	}
}
