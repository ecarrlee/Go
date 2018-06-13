package main

import (
	"fmt"
)

func main() {
	/*Defers go on a stack. This 
	follows a last in first out
	order, so the numbers will
	print in "reverse" order.*/
	for i:= 0; i < 9; i+=2 {
		defer fmt.Println(i)
	}
}
