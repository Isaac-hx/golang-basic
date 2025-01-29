package main

import "fmt"

func Interface2() {
	var listOfPerson []map[string]any

	listOfPerson = []map[string]any{
		{"Nama": "Dimas Ananda Riyadi", "Umur": 29},
		{"Nama": "Mia ahayu", "Umur": 19},
	}
	fmt.Println(listOfPerson)

	//initialization of object struct person
	var listPeopleCountry []map[string]any

	listPeopleCountry = []map[string]any{
		{"Nama": "Dimas ANanda riyadi", "Umur": 12},
		{"Nama": "Azizeh", "Umur": 18},
		{"Nama": "Juan Mayerel", "Umur": 30},
		{"Nama": "Findarita", "Umur": 100},
		{"Nama": "Agus", "Umur": 22},
	}

	for _, i := range listPeopleCountry {
		fmt.Println(i["Nama"], i["Umur"])
	}
	fmt.Println(listPeopleCountry)
}
