package main

import (
"fmt"
)

func simplePrint(i int, c chan int) {
	c <- i
	return
}

/* If you run the program repeatedly they won't always
print in the 0...9 order because one of the go "threads"
may send to c before one of the threads that was
created before it*/
func main() {
	c:= make(chan int)
	for i := 0; i < 10; i++ {
		go simplePrint(i,c)
	}

	for j := 0; j < 10; j++ {
	x:= <-c
	fmt.Println(x)
	}
}
