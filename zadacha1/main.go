package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter number: -5 ")

	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("Namber is Positive")
	} else if number < 0 {
		fmt.Println("Namber is negative")
	} else {
		fmt.Println("Number is zero")
	}
}
