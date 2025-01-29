package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

type mahasiswa struct {
	nim   uint
	umur  uint
	prodi string
}

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	TotalPage int    `json:"total_page"`
	Publisher string `json:"publisher"`
}

const path string = "/home/isaaachx/Documents/log.txt"

func LogServer(textLog string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	timeLog := fmt.Sprintf("[%s] %s", timestamp, textLog)
	WriteFileWithOS(path, timeLog)
	log.Println(textLog)
}

func index(w http.ResponseWriter, r *http.Request) {
	defer LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	fmt.Fprintln(w, "<h1>Halo Apa Kabar !!!</h1>")

}
func welcome(w http.ResponseWriter, r *http.Request) {
	defer LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	fmt.Println(r.Method)
	fmt.Fprintln(w, "Damn")

}

func controllerGetBookList(w http.ResponseWriter, r *http.Request) {
	defer LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	var listBooks []Book
	listBooks = []Book{
		{Title: "Sapiens", Author: "Yuval Noah Harari", TotalPage: 527, Publisher: "KPG"},
		{Title: "The Little Book of Sideways Markets", Author: "Vitaly", TotalPage: 256, Publisher: "ELEK Media"},
		{Title: "Designing Data-intensive Application", Author: "Martin Kleppman", TotalPage: 562, Publisher: "Kompas Gramedia"},
	}

	for _, data := range listBooks {
		fmt.Fprintln(w, fmt.Sprintf(
			`<div>
				<h1>%s</h1>
				<h4>%s</h4>
				<p>%d</p>
				<p>%s</p>	
			</div> <br>`, strings.ToUpper(data.Title), data.Author, data.TotalPage, data.Publisher,
		))
	}
}

func WebServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

		fmt.Fprintln(w, "Halo!!")
		data := map[string]string{
			"name":    "Dimas Ananda Riyadi",
			"message": "Hello world",
		}
		//Mengembalikan file parse ke template.html
		t, err := template.ParseFiles("template.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		var collectMahasiswa map[string]mahasiswa
		collectMahasiswa = map[string]mahasiswa{
			"Isaac": {
				prodi: "Teknik Informatika", umur: 20, nim: 3922250001,
			},
			"Nabila": {
				prodi: "Kedokteran", umur: 20, nim: 392225010,
			},
			"Aca": {
				prodi: "Hukum", umur: 22, nim: 2910035403,
			},
		}

		for indexNama, data := range collectMahasiswa {
			fmt.Fprintln(w, fmt.Sprintf(
				`<ul>
				<li><h1>Nama : %s</h1></li>
				<h2>Nim : %d</h2>
				<h2>Prodi : %s</h2>
				<h2>Umur : %d</h2>
				</ul>`, indexNama, data.nim, data.prodi, data.umur,
			))
		}
	})

	//fungsi yang digunakan untuk menghandle routing aplikasi web
	http.HandleFunc("/index", index)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/books", controllerGetBookList)
	fmt.Println("Start web server at http://localhost:8082/")
	listRouting := []string{
		"/index", "/welcome", "/mahasiswa", "/books",
	}
	ListRoute(listRouting...)
	http.ListenAndServe(":8082", nil)

}

func ListRoute(dataRoute ...string) {
	fmt.Println("List route registered : ")
	for _, data := range dataRoute {
		fmt.Println(data)
	}
}
