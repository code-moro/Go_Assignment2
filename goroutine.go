package main

import "fmt"

func routone(arr []int,ch chan int){
	slice:=arr[:50]
	sum:=0
	for i:=0;i<len(slice);i++{
     sum=sum+slice[i]
	}
	ch<-sum
} 

func routtwo(arr []int,ch chan int){
	slice:=arr[50:]
	sum:=0
	for i:=0;i<len(slice);i++{
     sum=sum+slice[i]
	}
	ch<-sum
} 
func main(){
    // unbuffered channel created 
	ch:=make(chan int)
	// arr created of size 100
	arr := make([]int, 100)
	for i:=0;i<100;i++{
		arr[i]=i
	}
	// call for 1st go routine to calculate sum of 1st 50 numbers
    go routone(arr,ch)
	fmt.Println("sum of 0 to 50 range numbers",<-ch)
	// call for 1st go routine to calculate sum of 1st 50 numbers
	go routtwo(arr,ch)
	fmt.Println("sum of 51 to 100 range numbers",<-ch)

}