package main

import ("fmt"; "math")

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) DoubleYs(){
	c.x *= 2; c.y *= 2; c.r *= 2;
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r*c.r	
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
	c := *cPtr

	fmt.Println(CircleArea(c1,c2, c))
	fmt.Println(cPtr.Area())
	fmt.Println(c1.Area())
	fmt.Println(c2.Area())

	//fmt.Println("Before =",c1, "After =", c1.DoubleYs())
	fmt.Println("Before =", *cPtr); 
	cPtr.DoubleYs();
	fmt.Println("After =", *cPtr);

	fmt.Println("Before =", c2); 
	c2.DoubleYs();
	fmt.Println("After =", c2);
	//fmt.Println("Before =",c1); c1.DoubleYs(); fmt.Println("After =", c1);
	
	// *cPtr.x = 10.0
	// *cPtr.y = 20.0
	// *cPtr.r = 5.0
	// fmt.Println(*cPtr)
}