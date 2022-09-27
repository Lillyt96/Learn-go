package main

import "fmt"

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (Prefix string) {
	switch language {
	case "French":
		Prefix = "Bonjour "
	case "Spanish":
		Prefix = "Hola "
	default:
		Prefix = "Hello "
	}

	return
}

func main() {
	fmt.Println(Hello("Lilly", ""))
}
