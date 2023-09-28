package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == ""{
		name = "World"
	}

	return englishHelloPrefix + name
}

func Bye() string{
	return "Bye, world"
}

func main() {
	fmt.Println(Bye())
}
