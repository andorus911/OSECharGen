package main

import (
	"fmt"
	"sync"
)

func main() {
	portal := make(chan Character, 3)
	var wg sync.WaitGroup

	go func() {
		defer close(portal)
		wg.Wait()
	}()

	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			for i := 0; i < 100; i++ {
				portal <- genChar_0lvl()
			}
		}()
	}

	i := 1
	for range portal {
		fmt.Printf("%v %+v\n", i, <-portal)
		i++
	}
}