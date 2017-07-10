package main

import (
	"fmt"
	"time"
)

func main() {
	name := "anıl"
	fmt.Println("HELLO WORLD BABUŞ")
	fmt.Println("My name is " + name)

	isTrue := true
	fmt.Println("Is it true?", isTrue)

	age := 22
	fmt.Println("I am ", age, "years old ")

	for i := 0; i <= 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	t := time.Now()

	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!")
	default:
		fmt.Println("It's weekday")
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	}
}
