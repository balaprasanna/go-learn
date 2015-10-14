package main

import "fmt"

func main() {
	var x string = "Hello world"
	var y string
	y = "bala"; y = y + "Prasanna"
	z := 12345; z = z + 100
	
	const c string = "I won't change again"
	fmt.Println(x + y)
	fmt.Println(z)
	fmt.Printf("%T %T %T", x,y,z)
	fmt.Println(c)

	/* Program to enter a number and multipy it by 2*/
	fmt.Println("Enter a number")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("output =", output)
}