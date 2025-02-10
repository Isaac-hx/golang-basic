package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-basic/library"
	"net/http"
	"strconv"
	"strings"
)

func main() {

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

	//execute filter func findAzzahra
	fmt.Println(filter([]string{"azzahra", "ara", "ayu", "AZZAHRA"}, findAzzahra))

	//execute file pointer
	Pointer()

	//execute file struct
	Struct()

	//execute file method
	Method()

	//execute package init library
	fmt.Println(library.OSLauncher.OSName)
	fmt.Println(library.OSLauncher.KernelName)

	//execute file interface
	Interface()

	//execute file interface-2
	Interface2()

	//execute file reflect
	Reflect()

	// //execute file goroutine
	// Goroutine()

	// //execute file channel
	// Channel()

	// //execute file buffer
	// Buffer()

	// //execute file SelectChannel
	// SelectChannel()

	// //execute file rangeClose
	// RangeClose()

	// //execute file timeout
	// TimesOut()

	// //execute file deferExit
	// DeferExit()

	//execute file errors
	// GetErrors()

	//Menghitung berapa lama waktu eksekusi dari variabel start ke variabel duration
	// start := time.Now()
	// time.Sleep(5 * time.Second)
	// duration := time.Since(start)
	// fmt.Println("time elapsed in seconds : ", duration.Seconds())
	// fmt.Println("time elapsed in minutes : ", duration.Minutes())
	// fmt.Println("time elapsed in hours : ", duration.Hours())

	// //Menghitung selisih lama waktu eksekusi dari variabel t1 ke t2
	// t1 := time.Now()
	// time.Sleep(10 * time.Second)
	// t2 := time.Now()

	// timeGap := t2.Sub(t1)

	// fmt.Println("Selisih waktu dari t1 ke t2 : ", timeGap.Seconds())

	// //execute file regex
	// RegexP()

	// //execute file args
	// Flag()

	// execute file exec
	// Exec()

	//execute file webserver
	// WebServer()

	//execute file jsonParse
	// JsonParse()

	//execute file webservice
	// Webservice()

	//test
	// dataTest, err := fetchBook()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// unitTestGet(dataTest, ListofBooks)

	//execute file Wg Sync group
	// WaitGroupExec()

	//execute file wg sync mutex
	MutexExcute()

}

func unitTestGet(input []Book, unit []Book) {
	for iter, data := range input {
		if data == unit[iter] {
			fmt.Println(iter+1, "Test succes!!")
			continue
		}
		fmt.Println(iter+1, "Test Failed!!")
	}
}

func validateNumber(input string) error {
	convertToFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Error : input bukan angka desimal yang valid!")
	}
	convertToInt := int64(convertToFloat)
	convertToString := strconv.Itoa(int(convertToInt))
	fmt.Println("Angka desimal (Float64) : ", convertToFloat)
	fmt.Println("Angka integer (Int64) : ", convertToInt)
	fmt.Println("Angka string (String) : ", convertToString)
	return nil
}
func greet(setence1 string, setence2 []string) string {
	joinSetence2 := strings.Join(setence2, " ")
	return setence1 + " " + joinSetence2
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

// if word == azzahra
func findAzzahra(each string) string {
	if each == "azzahra" || each == "Azzahra" || each == "AZZAHRA" {
		return fmt.Sprintf("Ketemu!!, %s aku sayang kamu!", each)
	} else if each == "ara" {
		return fmt.Sprintf("Ketemu !!,%s apakah kamu mau jadi pacarku?", each)

	}
	return ""
}

var baseUrl = "http://localhost:8085"

func fetchBook() ([]Book, error) {
	var err error
	//Mendeklarasikan nilai dereference instance object http.client
	var client = &http.Client{}
	var data []Book

	//Memanggil fungsi http.NewRequest yang mengembalikan nilai pointer reference http.Request
	//Parameter pertama,berisikan tipe method yang akan digunakan seperti `POST` atau `GET`
	//Parameter kedua,adalah tujuan url yang ingin di request
	//Parameter ketiga,adalah form data request jika terdapat body
	req, err := http.NewRequest("GET", baseUrl+"/books", nil)
	if err != nil {
		return nil, err
	}

	//Memanggil method yang ada pada instance `client` dan mengeksekusi method Do
	//,dengan menyisipkan argumen req yang telah dibuat object sebelumnya pada variabel `req`
	//Method ini akan mengembalikan object http.Response
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//response data yang telah diambil perlu diclose setelah tidak dipakai.
	defer res.Body.Close()

	//data pada variabel `res` yang terdapat pada property body
	//Mendecode data menjadi data bertipe pointer `&data`
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
