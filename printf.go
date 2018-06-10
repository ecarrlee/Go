
package main
//import statements are weird in go, although you can also do them
//the "normal way"
import (
	"fmt"
	"math"
)

//Printf style string formatting in go!
//unlike Println this does not add a 
//newline character, so don't forget to
//add one! 
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
