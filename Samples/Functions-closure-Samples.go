package main

import "fmt"

func sum(nums []float64) float64{
	total := 0.0
	for _,value := range nums {
		total += value;
	}
	return total/ float64(len(nums))
}

func halfIt(i int) (half int, status bool ) {
	half = i/2
	if i % 2 == 0 {
		status = true
	}else{
		status = false
	}
	return
}

func main() {
	nums := []float64{2,4,5,5,6}
	fmt.Println(sum(nums))
	fmt.Println(halfIt(2))
	fmt.Println(halfIt(1))
}