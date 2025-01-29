package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetErrors() {
	//Membuat fungsi catch dieksekusi terakhir
	//defer catch()
	var input string
	fmt.Println("Masukan nomer yang ingin dikonversi!!! ")
	fmt.Scanln(&input)

	//Akan mengembalikan dua nilai error dan hasil dari konversi string to integer
	numberConv, err := strconv.Atoi(input)
	//kondisi dimana jika variabel error terdapat kesalahan,maka statement ini akan dieksekusi
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Terjadi kesalahan dalam konversi angka!!")
		os.Exit(1)
	}
	//blok kode jika hasil konversi berhasil.
	fmt.Println("Angka berhasil dikonversiii nilainya sama dengan : ", numberConv)

	// var name string
	// fmt.Print("Type your name : ")
	// fmt.Scanln(&name)

	// valid, err := validate(name)

	// if valid {
	// 	fmt.Println("Hello", name)
	// } else {
	// 	panic(err.Error())
	// 	fmt.Println("Endline...")
	// }

	// data := []string{"Superman", "Aquaman", "Wonder woman"}

	// for _, each := range data {
	// 	func() {
	// 		defer func() {
	// 			r := recover()
	// 			if r != nil {
	// 				fmt.Println("Panic occured looping", each, "Message")
	// 			} else {
	// 				fmt.Println("Applicatrion running perfectly")
	// 			}
	// 		}()
	// 		panic("Some err")
	// 	}()

	// }

	//divisionByTwoNumbers(2, 0)
	go goroutine1()
	go goroutine2()
	go goroutine3()

	time.Sleep(2 * time.Second)
}

func validate(input string) (bool, error) {
	//Jika kondisi dimana input user kosong
	if strings.TrimSpace(input) == "" {
		//Membuat kustom errror
		return false, errors.New("Input Cannot be empty")
	}
	return true, nil

}

func catch() {
	//Menjalankan fungsi recover untuk mengambil alih goroutine jika mengalami panic
	r := recover()
	if r != nil {
		fmt.Println("Error occured : Division by zero is not allowed")
	}
}

func divisionByTwoNumbers(a, b int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Zero divisiozn not allowed")
		}
	}()
	result := a / b
	fmt.Println("Result of division : ", result)

}

func goroutine1() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Goroutine 1 : Panic occured : Random failure")
		}
	}()
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberRandom := randomizer.Int()%10 + 1
	if numberRandom <= 2 {
		panic("Panic occured : Random Failure")
	}
	fmt.Println("Goroutine 1 : No panic,executed completed", numberRandom)
}
func goroutine2() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Goroutine 2 : Panic occured : Random failure")
		}
	}()
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberRandom := randomizer.Int()%10 + 1
	if numberRandom <= 2 {
		panic("Panic occured : Random Failure")
	}
	fmt.Println("Goroutine 2 : No panic,executed completed", numberRandom)
}

func goroutine3() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Goroutine 3 : Panic occured : Random failure")
		}
	}()
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberRandom := randomizer.Int()%10 + 1
	if numberRandom <= 2 {
		panic("Panic occured : Random Failure")
	}
	fmt.Println("Goroutine 3 : No panic,executed completed", numberRandom)
}
