package main

import (
	"fmt"
	"time"
)

func task(value int) {
	fmt.Printf("Taks %d done!\n", value)
}

func main() {
	for i := 0; i < 10; i++ {
		go task(i)
	}
	time.Sleep(time.Second)
	fmt.Println("All done!")
}
