package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const french = "French"

const latinHelloPrefix = "Salve, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Greet(name string, language string) string  {
	if name == "" {
		name = "World"
	}

	prefix := latinHelloPrefix

	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case english:
			prefix = "Hello, "
		case french:
			prefix = "Bonjour, "
	}

	return prefix + name
}

func main() {
	fmt.Println(Greet("world", "english"))
}
