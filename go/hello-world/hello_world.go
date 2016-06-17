package hello

import "fmt"

const testVersion = 2

// HelloWorld says 'Hello, World!'
func HelloWorld(name string) string {
	if len(name) == 0 {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
