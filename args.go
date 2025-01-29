package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func calculate(operand, n1, n2 string) int {
	number1Conv, _ := strconv.Atoi(n1)
	number2Conv, _ := strconv.Atoi(n2)
	if operand == "+" {
		return number1Conv + number2Conv
	} else if operand == "-" {
		return number1Conv - number2Conv
	} else if operand == "*" {
		return number1Conv * number2Conv
	}
	return number1Conv / number2Conv
}

func Args() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Error : Kesalahan format penulisan !!!")
		}
	}()
	catchArgs := os.Args
	// fmt.Printf("-> %#v\n", argsAw)
	// fmt.Println(argsAw)

	// args := argsAw[1:]
	// fmt.Printf("-> %#v\n", args)
	argsList := catchArgs[1:]
	fmt.Println(calculate(argsList[1], argsList[0], argsList[2]))

}

func Flag() {
	var name, country string
	var age int
	flag.StringVar(&name, "name", "anonymous", "Type your name?")
	flag.StringVar(&country, "country", "default", "Type your country?")
	flag.IntVar(&age, "age", 10, "Type your age?")
	flag.Parse()

	fmt.Println(createFlagsSetence(age, name, country))
}

func createFlagsSetence(age int, name, country string) string {
	return fmt.Sprintf("Halo,%s! Anda berusia %d tahun dan berasal dari %s", name, age, country)
}
