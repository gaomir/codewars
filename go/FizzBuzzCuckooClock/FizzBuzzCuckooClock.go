package main

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzzCuckooClock is fun
func FizzBuzzCuckooClock(time string) string {

	timeElements := strings.Split(time, ":")
	hour, _ := strconv.Atoi(timeElements[0])
	minutes, _ := strconv.Atoi(timeElements[1])

	response := "tick"

	if minutes%3 == 0 {
		response = "Fizz"
	}

	if minutes%5 == 0 {
		response = "Buzz"
	}

	if minutes%15 == 0 {
		response = "Fizz Buzz"
	}

	if minutes == 0 {

		if hour == 0 {
			hour = 12
		}

		if hour > 12 {
			hour = hour % 12
		}

		response = strings.TrimRight(strings.Repeat("Cuckoo ", hour), " ")
	}

	if minutes == 30 {
		response = "Cuckoo"
	}

	return response
}

func main() {
	fmt.Println(FizzBuzzCuckooClock("13:34")) // tick
	fmt.Println(FizzBuzzCuckooClock("21:00")) // Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo
	fmt.Println(FizzBuzzCuckooClock("11:15")) // Fizz Buzz
	fmt.Println(FizzBuzzCuckooClock("03:03")) // Fizz
	fmt.Println(FizzBuzzCuckooClock("14:30")) //Cuckoo
	fmt.Println(FizzBuzzCuckooClock("08:55")) //Buzz
	fmt.Println(FizzBuzzCuckooClock("00:00")) //Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo
	fmt.Println(FizzBuzzCuckooClock("12:00")) //Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo
}
