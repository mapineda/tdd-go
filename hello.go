package main

import "fmt"

func Hello(name string) string  {
	return "Salve, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
