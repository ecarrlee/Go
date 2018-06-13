package main

import (
"fmt"
)

func main() {
	x:= 7
	defer fmt.Println(x)
	fmt.Println("This will be printed first.")
	/*since the defer evaluates its argument 
	immediately but waits to execute, this
	increment doesn't change the value of x
	that is printed*/
	x++
}
