package main

import (
	"fmt"
	"strings"
)

type Laptop struct {
	brand     string
	brandType string
}

type Smartphone struct {
	brand string
	imei  int
}

func getGreetBrand(brand, greet string) string {
	return fmt.Sprintf("Object brand is indicate %s!!!,%s", brand, greet)

}

func (s Smartphone) getBrand() string {
	switch strings.ToLower(s.brand) {
	case "xiomi":
		return getGreetBrand(s.brand, "You're  poor user")
	case "asus":
		return getGreetBrand(s.brand, "You're middle user")
	case "samsung":
		return getGreetBrand(s.brand, "You're high user")
	case "iphone":
		return getGreetBrand(s.brand, "You're exclusive user")
	}
	return "You're brand not registered"
}

func (l Laptop) getBrand() string {
	return l.brand
}
func (l Laptop) getBrandType() string {
	return l.brandType
}

func (l *Laptop) setBrandType(newBrandType string) {
	l.brandType = newBrandType
	fmt.Println(l.brandType)
}
func Method() {
	lp1 := Laptop{
		brand:     "Asus",
		brandType: "X441UA",
	}
	fmt.Println(lp1.getBrand())
	fmt.Println(lp1.getBrandType())
	lp1.setBrandType("Vivobook yx86")
	fmt.Println(lp1.brandType)

	sm1 := Smartphone{
		brand: "Samsung",
		imei:  3922250001,
	}

	fmt.Println(sm1.getBrand())
	sm2 := Smartphone{
		brand: "Xiomi",
		imei:  10040123,
	}
	fmt.Println(sm2.getBrand())

	sm3 := Smartphone{
		brand: "Iphone",
		imei:  12391293219,
	}
	fmt.Println(sm3.getBrand())
}
