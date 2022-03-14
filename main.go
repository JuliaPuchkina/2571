package main

import (
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Print("Введите произвольные данные:")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %v\n", n)
}
