package main

import (
"fmt"
)

/*Since the channel has a buffer of
3, it can hold up to 3 values
before blocking*/
func main() {
	c := make(chan int, 3)
	c <- 5
	c <- 4
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
