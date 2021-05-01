package main

import "fmt"

func HelloWorld() string {
	return "hello world"
}

func UrMom() string {
	return "ur mom"
}

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("bashbunni"))
}
