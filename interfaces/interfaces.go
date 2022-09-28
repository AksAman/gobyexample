// pointer receivers
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

func (r *rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("g: %T, %#v\n", g, g)
	fmt.Printf("&g: %T, %#v\n", &g, &g)
	fmt.Printf("g.area(): %v\n", g.area())
	fmt.Printf("g.perim(): %v\n", g.perim())
	fmt.Println("------------------------")
}

func Run() {
	r := rectangle{width: 3, height: 4}
	fmt.Printf("r: %T, %#v\n", r, r)
	fmt.Printf("r: %T, %#v\n", r, &r)
	measure(&r)

	c := circle{radius: 5}
	fmt.Printf("c: %T, %#v\n", c, c)
	fmt.Printf("c: %T, %#v\n", &c, &c)
	measure(&c)
}
