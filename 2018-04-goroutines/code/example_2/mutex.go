package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)

	var counter int32 = 1
	atomic.AddInt32(&counter, 2) // Operação atômica

	wg.Add(1)
	go func() {
		mutex.Lock() // bloqueia acesso a contador
		counter = counter + 2
		mutex.Unlock() // libera acesso a contador
		wg.Done()
	}()
	wg.Wait()

	println("Counter = ", counter)
}
