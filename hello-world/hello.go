package main

import (
	"fmt"
)

const french = "French"
const czech = "Czech"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Salut, "
const czechHelloPrefix = "Cao, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case czech:
		prefix = czechHelloPrefix
	}

	return prefix + name
}

func main () {
	fmt.Println(Hello("Luke", ""))
}
