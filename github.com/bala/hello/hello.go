package main

import("fmt"
		"math/rand"
		"life"
		)

func f(n int) {
	fmt.Println("In Go Routine")
	for i := 0; i < 10; i++ { fmt.Println(n, ":", i) }
}

func main() {	
	fmt.Println("Random Number" ,rand.Int())
	fmt.Println("Life" , life.GetLifeTimeSince())
	fmt.Println("Before Go Routine")  	
  	go f(0)
  	fmt.Println("After Go Routine")
  	var input string
  	fmt.Scanln(&input)
}