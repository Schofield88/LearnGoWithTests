package main

import (
	"fmt"
)

const french = "French"
const czech = "Czech"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Salut, "
const czechHelloPrefix = "Cao, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case czech:
		prefix = czechHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main () {
	fmt.Println(Hello("Luke", ""))
}
