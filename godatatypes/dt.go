package main

import (
	"fmt"
	"reflect" //used to get the type of a variable
)

func main() {

	fmt.Println("Type of 20:", reflect.TypeOf(20))
	fmt.Println("Type of 3.14:", reflect.TypeOf(3.14))
	fmt.Println("Type of true:", reflect.TypeOf(true))
	fmt.Println("Type of \"Hello, World\":", reflect.TypeOf("Hello, World"))             //string is a collection of characters
	fmt.Println("Type of 'R':", reflect.TypeOf('R'))                                     //rune is a single character
	fmt.Println("Type of nil:", reflect.TypeOf(nil))                                     //nil is a special type that represents no value
	fmt.Println("Type of struct:", reflect.TypeOf(struct{ name string }{name: "John"}))  //struct is a collection of fields
	fmt.Println("Type of slice:", reflect.TypeOf([]int{1, 2, 3}))                        //slice is a collection of elements
	fmt.Println("Type of map:", reflect.TypeOf(map[string]int{"apple": 1, "banana": 2})) //map is a collection of key-value pairs
	fmt.Println("hello harsha")

	//math ops

	var originalcount int = 20
	fmt.Println("i had", originalcount, "apples")

	var eaten int = 5
	fmt.Println("some apples are eaten by harsha", eaten, "fuck")

	fmt.Println("remaining are", originalcount-eaten, "apples thats it")

	//tax example
	var price int = 100
	fmt.Println("total price is: ", price)

	var taxrate float64 = 0.08
	var tax float64 = float64(price) * taxrate //float64(price) converts price to float datatype to calculate tax
	fmt.Println("tax rate is", tax, "bucks")

	var total float64 = float64(price) + tax
	fmt.Println("total cost is", total, "bucks")

	var availfunds int = 120
	fmt.Println(availfunds, "bucks available to spned")
	fmt.Println("within budget?", total <= float64(availfunds))

}
