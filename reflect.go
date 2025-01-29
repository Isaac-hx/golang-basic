package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age  int
}

// method getPropertyInfo
func (s *student) getPropertyInfo() {
	// Ambil nilai refleksi dari s
	var reflectValue = reflect.ValueOf(s)

	// Periksa jika nilai adalah pointer
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		fmt.Println("Variabel ini adalah pointer", reflectValue)
	}

	// Ambil tipe refleksi
	var reflectType = reflectValue.Type()
	fmt.Println("Nilai dari ", reflectType.Field(0))

	// Iterasi setiap field
	for i := 0; i < reflectValue.NumField(); i++ {
		// Cetak informasi field
		fmt.Println(reflectType.Field(i))
		fmt.Println("Nama : ", reflectType.Field(i).Name)
		fmt.Println("Tipe data : ", reflectType.Field(i).Type)

		// Cetak nilai field
		fmt.Println("Nilai : ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

// method setname student
func (s *student) SetName(name string) {
	s.Name = name
}

// function that will accept any argument pass of all type data
func printOutTheStatement(statement reflect.Value) {
	fmt.Println(statement.Interface())
}

// check if integer
func checkNumberIsInteger(number reflect.Value) bool {

	if number.Kind() == reflect.Int {
		return true
	}
	return false

}

// check type data
func checkTypeData(data reflect.Value) {
	fmt.Println(data.Type())
}

func Reflect() {
	number := 23
	string1 := "Dimas"
	reflectValueString1 := reflect.ValueOf(string1)
	reflectValue := reflect.ValueOf(number)
	reflectType := reflect.TypeOf(string1)
	fmt.Println(reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Statement ini dijalankan")
	}

	// printOutTheStatement(reflectValue)
	fmt.Println(checkNumberIsInteger(reflectValue))
	checkTypeData(reflectValue)
	checkTypeData(reflectValueString1)
	fmt.Println("Hasil dari trype of", reflectType)
	fmt.Println("Hasil dari valueof", reflectValueString1)

	//run method getPropertyInfo
	siswa1 := student{Name: "Dimas Ananda Riyadi", Age: 20}
	siswa1.getPropertyInfo()

	//run method setName
	siswa2 := &student{Name: "Azzahara", Age: 21}

	var reflectValue2 = reflect.ValueOf(siswa2)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("cantik"),
	})

	fmt.Println(siswa2.Name)
}
