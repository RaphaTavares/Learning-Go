package main

import "fmt"

const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix(language) + name
}

func helloPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = helloPrefixFrench
	case "spanish":
		prefix = helloPrefixSpanish
	default:
		prefix = helloPrefixEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", "english"))
}
