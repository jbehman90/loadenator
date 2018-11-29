package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	//out := make(chan string)

	// for i := 0; i < 100000; i++ {
	// 	fmt.Printf("hello dev %v \n", i)
	// }

	var wg sync.WaitGroup
	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("hello async %v\n", i)
		}(i)

	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

}
