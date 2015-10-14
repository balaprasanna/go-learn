package main

import "fmt"

func change(ptr *int) {
	*ptr = 10
}

func swap(a ,b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func main() {
	/*SWAP TWO INTS*/
	a,b := 2,4
	fmt.Println("a=",a,"b=",b)
	swap(&a,&b)
	fmt.Println("a=",a,"b=",b)
	
	/* USING NEW OPERATOR*/
	xPtr := new(int)
	*xPtr = 2
	fmt.Println("Before",*xPtr)
	change(xPtr)
	fmt.Println("After",*xPtr)
	
	/*NORMAL WAY OF DOING*/
	x := 5
	fmt.Println("Before change",x)
	change(&x)
	fmt.Println("After change",x)
}