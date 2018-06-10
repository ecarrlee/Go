
package main

import "fmt"

/* The two arguments are both ints
so you can specify the type once.
Also, multiline comments work
in go!*/ 
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
