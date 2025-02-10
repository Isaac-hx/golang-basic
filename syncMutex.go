// Sync mutex merupakan thread safe sehingga aman digunakan untuk mengeksekusi goroutine
// Konsep mutual exclusion memastikan bahwa hanya satu proses yang bisa mengakses suatu resources,atau biasa disebut
// dengan race condition
package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	//mengembed object mutex agar bisa digunakan menerapkan method mutex pada method counter
	sync.Mutex
	val int
}

func (c *counter) addCounter() {
	//Rangkaian lock dan unlock ini disebut juga dengan	proses mutual exclusion
	//Menandakan bahwa semua operasi setelah baris ini akan dikunci dan menjadi sync
	c.Lock()
	c.val++ //critical section
	//Menandakan bahwa operasi yang sudah dilakukan sebelumnya sudah dibuka,dan kembali bisa mengakses properti

	c.Unlock()

}

func (c *counter) showValue() int {
	return c.val
}

func MutexExcute() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var count counter

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//Menjalankan goroutine sebanyak 1000 kali
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				//Menambah counter dengan memanggil addCounter.
				count.addCounter()

			}
		}()
	}
	wg.Wait()
	fmt.Println(count.showValue())
}

//Question
//1.Bagaimana cara melakukan mutual exclusion tanpa mengembed object?
//2.Apakah kita perlu melakukan locking dua arah di database level dan di service level?
//3.Kasus apa saja yang bisa terjadi race conditio pada realcase ?

//Answer
//1.Bisa menggunakan sebagai variabel global dan akes method lock sama seperti object lainya.
//2.Locking pada database diperlukan untuk memastikan konsistensi data.Locking pada tingkat service digunakan
//untuk mengelola resource pada tingkat aplikasi,jika locking pada database dibutuhkan saat konsistensi data dan transparansi
//agar logika operasi aplikasi dapat lebih sederhana,sedangkan untuk pada tingkat service digunakan untuk caching dan
//Distribusi beban,dan untuk kasus yang tepat ialah optimalkan satu skema locking antara database atau service level
//3.Transfer uang rekening yang melibatkan kejadian race condiion,atau mengupdate suatu stock.
