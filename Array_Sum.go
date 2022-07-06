package main

import (
	"fmt"
	"sync"
)



func routineone(arr []int,ch chan int,wg *sync.WaitGroup){
	slice:=arr[:51]
	sum:=0
	for i:=0;i<len(slice);i++{
     sum=sum+slice[i]
	}
	ch<-sum
	wg.Done()
} 

func routinetwo(arr []int,ch chan int,wg *sync.WaitGroup){
	slice:=arr[50:]
	sum:=0
	for i:=0;i<len(slice);i++{
     sum=sum+slice[i]
	}
	ch<-sum
	wg.Done()
} 

func main(){
	wg := new(sync.WaitGroup)

	wg.Add(2)
    // unbuffered channel created 

	ch:=make(chan int)
	// arr created of size 100

	arr := make([]int, 100)
	for i:=0;i<100;i++{
		arr[i]=i
	}

	// call for 1st go routine to calculate sum of 1st 50 numbers
    go routineone(arr,ch,wg)

	fmt.Println("sum of 0 to 50 range numbers",<-ch)
	// call for 1st go routine to calculate sum of 1st 50 numbers

	go routinetwo(arr,ch,wg)

	fmt.Println("sum of 51 to 100 range numbers",<-ch)
    
	wg.Wait()

}