package main

import "fmt"



func routineone(arr []int,ch chan int,a,b int){
	sum:=0
	for i:=a;i<b;i++{
     sum=sum+arr[i]
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


    go routineone(arr,ch,0,50)

	fmt.Println("sum of 0 to 50 range numbers",<-ch)
	

	go routineone(arr,ch,51,100)

	fmt.Println("sum of 51 to 100 range numbers",<-ch)
	close(ch)
    

}