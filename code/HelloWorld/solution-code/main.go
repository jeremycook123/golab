//CODE1:define the main package
package main

//CODE2:import the string formatting package
import (
	"fmt"
)

//CODE3:define a basic helloworld function that takes an input message parameter
func helloworld(message string) {
	fmt.Println("helloworld!!")
	fmt.Println(message)
	return
}

//CODE4:define an init function, this is called automatically at startup time
func init() {
	fmt.Println("init called automatically...")

	//initialisation code goes here
}

//CODE5:define the main function, a main function must be present in the main package
func main() {
	fmt.Println("main started...")

	message := "cloudacademy + go = awesomeness!"

	helloworld(message)

	fmt.Println("main finished...")
}
