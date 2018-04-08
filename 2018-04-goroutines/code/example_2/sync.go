package main

import (
	"fmt"
	"math/rand"
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

func randVectorSync(size int) []float32 {
	result := make([]float32, size)
	for i := range result {
		result[i] = float32(rand.Intn(99) + 1)
	}
	return result
}

func main() {
	startTime := time.Now()
	result := make([]float32, 10000)
	values := randVectorSync(10000)

	for i, value := range values {
		result[i] = sqrtSync(value)
	}
	fmt.Println("Time: ", time.Now().Sub(startTime))
}
