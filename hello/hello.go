package main

import "fmt"

const englishHelloprefix = "Hello,"
const spanish = "Spanish"
const spanishHelloprefix = "Hola,"
const frenchHelloprefix = "Bonjour,"
const french = "French"

func Hello(s, language string) string {
	if s == "" {
		s = "World"
	}
	return greetingPrefix(language) + s
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloprefix
	case french:
		prefix = frenchHelloprefix
	default:
		prefix = englishHelloprefix
	}
	return
}
func main() {
	fmt.Println(Hello("world", ""))
}
