package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("haya", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "spanish" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}
