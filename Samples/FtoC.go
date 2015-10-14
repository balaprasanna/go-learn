package main

import "fmt"

func main() {
	fmt.Println("Enter Temperature in Farenheat")
	var f float64
	fmt.Scanf("%f", &f)
	c := (f - 32) * 5 / 9
	fmt.Println("Celcius =",c)
}