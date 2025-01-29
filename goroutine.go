package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintMessage(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(till, message)
	}
}

func doSomething() {
	fmt.Println("Do Something heavy!!")
	time.Sleep(2 * time.Second)
	fmt.Println("Task Done!!")
}
func Goroutine() {
	//fungsi yang menentukan jumlah core
	runtime.GOMAXPROCS(2)

	// go PrintMessage(5, "ini dijalankan go routine")
	// PrintMessage(5, "ini tidak dijalankan go routine")

	go doSomething()
	go doSomething()
	fmt.Println("damn")
}
