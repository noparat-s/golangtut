package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		runa()
	}()
	fmt.Println("func has run")
	time.Sleep(time.Second * 10)
}

func runa() {
	for i := 1; i < 11; i++ {
		fmt.Printf("I = %v\n", i)
		time.Sleep(time.Second)
	}
}
