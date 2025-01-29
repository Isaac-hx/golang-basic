package main

import (
	"fmt"
	"time"
)

func Buffer() {
	// messages := make(chan int, 3)

	// // Goroutine penerima, tetapi proses menerima data lambat
	// go func() {
	// 	for msg := range messages {
	// 		fmt.Println("Receive data:", msg)
	// 		time.Sleep(2 * time.Second) // Simulasi proses lambat
	// 	}
	// }()

	// // Pengirim mengirimkan lebih banyak data daripada kapasitas buffer
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Sending data:", i)
	// 	messages <- i // Akan terblokir saat buffer penuh
	// 	fmt.Println("Data sent:", i)
	// }

	// // Memberi waktu untuk goroutine penerima menyelesaikan pekerjaannya
	// fmt.Println("All data sent, waiting for receiver...")
	// time.Sleep(15 * time.Second)
	// close(messages) // Menutup channel setelah selesai

	//soal 1
	//Membuat maksimum channel dengan maksimum pelanggan  5
	antrian := make(chan int, 5)

	// Goroutine pelanggan yang mengirim data ke channel setiap 500ms
	go func() {
		for i := 1; i <= 10; i++ {
			// Kirim data pelanggan ke channel
			select {
			case antrian <- i:
				fmt.Printf("Pelanggan %d masuk antrian.\n", i)
			default:
				// Jika antrian penuh, tidak bisa mengirim, pelanggan tetap terblokir
				fmt.Printf("Antrian penuh, pelanggan %d harus menunggu.\n", i)
			}
			time.Sleep(500 * time.Millisecond) // Delay 500ms antara pelanggan
		}
		close(antrian) // Menutup channel setelah semua pelanggan dikirim
	}()

	// Goroutine pelayan yang memproses pelanggan dari channel
	go func() {
		for pelanggan := range antrian {
			// Proses pelanggan
			fmt.Printf("Pelayan sedang melayani pelanggan %d...\n", pelanggan)
			time.Sleep(1 * time.Second) // Pemrosesan 1 detik per pelanggan
		}
	}()

	// Menunggu sampai semua pelanggan dilayani
	// time.Sleep(15 * time.Second) // Tunggu selama 15 detik untuk memastikan semua pelanggan dilayani

	// Menampilkan pesan setelah semua pelanggan dilayani
	fmt.Println("Semua pelanggan telah dilayani.")

}
