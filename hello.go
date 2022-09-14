package main

import "fmt"

const englishPrefixHello string = "Hello "

func Hello(name string) string {
	if len(name) == 0 {
		return englishPrefixHello + "World"
	}
	return englishPrefixHello + name
}

func main() {
	fmt.Println(Hello("Lilly"))
}
