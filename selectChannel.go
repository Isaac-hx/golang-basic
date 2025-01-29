package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getAverages(listNum []int, ch chan float64) {
	sum := 0
	for _, i := range listNum {
		sum += i
	}
	ch <- float64(sum) / float64(len(listNum))
}

func getMax(listNum []int, ch chan int) {
	var max int

	for _, i := range listNum {
		if i > max {
			max = i
		}
	}

	ch <- max
}

func SelectChannel() {
	// numbers := []int{3, 4, 4, 5, 6, 6, 7, 234, 5, 67, 7}
	// //Menjalankan kedua fungsi tersebut secara konkuren
	// ch1 := make(chan float64)
	// go getAverages(numbers, ch1)

	// ch2 := make(chan int)
	// go getMax(numbers, ch2)

	// //Selama menjalankan program secara konkuren,kode dibawah ini akan dilakukan eksekusi
	// for i := 0; i < 2; i++ {
	// 	//Pada setiap iterasi dari perulangan akan menunggu secara synchronous hingga salah satu ch1 dan ch2 telah mengirimkan data
	// 	select {
	// 	//case dimana average telah mengirim data dan diterima pada variabel avg
	// 	case avg := <-ch1:
	// 		fmt.Println("Average : ", avg)
	// 	//case dimana max telah mengirim data dan diterima oleh variabel max
	// 	case max := <-ch2:
	// 		fmt.Println("Max  : ", max)
	// 	}

	// }

	//soal
	//membuat dua sensor yang mengirimkan data dari goroutine ke goroutine utama dengan masing masing interval
	//suhu 700ms dan kelembapan 1s

	temperature := make(chan float64)
	// go func(ch chan float64) {
	// 	for {
	// 		temperature := 20.0 + rand.Float64()*(35.0-20.0)
	// 		ch <- temperature
	// 		time.Sleep(700 * time.Millisecond)
	// 	}

	// }(temperature)

	humidity := make(chan int)
	// go func(ch chan int) {
	// 	for {
	// 		humidity := rand.Intn(90-40) + 40
	// 		ch <- humidity
	// 		time.Sleep(1 * time.Second)
	// 	}

	// }(humidity)

	go temperatureSensor(temperature)
	go humiditySensor(humidity)

	for {
		select {
		case getTemperature := <-temperature:
			fmt.Println(fmt.Sprintf("Temperature Sensor : %v", getTemperature))
		case getHumidity := <-humidity:
			fmt.Println(fmt.Sprintf("Humidity Sensor : %v", getHumidity))
		case <-time.After(5 * time.Second):
			fmt.Println("Pengumpulan data selesai")
			return

		}
	}
}

func temperatureSensor(ch chan float64) {
	ticker := time.NewTicker(700 * time.Millisecond) // Membuat ticker dengan interval 700 ms
	defer ticker.Stop()                              // Pastikan ticker dihentikan setelah selesai

	for {
		select {
		case <-ticker.C:
			// Mengirimkan suhu acak antara 20 hingga 35
			temp := 20.0 + rand.Float64()*(35.0-20.0)
			ch <- temp
		}
	}
}

func humiditySensor(ch chan int) {
	ticker := time.NewTicker(1 * time.Second) // Membuat ticker dengan interval 1 detik
	defer ticker.Stop()                       // Pastikan ticker dihentikan setelah selesai

	for {
		select {
		case <-ticker.C:
			// Mengirimkan kelembapan acak antara 40 hingga 90%
			humidity := rand.Intn(51) + 40 // Bilangan bulat antara 40 hingga 90
			ch <- humidity
		}
	}
}
