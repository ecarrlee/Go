package main

import (
"fmt"
)

func sender(c chan int, duration int) {
	for i:=0; i < duration; i++{
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go sender(c,10)
	/*This loop runs until 
	the channel is closed,
	and each time reads the
	next value from chan into
	x*/
	for x:= range c{
		fmt.Println(x)
	}

}
