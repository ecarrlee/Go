package main

import "fmt"
/* The return values are treated as
being defined at start of function,
so if I don't define one of them then
it defaults to 0 because it is an int.
*/
func double (input int) (x, y int) {
x = input *2
return
}

func main() {
fmt.Println(double(4))
}
