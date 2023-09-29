package main


const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Fake Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix

	case "Spanish":
		prefix = spanishHelloPrefix

	}
	return prefix
}

func new_greetingPrefix(language string) (prefix string){
	switch language{
	case "French":
		prefix = frenchHelloPrefix
	
	case "Spanish":
		prefix = spanishHelloPrefix
	
	default:
		prefix = englishHelloPrefix
	}
	return
}
