package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			boiler := GetInstance()
			fmt.Printf("Goroutine %d: %p\n", id, boiler)
		}(i)
	}

	wg.Wait()

	boiler := GetInstance()
	boiler.Fill()
	boiler.Boil()
	boiler.Drain()
}
