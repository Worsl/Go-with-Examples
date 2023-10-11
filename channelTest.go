package main

import (
	"fmt"
)

// the goal is to see if channel go routines are truly random
func main(){

	dogChannel := make(chan string)
	go func (){dogChannel <- "woof!"}()

	catChannel := make(chan string)
	go func(){catChannel <- "Meow!"}()

	select{

	case msgFromDog := <-dogChannel:
		fmt.Println(msgFromDog)

	case msgFromCat := <- catChannel:
		fmt.Println(msgFromCat)

	}


}