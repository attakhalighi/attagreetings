package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func Goodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}
