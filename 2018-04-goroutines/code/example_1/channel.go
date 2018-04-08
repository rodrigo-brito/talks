package main

import (
	"fmt"
)

func task(value int, done chan<- bool) {
	fmt.Printf("Task %d done!\n", value)
	done <- true
}

func main() {
	done := make(chan bool)
	go task(1, done)
	go task(2, done)
	go task(3, done)
	<-done
	<-done
	<-done
	fmt.Println("All done!")
}
