package creategomod

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. welcome", name)
	return message
}

/*here we are creating a module named create_go_mod
also were are defining a function HelloFucker which takes a string input and returns a string

we are using Sprintf function of fmt package to format the string and return it
Sprintf formats according to a format specifier and returns the resulting string
%v is a placeholder for the value of the variable name

to use this module in other files we need to import it using import "create_go_mod"
and then we can call the function using create_go_mod.hello_fucker("name")
we can also create a go.mod file using command go mod init create_go_mod
this will create a go.mod file in the current directory with the module name create_go_mod */
