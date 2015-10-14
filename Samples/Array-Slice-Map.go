package main

import "fmt"

func main() {
	x := []int{
			  48,96,86,68,
			  57,82,63,70,
			  37,34,83,27,
			  19,97, 9,17,
			}
	temp := x[0]
	for _,value := range x {
		if temp < value { temp = value }
	}
	fmt.Println("Smallest", temp)
	// numbers := [5]int{ 5,6,7,7,9 }
	// total := 0
	// // for i:= 0; i< len(numbers); i++ {
	// // 	fmt.Scanf( "%d", &numbers[i])
	// // 	total += numbers[i]
	// // 	fmt.Println("%p", &numbers[i])
	// // }
	// for i, value := range numbers {
	// 	//fmt.Scanf( "%d", &value); 
	// 	total += value
	// 	fmt.Printf("%p , %p", &value, &numbers[i])
	// 	fmt.Println(" " , value , " ")
	// }
	// output := total/len(numbers)
	// fmt.Println(output)
}