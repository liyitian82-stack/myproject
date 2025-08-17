package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for {
			num := rand.Intn(50)
			ch <- num
			time.Sleep(200 * time.Millisecond)
		}
	}()
	for i := 1; i <= 2; i++ {
		go func(id int) {
			for num := range ch {
				fmt.Printf("Reader %d num: %d\n", id, num)
			}
		}(i)
	}

	select {}
}
