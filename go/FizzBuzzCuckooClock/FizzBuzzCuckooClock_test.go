package main

import "testing"

func TestFizzBuzzCuckooClock(t *testing.T) {
	response := FizzBuzzCuckooClock("13:34")
	if response != "tick" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "tick")
	}
}

func TestFizzBuzzCuckooClockTwentyNineOclock(t *testing.T) {
	response := FizzBuzzCuckooClock("21:00")
	if response != "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo")
	}
}

func TestFizzBuzzCuckooClockFifteenMinutes(t *testing.T) {
	response := FizzBuzzCuckooClock("11:15")
	if response != "Fizz Buzz" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Fizz Buzz")
	}
}

func TestFizzBuzzCuckooClockMinutesDivisibleBy3(t *testing.T) {
	response := FizzBuzzCuckooClock("03:03")
	if response != "Fizz" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Fizz")
	}
}

func TestFizzBuzzCuckooClockMinutesEqual30(t *testing.T) {
	response := FizzBuzzCuckooClock("14:30")
	if response != "Cuckoo" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Cuckoo")
	}
}

func TestFizzBuzzCuckooClockMinutesDivisibleBy5(t *testing.T) {
	response := FizzBuzzCuckooClock("08:55")
	if response != "Buzz" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Buzz")
	}
}

func TestFizzBuzzCuckooClock12AM(t *testing.T) {
	response := FizzBuzzCuckooClock("00:00")
	if response != "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo")
	}
}

func TestFizzBuzzCuckooClock12PM(t *testing.T) {
	response := FizzBuzzCuckooClock("12:00")
	if response != "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo" {
		t.Errorf("Response was incorrect, got: %v, want: %v.", response, "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo")
	}
}
