package main

import "fmt"

func main() {
	//consumeChannel function to consume channel 
	
     ch:=ConsumeChannel() 
	 // for for Producing the channel
		for i:=0;i<10;i++{
		ch<-i+1
		fmt.Println("Produced",i+1)
	}
}
func ConsumeChannel() chan<- int {

	ch1:=make(chan int)

	go func(){
		for i:= range ch1{
			fmt.Println("consumed",i)
		
		}
	}()

	return ch1
}
