package main

import (
	"fmt"

	"github.com/nombregithub/greetings"
)

func main() {
	message := greetings.Hello("manuel")
	fmt.Println(message)
}
