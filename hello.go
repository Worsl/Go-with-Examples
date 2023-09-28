package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func Bye() string{
	return "Bye, world"
}

func main() {
	fmt.Println(Bye())
}
