package main

import ("fmt")

func add_values(a int, b int) int {

	var c= a + b

	return c

}

func main () {

	var a, b, c int

	fmt.Println("Enter the first number: ")

	fmt.Scan(&a)

	fmt.Println("Enter the second number: ")

	fmt.Scan(&b)

	c=add_values(a , b)

	fmt.Printf("sum: %d \n" , c)

	
}

