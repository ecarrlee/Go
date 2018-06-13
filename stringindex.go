package main

import (
"fmt"
)

func main() {
	str := "example 123"
	for i := 0; i < len(str); i++ {
		//This will print numerical byte values
		fmt.Println(str[i])
		//This will print individual characters
		fmt.Println(string(str[i]))
	}
}
