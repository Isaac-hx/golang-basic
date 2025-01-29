package library

import (
	"fmt"
	"math/rand"
	"time"
)

var OSLauncher = struct {
	OSName     string
	KernelName string
}{}

func generateRandomString() {
	//Membuat object randomizer yang dapat digunakan untuk melakukan randomier number
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	fmt.Println(letters)
	randomString := func(length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[randomizer.Intn(len(letters))]
		}
		return string(b)
	}(5)

	fmt.Println("Sesudah di parse ke string", randomString)
}

func init() {
	OSLauncher.OSName = "Linux"
	OSLauncher.KernelName = "Ubuntu-x441"
	generateRandomString()

	//kita mendefiniskan object waktu dengan menggunakan function now
	time1 := time.Now()
	fmt.Println("Waktu Sekarang adalah", time1)
	fmt.Println("Tahun sekarang adalah : ", time1.Year())
	fmt.Println("Bulan sekarang adalah : ", time1.Month())
	fmt.Println("Waktu sekarang adalah : ", time1.Hour())
	//mendefinisikan tanggal dan waktu sendiri dengan menggunakan function Date
	date1 := time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Println("Tanggal sekarang", date1)
	//parsing data dari string ke date time

	//parsing string ke object time.Time dengan layoutFormat
	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	date, _ := time.Parse(layoutFormat, value)
	fmt.Println(date.String())
	//parsing string ke object time.Time dengan layout format yang berbeda
	layoutFormat2 := "02/01/2006 MST"
	value2 := "10/10/2010 WIB"
	date2, _ := time.Parse(layoutFormat2, value2)
	fmt.Println(date2.String())
	//parsing string ke object time.Time dengan menggunakan predefined format layout
	date3, _ := time.Parse(time.RFC822, "10 Sep 15 10:00 WIB")
	fmt.Println("Date ke - 3", date3.String())

	//format data object time.Time ke string
	date4, _ := time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	stringDate4 := date4.Format("Monday 02,January 2006 15:04 MST")
	fmt.Println("Stringdate", stringDate4)
	fmt.Println("Package library---> BELAJAR GOLANG - RESERT >>> imported")

}
