package main

import "fmt"

func AverageVardiac(symbol string, args ...float64) (result float64) {
	fmt.Println("symbol ",symbol)
	total := 0.0
	for _, value := range args { total += value }
	result = total/ float64(len(args))
	return
}
func Average(numbers []float64) (result float64, total float64) {
	total = 0.0
	for _,value := range numbers { total += value; }
	result = total/ float64(len(numbers))
	return
}

func main() {
	x := []float64{5,5,5,5,5,5,5}
	fmt.Println("Input ",x)
	
	result,total := Average(x)
	fmt.Println(result, total)
	
	fmt.Println(AverageVardiac("$",1,2,3,4,5,5,6,6,6,7,7,7))
	fmt.Println(AverageVardiac("^",x...))
}