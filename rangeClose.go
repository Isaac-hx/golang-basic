package main

import (
	"fmt"
	"time"
)

// fungsi hanya untuk mengirim data ke channel
func sendMessages(ch chan<- string) {
	//mengirim data 20 kali ke dalam channel ch
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data ke - %d", i)

	}
	//Menutup channel untuk tidak digunakan kembali
	close(ch)
}

// fungsi hanya untuk membaca dari channel
func printMessages(ch <-chan string) {

	//melakukan perulangan sebanyak data yang ada pada channel
	for message := range ch {
		fmt.Println(message)
	}
}

func multiplyByTwo(listNumber []int, ch chan []int) {
	var listOutput []int
	for _, i := range listNumber {
		listOutput = append(listOutput, i*2)
	}
	ch <- listOutput
	close(ch)
}

func sumList(listNumber chan []int, ch chan int) {
	sum := 0
	for iterateListNumber := range listNumber {
		for _, dataListNumber := range iterateListNumber {
			sum += dataListNumber
		}
	}
	ch <- sum
}

func RangeClose() {

	listDataNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	multiplyDataNumber := make(chan []int)
	sumDataNumber := make(chan int)
	go multiplyByTwo(listDataNumber, multiplyDataNumber)
	go sumList(multiplyDataNumber, sumDataNumber)
	fmt.Println("Hasil penjumlahan", <-sumDataNumber)
	time.Sleep(2 * time.Second)

}
