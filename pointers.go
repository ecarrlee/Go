package main

import (
"fmt"
)

func main() {
i := 7
p := &i
fmt.Println(*p)

*p++
fmt.Println(*p)
}
