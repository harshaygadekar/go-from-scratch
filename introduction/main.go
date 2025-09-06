package main

import (
	"fmt"

	"rsc.io/quote"
) //importing quote package

func main() {
	fmt.Println("hi i want to learn go", "i hate javascript")
	yes()

	//calling code in an external package,
	fmt.Println(quote.Go())

}

//Test cases:

/*printed the line using only print() but prints normal no changes in output
used multiple values separated by commas, ex: fmt.Print("hi i want to learn go","i hate javascript")
but here there is no space between values printed in output screen.

when used with println it detects the end of the values and adds space betwee values printed.
*/
