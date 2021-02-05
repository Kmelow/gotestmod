package gotestmod

import "fmt"

func GreetGoTestMod(name string) string {
	return fmt.Sprintf("Greetings %s, you have been hacked halfling\n", name)
}
