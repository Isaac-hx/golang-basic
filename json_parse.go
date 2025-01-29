package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  uint
}

//Mendecode nilai json menjadi object/type valid di golang
//Contoh menggunakan map[string]interface{}

func decodeJson(data string) []map[string]interface{} {
	jsonData := []byte(data)
	var collectData []map[string]interface{}
	err := json.Unmarshal(jsonData, &collectData)
	if err != nil {
		log.Fatal(err.Error())
	}
	return collectData
}
func JsonParse() {
	//iniasilisasi data dump json
	jsonString := `[
    {"name": "Dimas Ananda Riyadi", "age": 20},
    {"name": "Laila Nurfadilah", "age": 22},
    {"name": "Andi Saputra", "age": 25},
    {"name": "Siti Aisyah", "age": 19},
    {"name": "Rizki Aditya", "age": 21},
    {"name": "Putri Kartika", "age": 24},
    {"name": "Bayu Pratama", "age": 23},
    {"name": "Hendra Wijaya", "age": 26},
    {"name": "Citra Melinda", "age": 20},
    {"name": "Agung Wahyudi", "age": 22}
]
`
	// 	jsonData := []byte(jsonString)

	// 	var collectData []User

	// 	err := json.Unmarshal(jsonData, &collectData)
	// 	if err != nil {
	// 		fmt.Println("Error : Gagal decode json!!", err.Error())
	// 		return
	// 	}
	collectData := decodeJson(jsonString)
	// fmt.Println("Nama : ", data.Name)
	// fmt.Println("Umur : ", data.Age)
	for _, data := range collectData {
		fmt.Println("Nama : ", data["name"])
		fmt.Println("Umur : ", data["age"])
	}
	extData := encodeJson()
	var dataUser []User
	err := json.Unmarshal(extData, &dataUser)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, data := range dataUser {
		fmt.Println(data.Name)
		fmt.Println(data.Age)
	}
	fmt.Println(dataUser)
}

func encodeJson() []byte {
	object := []map[string]interface{}{

		{"name": "Aris", "Age": 20},
		{"name": "Agus", "age": 19},
	}

	jsonData, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err.Error())
	}

	return jsonData
}
