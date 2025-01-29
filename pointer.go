package main

import "fmt"

func Pointer() {
	number1 := 10
	number2 := &number1

	var number3 int = 10
	var number4 *int = &number3 //tanda asterix '*' selalu diletakan di depan tipe variabel

	fmt.Println("Nilai dari deference Number 1 : ", number2)
	fmt.Println("Nilai dari deference Number 4 : ", number4)
	fmt.Println("Nilai dari reference Number 2 : ", *number2)
	fmt.Println("Nilai dari reference Number 4 : ", *number4)

	//before change value
	fmt.Println("Number 2 & 1 before change value : ", number1, *number2)
	fmt.Println("Number 3 & 4 before change value : ", number3, *number4)

	*number2 = 20
	*number4 = 30
	fmt.Println("Number 2 & 1 after change value : ", number1, *number2)
	fmt.Println("Number 3 & 4 after change value : ", number3, *number4)

	//create function change value within pointer method
	data := "dimas"
	fmt.Println(changeValuePointer(&data))
	fmt.Println(data)
}
func changeValuePointer(data *string) string {
	if *data == "dimas" {
		*data = "nilai pointer terhadap dimas diubah"
		return "data berhasil diubah"

	}
	return "tidak ada perubahan data"
}
