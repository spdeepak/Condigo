package main

import (
	"fmt"

	"github.com/spdeepak/condigo"
)

func main() {
	//When you want to get a string object out of the conditions using just Then & Else
	result := condigo.When(true).
		Then("Success").
		Else("Failure")
	fmt.Println(result)

	//When you want to get a string object out of the conditions using just Then, ElseIf & Else
	result = condigo.When(false).
		Then("First").
		ElseIf(true, "Second"). // You can have multiple ElseIf here
		Else("Default")
	fmt.Println(result)

	//When you can want a function in Then/ElseIf/Else
	resultFunc, _ := condigo.When(true).
		Then(func() bool {
			fmt.Println("Inside Then")
			return true
		}).
		Else(func() bool {
			fmt.Println("Inside Else")
			return false
		}).(func() bool)
	fmt.Println(resultFunc())
}
