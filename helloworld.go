package main

import "fmt"

// Hello says hello
func Hello() string {
	return "Hello World"
}

// HelloName says hello to somebody
func HelloName(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloName("packy"))
}
