package main

import "fmt"

const (
	 spanish = "Spanish"
	english = "English"
	french = "French"

	latinHelloPrefix = "Salve, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


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

func greetingPrefix(language string) (prefix string) {
		switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case english:
			prefix = "Hello, "
		case french:
			prefix = "Bonjour, "
	}
	return
}

func main() {
	fmt.Println(Greet("world", "english"))
}
