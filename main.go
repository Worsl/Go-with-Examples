package main

import "fmt"

func main(){
	array := [5] int {1,2,3,4,5}

	for i, number := range array{
		fmt.Println(i,number)
	}
}


// new timing is 5.26 pm
