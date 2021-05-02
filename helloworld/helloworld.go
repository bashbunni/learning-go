package main

import "fmt"

// prefixes by language
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

// languages
const french = "french"
const spanish = "spanish"

func UrMom() string {
	return "ur mom"
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return LanguagePrefix(lang) + name
}

func LanguagePrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("bashbunni", french))
}
