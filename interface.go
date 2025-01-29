package main

import "fmt"

// type rumus interface {
// 	luas() float64
// 	keliling() float64
// }

// type persegi struct {
// 	sisi float64
// }

// func (p persegi) luas() float64 {
// 	return p.sisi * p.sisi
// }

// type lingkaran struct {
// 	jari uint
// }

// func (l lingkaran) luas() float64 {
// 	return 3.14 * 2 * float64(l.jari)
// }

// func Intercface() {
// 	fmt.Println(persegi{10.0}.luas())
// 	bangun_2 := lingkaran{7.0}
// 	fmt.Println(bangun_2.luas())
// }

type shape interface {
	area() float64
	around() float64
}

type rectangle struct {
	width, height float64
}

type square struct {
	width float64
}

type circle struct {
	diameter float64
}

// method untuk mencari luas bangun datar
func (c circle) area() float64 {
	return c.diameter / 2 * c.diameter / 2 * 3.14
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (s square) area() float64 {
	return s.width * s.width
}

//method digunakan untuk mencari keliling bangun datar

func (c circle) around() float64 {
	return 2 * c.diameter * 3.14
}

func (r rectangle) around() float64 {
	return 2 * (r.width + r.height)
}

func (s square) around() float64 {
	return 4 * s.width
}

// Mengimplementasi interface car
type Car interface {
	Move() string
}

// Mendefinisikan struct
type Sedan struct{}

type SUV struct{}

type Sport struct{}

// Mengimplementasikan method `Move` untuk struct `sedan`
func (S Sedan) Move() string {
	return "Berjalan Sedan!"
}

// Mengimplementasikan method `Move` untuk struct `SUV`
func (S SUV) Move() string {
	return "Berjalan SUV!"
}

// Mengimplementasikan method `Move` untuk struct `Sport`
func (S Sport) Move() string {
	return "Berjalan Sport!!"
}

//method digunakan untuk mencari jari jari dari lingkaran

func (c circle) radius() float64 {
	return c.diameter / 2
}
func printArea(s shape) {
	fmt.Println(s.area())
}

func printAround(s shape) {
	fmt.Println(s.around())
}

func Interface() {
	var bangundatar shape

	bangundatar = square{width: 10.0}
	fmt.Println(bangundatar.area()) //100

	bangundatar = rectangle{width: 1.0, height: 1.0}
	fmt.Println(bangundatar.area()) //1

	bangundatar = circle{diameter: 7}
	fmt.Println(bangundatar.area())
	// fmt.Println(bangundatar.(circle).radius())

	var bangunDatar2 shape

	bangunDatar2 = circle{diameter: 14}
	printArea(bangunDatar2)

	var bangunDatar3 shape

	bangunDatar3 = rectangle{width: 2, height: 2}
	printArea(bangunDatar3)

	var bangunDatar4 shape

	bangunDatar4 = square{width: 4}
	printAround(bangunDatar4)

	//Execute interface Car
	cars := []Car{Sedan{}, SUV{}, Sport{}}
	for _, iteration := range cars {
		fmt.Println(iteration.Move())
	}

}
