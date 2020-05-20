package main

import "fmt"

func main() {
	functions := map[string]func(int, int) int {
		"add":		add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":	divide,
	}

	var currentNumber int
	fmt.Println("Enter the start number: ")
	fmt.Scan(&currentNumber)

	for true {
		var functioName string
		var number int
		fmt.Println("What function do you want to use")
		fmt.Scan(&functioName)
		if functioName = "done" {
			break
		}

		fmt.Println("Enter the number")
		fmt.Scan(&number)

		currentNumber = functions[functioName](currentNumber, number)
	}

	fmt.Println("Your number is")
	fmt.Println(currentNumber)


func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func mulitply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

