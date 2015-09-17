//package main
package shapepk

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
    Perimeter() float64
}

type Circle struct {
	r float64
}

type Rectangle struct {
	l float64
	b float64
}

func (r *Rectangle) Area() float64 {
	return r.l * r.b
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.b)
}

func main() {
	c := Circle{3.0}
	fmt.Println("The area of the circle", c.Area())
	fmt.Println("The perimeter of the circle", c.Perimeter())
	r := Rectangle{2.0, 4.0}
	fmt.Println("The area of the Rectangle", r.Area())
	fmt.Println("The perimeter of the Rectangle", r.Perimeter())

}
