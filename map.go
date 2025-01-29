package main

import "fmt"

func Map() {
	var data map[string]string
	data2 := []map[string]string{} //deklarasi yang cukup simple untuk di implementasikan
	data2 = append(data2, map[string]string{"Nama": "Dimas"})
	data3 := []map[string][]string{}
	su := []map[string]string{}
	su = append(su, map[string]string{
		"Nama": "Dimas",
	},
		map[string]string{
			"Nama": "Reza",
		},
	)
	data3 = append(data3, map[string][]string{
		"FST": {"Teknik Informatika", "Sistem Informasi", "Management Informasi"},
	},
		map[string][]string{
			"FIP": {"Pend.Bahasa Inggris", "Pend.Eknomi", "Pend.Matematika", "Pend.IPA"},
		},
	)
	data["Nama"] = "Dimas Ananda Riyadi"
	addData(&data2, "Azzahra")
	addData(&data2, "Nabila")
	addData(&data2, "Untirta")
	printData(data2)
	updateData(data2, 1, "Azahra sangat cantik")
	deleteData(&data2, "Nabila")
	fmt.Println(su)
	fmt.Println(data3)
	fmt.Println(data3, "ini adalah data3")
	for row, data := range data3 {
		fmt.Println(row, data)
	}

	xs := map[string][]string{}
	xs["Dimas"] = []string{"tinggi", "ganteng", "pinter"}
	xs["Azzahra"] = []string{"tinggi", "cantik", "pinter"}
	for index, _ := range xs {
		fmt.Println(index)
		fmt.Println(xs[index])
	}
	fmt.Println(xs["Dimas"])
}

func addData(data *[]map[string]string, nama string) {
	*data = append(*data, map[string]string{"Nama": nama})
}

// fungsi ini digunakan untuk memprint semua isi data dari map
func printData(data []map[string]string) {
	for index, row := range data {
		fmt.Println(index+1, row["Nama"])
	}
}

// fungsi ini digunakan untuk menghapus data dari slice map
func deleteData(data *[]map[string]string, nama string) {
	for i := len(*data) - 1; i >= 0; i-- { // iterate in reverse to avoid index shifting issues
		if (*data)[i]["Nama"] == nama {
			*data = append((*data)[:i], (*data)[i+1:]...) // remove the element
		}
	}
}

// fungsi ini digunakan untuk mengupdate nilai data pada slice map
func updateData(data []map[string]string, idx int, nama string) {
	data[idx]["Nama"] = nama
}
