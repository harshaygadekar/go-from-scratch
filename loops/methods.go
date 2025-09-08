package main

import (
	"fmt"
	"strings" //importing strings package to use its functions
	"time"    //importing time package to use its functions
)

func main() {
	var now time.Time = time.Now() //here time.Time represents a built in data type whici used to store formatted time values, in go that represents a specific point in time
	var year int = now.Year()      //now.Year() is a method that returns the year of the time.Time object
	//now is a variable of type time.Time and Year() is a method that returns the year of the time.Time object
	//the value returned by now.Year() is assigned to the variable year
	//now.Year() is a method that returns the year of the time.Time object.
	var month int = int(now.Month())
	fmt.Println("current date and time is:", now)
	fmt.Println(month)
	fmt.Println(year)

	brokenstring := "G$$ O$$ is"
	replacer := strings.NewReplacer("$", " ", "G", "Go") //strings.NewReplacer is a function that creates a new Replacer object that can be used to replace substrings in a string
	fmt.Println(replacer.Replace(brokenstring))          //replacer.Replace is a method that replaces the substrings in the string with the values specified in the Replacer object
}
