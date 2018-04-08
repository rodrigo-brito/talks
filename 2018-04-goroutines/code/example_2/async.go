package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sqrtSync(value float32) float32 {
	result := value
	for i := 0; i <= int(value); i++ {
		result = (result + value/result) / 2
	}
	time.Sleep(time.Millisecond) // Simulando processo pesado
	return result
}

func randVectorAsync(size int, values chan<- float32) {
	for i := 0; i < size; i++ {
		values <- float32(rand.Intn(99) + 1)
	}
	close(values)
}

func sqrtAsync(values <-chan float32, results chan<- float32, wg *sync.WaitGroup) {
	for value := range values {
		results <- sqrtSync(value)
	} // finaliza quando canal `values` fechar
	wg.Done()
}

func main() {
	const asyncTasks = 8

	startTime := time.Now()
	values := make(chan float32, 10000)
	results := make(chan float32, 10000)

	// generator
	go randVectorAsync(10000, values)

	wg := new(sync.WaitGroup)
	wg.Add(asyncTasks)
	for i := 0; i < asyncTasks; i++ {
		// consumers
		go sqrtAsync(values, results, wg)
	}

	wg.Wait()
	fmt.Println("Time: ", time.Now().Sub(startTime))
}
