package main

import (
	"fmt"
	"sync"
)

func task(value int, wg *sync.WaitGroup) {
	defer wg.Done() // executa ao finalizar
	fmt.Printf("Taks %d done!\n", value)
}

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, wg)
	}

	wg.Wait()
	fmt.Println("All done!")
}
