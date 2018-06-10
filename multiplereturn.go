package main

import "fmt"

func clone(x string) (string,string) {
return x, x
}

func main() {
original, copy := clone("Dolly")
fmt.Println(original,copy)
fmt.Println(clone("Dolly"))
}
