package main

import "fmt"

func main() {

	age := 33

	if age < 13 {
		fmt.Println("Wow you are young!")
	}
	if age > 12 && age < 20 {
		fmt.Println("You are a teenager")
	}
	if age >19 && age < 30 {
		fmt.Println("You are in your twenties")
	}
	if age > 29 && age < 40 {
		fmt.Println("You are in your thirties")
	}
	if age > 39 && age < 50 {
		fmt.Println("You are getting there!")
	}
	if age > 49 {
		fmt.Println("Over the hill!")
	}

}
