package main

import (
	"fmt"
	"regexp"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^([a-z0-9._%+\-]{2,20})+@([a-z0-9.\-]{2,20})+\.(co.id||id)$`)
	return Re.MatchString(email)
}

func main() {

	var e string

	fmt.Scan(&e)

	if !validateEmail(e) {
		fmt.Println("Format Email yang dimasukan Salah")
	} else {
		fmt.Println("Format Email yang dimasukan Benar")
	}

}
