package main

import (
	"fmt"
	"math"
)

//Shape definitions

type circle struct {
	radius float64
}

type sphere struct {
	radius float64
}

type square struct {
	side float64
}

type cube struct {
	side float64
}

type rectangle struct {
	length, width float64
}

type rectPrism struct {
	length, width, height float64
}

type triangle struct {
	base, height float64
}

type triPrism struct {
	length, width, height float64
}

//Slices of each shape type

type squares []*square
type circles []*circle
type rectangles []*rectangle
type triangles []*triangle

type cubes []*cube
type spheres []*sphere
type rectPrisms []*rectPrism
type triPrisms []*triPrism

//Shape definitions for 2D and 3D

type shape interface {
	area() float64
}

type shape3d interface {
	surfArea() float64
	volume() float64
}

//Shape Wrappers

type shapeWrapper struct {
	s []*shape
}

type shape3dWrapper struct {
	s []*shape3d
}

//Area function definitions

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, float64(2))
}

func (s square) area() float64 {
	return math.Pow(s.side, float64(2))
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (t triangle) area() float64 {
	return t.base / 2 * t.height
}

//Volume function definitions

func (s sphere) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(s.radius, float64(3))
}

func (c cube) volume() float64 {
	return math.Pow(c.side, float64(3))
}

func (r rectPrism) volume() float64 {
	return r.height * r.length * r.width
}

func (t triPrism) volume() float64 {
	return t.length * t.width * t.height / 3
}

//Surface area function definitions

func (s sphere) surfArea() float64 {
	return 4 * math.Pi * math.Pow(s.radius, float64(2))
}

func (c cube) surfArea() float64 {
	return 6 * math.Pow(c.side, float64(2))
}

func (r rectPrism) surfArea() float64 {
	return r.length*r.width*2 + r.length*r.height*2 + r.width*r.height*2
}

func (t triPrism) surfArea() float64 {
	return t.length*t.width + t.length*t.height + t.width*t.height
}

//Functions to list shape properties

func info(z shape) {
	fmt.Printf("%+v\n", z)
	fmt.Println("Area: ", z.area())
}

func info3d(z shape3d) {
	fmt.Printf("%+v\n", z)
	fmt.Println("Surface Area: ", z.surfArea())
	fmt.Println("Volume: ", z.volume())
}

//Functions calculating totals for all shapes

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalSurfArea(shapes3d ...shape3d) float64 {
	var surfArea float64
	for _, s := range shapes3d {
		surfArea += s.surfArea()
	}
	return surfArea
}

func totalVolume(shapes3d ...shape3d) float64 {
	var volume float64
	for _, s := range shapes3d {
		volume += s.volume()
	}
	return volume
}

func main() {

	//2D shape maps

	sm := map[int]*square{
		0: {2},
		1: {3},
		2: {5},
	}

	cm := map[int]*circle{
		0: {5},
		1: {7},
		2: {9},
	}

	rm := map[int]*rectangle{
		0: {4, 7},
		1: {12, 3},
		2: {8, 9},
	}

	tm := map[int]*triangle{
		0: {5, 12},
		1: {2, 11},
		2: {20, 3},
	}

	//3D shape maps

	cbm := map[int]*cube{
		0: {5},
		1: {7},
		2: {19},
	}

	spm := map[int]*sphere{
		0: {9},
		1: {11},
		2: {21},
	}

	rpm := map[int]*rectPrism{
		0: {4, 9, 10},
		1: {20, 13, 2},
		2: {12, 12, 6},
	}

	tpm := map[int]*triPrism{
		0: {9, 9, 10},
		1: {15, 6, 27},
		2: {3, 20, 5},
	}

	//Make 2D shape slices

	s := make(squares, 0)
	c := make(circles, 0)
	r := make(rectangles, 0)
	t := make(triangles, 0)

	//Make 3D shape slices

	cu := make(cubes, 0)
	sp := make(spheres, 0)
	rp := make(rectPrisms, 0)
	tp := make(triPrisms, 0)

	//Copy map data into slices

	for _, d := range sm {
		s = append(s, d)
	}

	for _, d := range cm {
		c = append(c, d)
	}

	for _, d := range rm {
		r = append(r, d)
	}

	for _, d := range tm {
		t = append(t, d)
	}

	for _, d := range cbm {
		cu = append(cu, d)
	}

	for _, d := range spm {
		sp = append(sp, d)
	}

	for _, d := range rpm {
		rp = append(rp, d)
	}

	for _, d := range tpm {
		tp = append(tp, d)
	}

	//Print data

	fmt.Println("Square Data:")

	for _, d := range s {
		info(*d)
	}

	fmt.Println("Circle Data:")

	for _, d := range c {
		info(*d)
	}

	fmt.Println("Rectangle Data:")

	for _, d := range r {
		info(*d)
	}

	fmt.Println("Triangle Data:")

	for _, d := range t {
		info(*d)
	}

	fmt.Println("Cube Data:")

	for _, d := range cu {
		info3d(*d)
	}

	fmt.Println("Sphere Data:")

	for _, d := range sp {
		info3d(*d)
	}

	fmt.Println("Rectangular Prism Data:")

	for _, d := range rp {
		info3d(*d)
	}

	fmt.Println("Triangular Prism Data:")

	for _, d := range tp {
		info3d(*d)
	}

}
