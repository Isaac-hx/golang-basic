package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var ListofBooks []Book = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", TotalPage: 380, Publisher: "Addison-Wesley"},
	{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", TotalPage: 464, Publisher: "Prentice Hall"},
	{ID: 3, Title: "Introduction to Algorithms", Author: "Thomas H. Cormen", TotalPage: 1312, Publisher: "MIT Press"},
	{ID: 4, Title: "Design Patterns", Author: "Erich Gamma", TotalPage: 395, Publisher: "Addison-Wesley"},
	{ID: 5, Title: "You Don’t Know JS: Scope & Closures", Author: "Kyle Simpson", TotalPage: 98, Publisher: "O'Reilly Media"},
	{ID: 6, Title: "The Pragmatic Programmer", Author: "Andrew Hunt", TotalPage: 352, Publisher: "Addison-Wesley"},
	{ID: 7, Title: "Refactoring", Author: "Martin Fowler", TotalPage: 448, Publisher: "Addison-Wesley"},
	{ID: 8, Title: "Effective Java", Author: "Joshua Bloch", TotalPage: 412, Publisher: "Prentice Hall"},
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	//Log server
	defer LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	//Menentukan tipe response apa yang akan dikembalikan client
	w.Header().Set("Content-type", "application/json")

	//jika method yang digunakan adalah get
	if r.Method == "GET" {
		res, err := json.Marshal(ListofBooks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.Write(res)
		return
	}
	http.Error(w, "Books not found!", http.StatusBadRequest)
}

func getBooksById(w http.ResponseWriter, r *http.Request) {
	//Log server

	defer LogServer(fmt.Sprintf("%s - %s - %s", r.Host, r.Method, r.URL))

	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		for _, data := range ListofBooks {
			//jika id ditemukan
			if data.ID == id {
				//proses encode data menjadi json
				response, err := json.Marshal(data)
				//jika terjadi error pada encode data
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				//response balik ke client dari json
				w.Write(response)
				return

			}
		}
		http.Error(w, "Book not found!", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func Webservice() {
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/book", getBooksById)
	ListRoute("/books", "/book")
	fmt.Println("Service berjalan di localhost:8085/")

	http.ListenAndServe(":8085", nil)

}
