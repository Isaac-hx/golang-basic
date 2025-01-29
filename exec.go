package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// membuat sebuah file dimas.txt jika tidak ada,dan return nilai error jika sudah ada
func Exec() {
	// //function untuk mereturn/mengembalikan nilai path pada pwd
	// exec.Command("cd", "/home")
	// output1 := exec.Command("pwd")
	// //method yang digunakan untuk mengembalikan hasil eksekusi command
	// resultOutput1, _ := output1.Output()
	// fmt.Println("Hasil function command : ", output1)
	// fmt.Println("Hasil dari methode Output : ", string(resultOutput1))
	// fmt.Println(runtime.GOOS)
	// createFiledimas()
	// path := "/home/isaaachx/Documents/coba.txt"
	// createFileWithOS(path)
	// writeFileWithExec(path, "Dimas ananda riyadi")
	// writeFileWithOS(path)
	// readFileWithExec(path)
	// readFileWithOS(path)
	// deleteFileWithExec(path)
	// deleteFileWithOS(path)
}

func createFileWithOS(path string) {
	if _, err := os.Stat(path); err != nil {
		file, err2 := os.Create(path)
		if err2 != nil {
			log.Fatal(err2.Error())
		}
		log.Println("Info : File berhasil dibuat!!!")
		defer file.Close()
		return
	} else {
		log.Fatal("Error : File dengan nama yang sama sudah digunakan!!!")
	}
}

func WriteFileWithOS(path, text string) {
	//Membuka file atau membuat file jika tidak ditemukan
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal("Error : Tidak ditemukan file log.txt", err.Error())
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", text))
	if err != nil {
		log.Fatal("Error : Kesalahan dalam menulis log!", err.Error())
	}

	err = file.Sync()
	if err != nil {
		log.Fatal("Error : Kesalahan dalam menyimpan file!", err.Error())
	}

}

func writeFileWithExec(path, text string) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("Ini dijalankan")
		cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo "%s" >> %s`, text, path))
		if _, err := cmd.Output(); err != nil {
			log.Fatal("Error : gagal menulis teks!!")
		}
		log.Println("Berhasil menambah teks!!!")
		return

	} else {
		log.Fatal("Error : Gagal menulis file!!!")
	}

}
func createFileWithExec() {

	if _, err := os.Stat("dimas.txt"); err == nil {
		cmd := exec.Command("sh", "-c", `echo "File : already exist" >> dimas.txt`)
		if _, err := cmd.Output(); err != nil {
			log.Fatal("Gagal menambahkan file", err)
		}
		log.Println("File dimas.txt sudah ada!!")
		return

	} else {
		log.Fatal("Gagal mendapatkan info file", err)
	}

	cmd := exec.Command("sh", "-c", `echo "Hello,from dimas" > dimas.txt`)
	if _, err := cmd.Output(); err != nil {
		log.Fatal("Error : file cant write", err)
	}
	log.Println("File dimas txt berhasil dibuat!!!")

}

func readFileWithOS(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}

	text := make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				break
			}

		}
		if n == 0 {
			break
		}
	}
	fmt.Println("===>Berhasil dibaca!!")
	fmt.Println(string(text))
}

func readFileWithExec(path string) {
	file, err := exec.Command("cat", path).Output()
	if err != nil {
		log.Fatal("Error : file can't be open")
	}
	text := string(file)
	fmt.Println(text)
}

func deleteFileWithExec(path string) {
	_, err := exec.Command("rm", path).Output()
	if err != nil {
		log.Fatal("Error : file can't deleted", err.Error())
	}
	fmt.Println("Info : Success deleted file!!!")
}

func deleteFileWithOS(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal("Erorr : file can't be deleted", err.Error())
	}
	fmt.Println("Info : Success deleted file")
}
