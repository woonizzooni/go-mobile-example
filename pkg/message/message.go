package message

import "fmt"

// Greetings ...
func Greetings(name string) string {
	return fmt.Sprintf("Hey, %s!", name)
}

// Goodbyes ...
func Goodbyes(name string) string {
	return fmt.Sprintf("Bye, %s!", name)
}
