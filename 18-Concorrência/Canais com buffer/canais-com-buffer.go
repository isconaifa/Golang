package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Ani"
	canal <- "Priscila"
	message := <-canal
	fmt.Println(message)
}
