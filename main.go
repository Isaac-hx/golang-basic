package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var data map[string]string
	data = map[string]string{}
	fmt.Println("Hello World !!")
	data2 := []map[string]string{} //deklarasi yang cukup simple untuk di implementasikan
	data2 = append(data2, map[string]string{"Nama": "Dimas"})
	data3 := []map[string][]string{}
	su := []map[string]string{}
	su = append(su, map[string]string{
		"Nama": "Dimas",
	},
		map[string]string{
			"Nama": "Reza",
		},
	)
	data3 = append(data3, map[string][]string{
		"FST": {"Teknik Informatika", "Sistem Informasi", "Management Informasi"},
	},
		map[string][]string{
			"FIP": {"Pend.Bahasa Inggris", "Pend.Eknomi", "Pend.Matematika", "Pend.IPA"},
		},
	)

	data["Nama"] = "Dimas Ananda Riyadi"
	addData(&data2, "Azzahra")
	addData(&data2, "Nabila")
	addData(&data2, "Untirta")
	printData(data2)
	updateData(data2, 1, "Azahra sangat cantik")
	deleteData(&data2, "Nabila")
	fmt.Println(su)
	fmt.Println(data3)
	fmt.Println(data3, "ini adalah data3")
	for row, data := range data3 {
		fmt.Println(row, data)
	}

	xs := map[string][]string{}
	xs["Dimas"] = []string{"tinggi", "ganteng", "pinter"}
	xs["Azzahra"] = []string{"tinggi", "cantik", "pinter"}
	for index, _ := range xs {
		fmt.Println(index)
		fmt.Println(xs[index])
	}
	fmt.Println(xs["Dimas"])

	setence1 := "Hello"
	setence2 := []string{"Sir", "Ananda"}

	fmt.Println(greet(setence1, setence2))
	fmt.Println(divideNumber(2, 0))

	//wordStrings
	setence3 := []string{"Azzahra", "Kamu", "Sangat", "Cantik"}
	fmt.Println(wordStrings("Halo", setence3...))

	//closure function/anonymous function digunakan
	getOddOrEven := func(n ...int) ([]int, []int) {
		var odd, even []int
		for _, iter := range n {
			if iter%2 == 0 {
				even = append(even, iter)
				continue
			}
			odd = append(odd, iter)
		}
		return odd, even
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd, even := getOddOrEven(numbers...)
	fmt.Println("Ganjil : ", odd)
	fmt.Println("Genap : ", even)

	//Saya ingin tau berapa banyak jumlah bilangan yang kurang dari 0 dan data semua bilangan yang kurang dari 0

	findNumberLessThenN := func(limit int, n ...int) (int, func() []int) {
		var listLessThenLimit []int
		for _, iter := range n {
			if iter < limit {
				listLessThenLimit = append(listLessThenLimit, iter)
			}
		}
		//fungsi sebagai nilai kembalian
		return len(listLessThenLimit), func() []int {
			return listLessThenLimit

		}
	}
	dataLessThen1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 0, -1, -2, -3, -4, -5}
	lengthData, listData := findNumberLessThenN(1, dataLessThen1...)
	fmt.Println("Banyak angka kurang dari 1 : ", lengthData)
	fmt.Println("List angka kurang dari 1 : ", listData())

	//execute fungsi filtered data
	fmt.Println(filteredWordMoreThan5("azzahra"))
	fmt.Println(filteredWordContainsO("DimOs"))

	//execute filter funciton
	bunchOfWord := []string{"dimas", "onondo", "riyadi"}
	fmt.Println(filter(bunchOfWord, filteredWordMoreThan5))
	fmt.Println(filter(bunchOfWord, filteredWordContainsO))
}

func greet(setence1 string, setence2 []string) string {
	joinSetence2 := strings.Join(setence2, " ")
	return setence1 + " " + joinSetence2
}

// fungsi ini digunakan untuk menambah data kedalam map slice
func addData(data *[]map[string]string, nama string) {
	*data = append(*data, map[string]string{"Nama": nama})
}

// fungsi ini digunakan untuk memprint semua isi data dari map
func printData(data []map[string]string) {
	for index, row := range data {
		fmt.Println(index+1, row["Nama"])
	}
}

// fungsi ini digunakan untuk menghapus data dari slice map
func deleteData(data *[]map[string]string, nama string) {
	for i := len(*data) - 1; i >= 0; i-- { // iterate in reverse to avoid index shifting issues
		if (*data)[i]["Nama"] == nama {
			*data = append((*data)[:i], (*data)[i+1:]...) // remove the element
		}
	}
}

// fungsi ini digunakan untuk mengupdate nilai data pada slice map
func updateData(data []map[string]string, idx int, nama string) {
	data[idx]["Nama"] = nama
}

func divideNumber(m, n int) (int, bool) {
	if n == 0 {
		return 0, true
	}
	return m / n, false
}

// Mechanic physic
// Speed in seconds
func velocity(distance int, time int) (int, error) {
	if time == 0 {
		return 0, errors.New("Error can't divide by zero")
	}
	return distance / time, nil
}

// fungsi yang digunakan untuk membuat semua huruf menjadi kapital
func wordStrings(greetSetence string, n ...string) string {
	var data []string
	for _, itr := range n {

		data = append(data, strings.ToUpper(itr))
	}

	return fmt.Sprintf("%s, %s", greetSetence, strings.Join(data, " "))

}

// Saya mempunyai sebuah list kata kata filter lah kata kata tersebut,jika kata kata itu lebih dari 5 maka buatlah kata tersebut kapital dan buatlah fungsi filter jika terdapat o maka ubah o menjadi a dan semua datanya tidak ilang
func filter(data []string, callback func(each string) string) []string {

	var filteredData []string

	for _, each := range data {

		filteredData = append(filteredData, callback(each))
	}
	return filteredData
}

// if word more than 5
func filteredWordMoreThan5(each string) string {
	if len(each) > 5 {
		capitalizeWord := strings.ToUpper(string([]rune(each)[0]))
		slicingWord := strings.ToLower(each[1:])
		return capitalizeWord + slicingWord
	}
	return each
}

// if word contains o change to a
func filteredWordContainsO(each string) string {
	var newWord string

	if strings.Contains(each, "o") == true {
		newWord = strings.Replace(each, "o", "a", len(each))
		return newWord

	} else if strings.Contains(each, "O") == true {
		newWord = strings.Replace(each, "O", "a", len(each))
		return newWord
	}

	return each
}
