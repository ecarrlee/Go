package main

import (
	"fmt"
)

func patchedfunc(x int)(y int) {
	defer func() {
		y++
	}()
	return x
}

func main() {
	/*prints 8 because the defer 
	increments the return value
	after the return happens!*/
	fmt.Println(patchedfunc(7))
}
