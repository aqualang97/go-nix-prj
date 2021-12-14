package main

import (
	"fmt"
	"math"
)

type Position struct {
	x, y float64
}
type Rectangle struct {
	Position
	x2, y2 float64
}
type Circle struct {
	x, y, r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x, r.y, r.x2, r.y2)
	w := distance(r.x, r.y, r.x2, r.y)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x, r.y, r.x2, r.y2)
	w := distance(r.x, r.y, r.x2, r.y)
	return 2*l + 2*w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

/*
func totalArea(c *Circle, r *Rectangle) (total float64) {
	total += c.area()
	total += r.area()
	return total
}
*/
// ... - neskolko

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) (total float64) {
	for _, shape := range shapes {
		total += shape.area()
	}
	return total
}
func totalPerimeter(shapes ...Shape) (totalPerimeter float64) {
	for _, shape := range shapes {
		totalPerimeter += shape.perimeter()
	}
	return totalPerimeter
}
func main() {

	rectangle := &Rectangle{
		Position: Position{
			x: 0,
			y: 0,
		},
		x2: 10,
		y2: 10,
	}
	circle := &Circle{
		x: 0,
		y: 0,
		r: 5,
	}

	fmt.Println(rectangle.area())
	fmt.Println(circle.area())
	//fmt.Println(totalArea(&circle,&rectangle))
	//interfases-->
	fmt.Println(totalArea(circle, rectangle, rectangle, rectangle, rectangle))

	fmt.Println(circle.perimeter())
	fmt.Println(rectangle.perimeter())
	fmt.Println(totalPerimeter(circle, rectangle))

}
