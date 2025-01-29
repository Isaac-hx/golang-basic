package main

import "fmt"

// fungsi menjumlahkan semua bilangan yang ada pada slice int
// dengan beberapa paarameter yaitu s[]int dan c chan int
// parameter c digunakan untuk mengirimkan hasil penjumlahan ke goroutine lain
func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum //mengirimkan hasil penjumlahan ke dalam channel parameter c disebut dengan sender
}
func Channel() {
	s := []int{7, 2, 1, 4, 5, 6}
	d := []int{1, 2, 3, 4, 5, 6}

	//membuat channel ditampung kedalam variabel c dengan tipe data int
	c := make(chan int)

	//eksekusi fungsi menggunakan goroutine
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum(d[:len(d)/2], c)
	//menerima data dari sender dengan menggunakan simbol <- yang disebut receiver
	fmt.Println(<-c, <-c)
	//x, y := <-c, <-c
	x := <-c

	//karena kedua goroutine berjalan secara pararel,urutan nilai x dan y bisa berbeda,tergantung waktu eksekusi selesainya
	fmt.Println(x)

	messages := make(chan string)

	sayHello := func(who string) {
		message := fmt.Sprintf("Hello,%s", who)
		messages <- message
	}

	go sayHello("Dimas")
	go sayHello("Nabilla")
	go sayHello("Azzahra")
	message1 := <-messages
	message2 := <-messages
	message3 := <-messages
	fmt.Println(message1, message2, message3)

}
