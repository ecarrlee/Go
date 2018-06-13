package main

import(
	"fmt"
)

func main() {
	i := 7
	j := 4
	/* Switch without a condition just
	switches upon "true". Since go
	switch's automatically break, this
	is equivalent to a long chain of 
	if-elses-if's.*/
	for {
		switch {
			case j > i:
				fmt.Println(1)
			case j == i:
				fmt.Println(2)
			case j <= i:
				fmt.Println(3)
			/*never evaluated because the previous
			case is true*/
			case j < 10:
				fmt.Println("not executed")

			}
		}
}
