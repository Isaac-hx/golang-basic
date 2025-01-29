package main

import (
	"fmt"
	"os"
)

func sayHello() {
	fmt.Println("Program ini berjalan")
	defer fmt.Println("Eksekusi perintah ini akan selalu setelah semua blok fungsi berakhir")
	fmt.Println("Program ini proses")
	fmt.Println("Program ini berakhir")
}

func itsEvenOrdOdd(number int) {
	fmt.Println("Ini adalah program genap atau ganjil")
	defer fmt.Println("Eksekusi perintah ini akan selalu setelah semua blok fungsi berakhir")

	if number%2 == 0 {
		fmt.Println("Angka adalah genap,dan program akan seluruh akan berhenti")
		os.Exit(1)
	}

	fmt.Println("Angka adalah ganjil,program akan segera berakhir!!")

}

func DeferExit() {
	sayHello()
	itsEvenOrdOdd(2)
}
