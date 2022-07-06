package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	ch1 := make(chan int)
	out := make(chan int)
   // Produced 1st Channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
			fmt.Println("produced channel1", i+1)
		}
		close(ch)
	}()
    // Produced 2nd Channel
	go func() {
		for i := 10; i < 21; i++ {
			ch1 <- i + 1
			fmt.Println("produced channel2 ", i+1)
		}
		close(ch1)
	}()

	go Produce(ch, ch1, out)

	for v := range out {
		fmt.Println("consumed channel3", v)
		time.Sleep(2 * time.Second)
	}

}
//Produced 3rd channel from 1st and 2nd channel

func Produce(ch, ch1, out chan int) {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func(){for v := range ch {
		out <- v
	}
	wg.Done()
	}()
	go func(){for v := range ch1 {
		out <- v
	}
	wg.Done()
	}()
	wg.Wait()
	close(out)
}
