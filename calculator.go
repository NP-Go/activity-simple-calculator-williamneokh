package main

import (
	"fmt"
	"log"
)

func add(a, b int) int {

	//Insert code here
	return a + b
}

func subtract(a, b int) int {
	//Insert code here
	return a - b
}

func multiply(a, b int) int {
	//Insert code here
	return a * b
}

func divide(a, b int) int {
	//Insert code here
	//consider for b = 0
	if b == 0 {
		log.Fatal("You cannot use 0 in division")
	}
	return a / b
}

func main() {
	var a, b int
	var process string
	var result int
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (+, -, *, /)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	//Insert code here
	switch process {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)

	}
	fmt.Println("The answer is: ", result)
}
