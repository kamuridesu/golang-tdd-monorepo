package main

import "fmt"

const francaisHello string = "Bonjour"
const francais string = "french"
const portuguesHello string = "Bom dia"
const portugues string = "portuguese"
const indonesianHello string = "Selamat pagi"
const indonesian string = "indonesian"

func getPrefix(language string) string {
	prefix := francaisHello
	switch language {
	case francais:
		prefix = francaisHello
	case portugues:
		prefix = portuguesHello
	case indonesian:
		prefix = indonesianHello
	}
	return prefix
}

func Hello(name string, language string) string {
	if name != "" {
		name = " " + name
	}

	return getPrefix(language) + name
}

func main() {
	fmt.Println(Hello("kamuri", "french"))
}
