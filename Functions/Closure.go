package main

import "fmt"

func makeEvenGenerator() (func() uint, func() uint) {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  },
  func() (ret uint) {
    ret = i
    i += 5
    return
  }
}

func first() {
  fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}
func third() {
  fmt.Println("3rd")
}
func close() {
  fmt.Println("Close")
}

func main() {  
	defer func () {
		str := recover()
		fmt.Println(str)
	}()
	f1, f2 := makeEvenGenerator()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f1())
	fmt.Println(f2())

	
	defer close()
	defer first()
	defer second()
	defer third()
	
	if x:= 5; x == 5 {
		fmt.Println("If")
		//return
		panic("Am runtime panics..")
	} else {
		fmt.Println("else")
	}
	fmt.Println("Need to close")

  // x := 0
  // increment := func() int {
  //   x++
  //   return x
  // }
  // x = 10
  // fmt.Println(increment())
  // fmt.Println(increment())
  // fmt.Println(increment())
  // x = 20
  // fmt.Println(increment())
	// third()
	// second()
	// third()
	// second()
}