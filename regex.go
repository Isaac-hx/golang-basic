package main

import (
	"errors"
	"fmt"
	"regexp"
)

var regexPatternEmail = regexp.MustCompile(`^([a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`)

func RegexP() {
	text := "dimas Ananda riyadi"
	text2 := "12w3"
	//compile regex expression dan mereturn object regexp.Regexp
	regex, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	//method yang digunakan untuk mencari semua nilai expresi yang cocok pada argument
	res1 := regex.FindAllString(text, 2)
	fmt.Println("Output res1 : ", res1)

	res2 := regex.FindAllString(text, -1)
	fmt.Println("Output res2 :", res2)

	//method yang digunakan untuk mendeteksi apakah nilai string cocok pada pola expresi regex
	res3 := regex.MatchString(text2)
	fmt.Println("Output res3 : ", res3)

	//method yang digunakan untuk mencari apakah nilai string memenuhi kriteri pada pola expresi regex
	res4 := regex.FindString(text)
	fmt.Println("Output res4 : ", res4)

	//method yang digunakan untuk mencari index string dari hasil operasi regexp
	idx := regex.FindStringIndex(text)
	fmt.Println("Output idx : ", idx, text[idx[0]:idx[1]])

	//method yang digunakan untuk mereplace nilai string berdasarkan pola expresi regex
	res5 := regex.ReplaceAllString(text, "dimas ganteng")
	fmt.Println("Output res5 : ", res5)

	//method yang digunakan untuk mereplace nilai string dan melakukan pengkodisian tertentu
	res6 := regex.ReplaceAllStringFunc(text, func(s string) string {
		if s == "dimas" {
			return "pecundang"
		} else if s == "riyadi" {
			return "adakah aku dipikiranmu?"
		}
		return s
	})
	fmt.Println("Output res6 : ", res6)

	regex2, err := regexp.Compile(`[a-i]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	//method yang digunakan untuk melakukan split terhadap pola expresi yang dicompile
	res7 := regex2.Split(text, -1)
	fmt.Println("Output res7 : ", res7)

	var inputEmail, inputPassword string
	fmt.Printf("Masukan email yang valid! : ")
	fmt.Scanln(&inputEmail)
	fmt.Printf("Masukan password yang valid! : ")
	fmt.Scanln(&inputPassword)
	resultEmail := validateEmail(inputEmail)
	if resultEmail != nil {
		fmt.Println(resultEmail.Error())
	}

}

func validateEmail(s string) error {
	result := regexPatternEmail.MatchString(s)
	if !result {
		return errors.New("Error : input bukan email yang valid!")
	}
	return nil

}
