package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.area(), r.perim())

	rp := &r
	rp2 := &r
	fmt.Println(rp.area(), rp.perim())
	fmt.Println(rp, rp2)
}
