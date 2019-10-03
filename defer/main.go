package main

import "fmt"

func main() {
	defer fmt.Println("1st defer")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("I=%v\n", i)
	}
	defer fmt.Println("end")
}
