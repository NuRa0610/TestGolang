package main

import "fmt"

func main() {
	for i := 3; i <= 100; i++ {
		fmt.Println("Masukkan Angka kelipatan 3 atau 5")
		fmt.Scan(&i)

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Hello World")
		} else if i%5 == 0 {
			fmt.Println("World")
		} else if i%3 == 0 {
			fmt.Println("Hello")
		} else {
			fmt.Println("Angka tidak berhubungan dengan 3 atau 5, coba lagi")
		}
	}
}
