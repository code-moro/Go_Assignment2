package main

import "fmt"

func main() {

	ch := make(chan int)
	ch1 := make(chan int)
	out := make(chan int)
   // Produced 1st Channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
			fmt.Println("produced ch one", i+1)
		}
		close(ch)
	}()
    // Produced 2nd Channel
	go func() {
		for i := 11; i < 21; i++ {
			ch1 <- i + 1
			fmt.Println("produced ch two", i+1)
		}
		close(ch1)
	}()

	go Produce(ch, ch1, out)

	for v := range out {
		fmt.Println("consumed ch three", v)
	}

}
//Produced 3rd channel from 1st and 2nd channel

func Produce(ch, ch1, out chan int) {
	for v := range ch {
		out <- v
	}
	for v := range ch1 {
		out <- v
	}
	close(out)
}
