package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}
func retrieveData(ch <-chan int) {

	for {
		// select {
		// case data := <-ch:
		// 	fmt.Println("receive :  ", data)
		// case <-time.After(time.Second * 10):
		// 	fmt.Println("timeout data will not receive again")
		// 	break loop
		// }
		fmt.Println(<-ch)
	}
}

func generateSendData(ch chan<- string) {
	randomTime := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; true; i++ {

		ch <- fmt.Sprintf("TRX-1204-%v", i)
		time.Sleep(time.Duration(randomTime.Int()%10+1) * time.Second)

	}
}

func TimesOut() {
	// messages := make(chan int)
	// go sendData(messages)
	// retrieveData(messages)
	data := make(chan string)
	go generateSendData(data)
	func(ch <-chan string) {
	loop:
		for {
			select {
			case receiveData := <-ch:
				fmt.Println("Data received : ", receiveData)
			case <-time.After(time.Second * 4):
				fmt.Println("Timeout!! Data not received")
				break loop

			}
		}
	}(data)
}
