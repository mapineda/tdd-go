package main

import "fmt"

const latinHelloPrefix = "Salve, "
func Hello(name string) string  {
	if name == "" {
		name = "World"
	}
	return latinHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
