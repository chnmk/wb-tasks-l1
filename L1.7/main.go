package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	m := make(map[string]int)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			m["var"] = rand.IntN(100)
			fmt.Println(m["var"])
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("result: %d\n", m["var"])
}
