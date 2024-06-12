package main

import "fmt"

const HelloPrefixEnglish = "Hello, "

func Hello(name string) string {
	return HelloPrefixEnglish + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
