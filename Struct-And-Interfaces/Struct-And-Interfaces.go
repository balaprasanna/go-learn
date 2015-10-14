package main

import ("fmt"; "math")

type Circle struct {
	x float64
	y float64
	r float64
}
func CircleArea(circles ...Circle) []float64 {
	areas := make([]float64, len(circles), len(circles) * 2)
	for i,c := range circles {
		areas[i] = math.Pi * c.r*c.r
	}
	return areas
}

func main() {
	c1 := Circle{x: 0,y: 10,r: 5}
	c2 := Circle{0, 4, 10}

	cPtr := new(Circle)
	cPtr = &c1
	localC := *cPtr;
	fmt.Println("X:",localC.x, "Y:",localC.y)
	fmt.Println(CircleArea(c1,c2, *cPtr))


	// *cPtr.x = 10.0
	// *cPtr.y = 20.0
	// *cPtr.r = 5.0
	// fmt.Println(*cPtr)
}