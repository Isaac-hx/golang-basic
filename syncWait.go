//Sync wait group merupakan salah satu tipe yang thread safe,variabel bertipe waitgroup aman dari potensi race condition,serta aman digunakan untuk banyak goroutine secara pararel.

package main

import (
	"fmt"
	"sync"
)

// Fungsi ini akan dijalankan sebagai go routine
// Variabel wg digunakan untuk sebagai kontrol dari sinkronisasi goroutine yang dijalankan.
func doPrint(wg *sync.WaitGroup, message string) {
	//Method done digunakan untuk memberikan informasi bahwa fungsi doPrint sudah selesai dijalankan
	fmt.Println("Pemrosesan data dimulai!!!")
	defer wg.Done()

	fmt.Println(message)
	fmt.Println("Pemrosesan data selesai!")
}
func WaitGroupExec() {
	//deklarasi variabel wg sebagai sync.WaitGroup
	var wg sync.WaitGroup

	//Pengulangan untuk memanggil fungsi doPrint sebagai goroutine sebanyak 5 kali iterasi
	//Pada setiap perulangan/iterasi kode tersebut memanggil method wg dan menginformasikan bahwa terjadi pemanggilan/penambahan goroutine sebanyak 1 proses (karena perulangan berjalan 5 kali,maka dari itu wg akan mengetahui bahwa proses goroutine sedang berjalan sebanyak 5kali)
	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("data %d", i)
		//Proses goroutine ditambah 1 setiap perulangan
		wg.Add(1)
		go doPrint(&wg, data)
	}

	//Method ini bersifat blocking,proses program tidak akan diteruskan kebaris selanjutnya,sebelum perulangan diatas selesai,dalam kasus ini terdapat 5 buah perulangan/goroutine yang sedang dijalankan.Pemanggilan wait juga harus sama dengan pemanggilan add,jika proses goroutine dipanggil sebanyak n,maka wait harus dipanggil sebanyak n juga
	wg.Wait()
	fmt.Println("Selesai program waitgroup!")
}

//Question
//1.Apa itu thread safe?
//2.Apa yang terjadi jika saya tidak menambahkan salah satu tiga method diatas?
//3.Contoh implementasinya pada real case?

//Answer
//1.Istilah dalam bahasa pemrograman yang menggambarkan kode atau fungsi yang dapat dijalankan secara bersamaan,tanpa perilaku tak terduga seperti race-condition,atau kerusakan data.Dalam aplikasi multithreading, beberapa thread dapat mengakses dan memodifikasi data yang sama secara bersamaan.
//2.- Saat method wg.Done dihilangkan program tidak berhenti atau terus berjalan sampai di interuppt,alasannya karena kembali ke pengertiannya diatas!
//- terjadi panic dan program langsung berakhir selesai,lalu pesan error yang muncul adalah panic:sync:negative WaitGroup Counter,alasan panic terjadi karena wg.Done lebih banyak dari wg.Add(),sehingga counter menjadi negatif dan tidak valid.Tipsnya adalah selalu pastikan wg.Add() dijalankan di thread utama sebelum goroutine dimulai.
//- Blok kode selanjutnya langsung dieksekusi tanpa menunggu semua goroutine yang berjalan selesai,alasannya program tidak akan menunggu goroutine selesai tetapi berjalan secara ansikron Tips pastikan selalu menggunakan wg.Wait(),agar proses pekerjaan pada goroutine tidak hilang.
//3.Memproses data secara pararel.Misalkan kita memiliki sejumlah data yang perlu diproses secara paralel untuk meningkatkan efisiensi. Kita dapat menggunakan goroutine untuk memproses setiap data dan sync.WaitGroup untuk menunggu hingga semua goroutine selesai.
