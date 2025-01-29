package main

import "fmt"

// deklarasi struct tanpa inisialisasi object
type Student struct {
	name string
	age  uint
}

type Class struct {
	classGrade string
	age        uint
	Student
}

type Animal struct {
	species string
	name    string
}

// anonymous struct atau bisa disebut struct yang dideklarasi dan langsung dan langsung di inisialisasi
var manusia = struct {
	name   string
	age    uint
	gender string
}{}

func Struct() {
	var student_1 Student
	student_1.name = "Dimas Ananda Riyadi"
	student_1.age = 12
	var student_2 = Student{name: "Dimas", age: 19}
	student_3 := Student{name: "Dunas", age: 120}
	fmt.Println(student_3)
	fmt.Println(student_2)
	fmt.Println(student_1)
	fmt.Println(student_1.name)

	//Pointer to struct
	var student_dereference *Student = &student_1
	fmt.Println("Ini sebelum nilai pointer diubah: ", student_dereference)
	student_dereference.name = "Azzahra"
	fmt.Println("Ini sesudah nilai pointer diubah : ", student_dereference, student_1)

	//initiated var to struct Class
	// student_4_data := Student{name: "Dimas Ananda Supriyadi", age: 1290}
	// student_4 := Class{classGrade: "X", Student: student_4_data}
	// fmt.Println(student_4)

	var student_4 Class
	student_4.name = "Azzahra cantik"
	student_4.age = 120
	student_4.classGrade = "XI"
	student_4.Student.age = 110
	fmt.Println("Embedded struct : ", student_4)

	//Anonymous struct
	student_5 := struct {
		nisn uint
		Student
	}{}
	student_5.nisn = 3922250001
	student_5.name = "Dimas Sayang Azzahra"
	student_5.age = 100
	fmt.Println(student_5)

	//slice & struct
	allStudents := []Class{
		{classGrade: "XII", age: 100, Student: Student{name: "Dimas Ananda Riyadi", age: 120}},
		{classGrade: "XII", age: 90, Student: Student{name: "Azzahra", age: 100}},
	}
	allStudents = append(allStudents, Class{classGrade: "XIII", age: 200, Student: Student{name: "Nabila", age: 100}})

	fmt.Println(allStudents)

	for _, data := range allStudents {
		fmt.Println()

		fmt.Println(fmt.Sprintf("Nama : %s \nKelas : %s \nUmur Palsu : %d \nUmur asli : %d", data.Student.name, data.classGrade, data.age, data.Student.age))
	}

	//inisialisasi object hewan_1

	hewan_1 := Animal{
		species: "Panthera leo",
		name:    "Singa/Lion",
	}
	fmt.Println(hewan_1)
	manusia.name = "Dimas Anando Riyado"
	manusia.age = 100
	manusia.gender = "Perempuan"
	fmt.Println(manusia)

}
