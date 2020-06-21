package main

import (
	"fmt"
	"os"
	"plugin"
)

// step:
// 1. go build -buildmode=plugin -o eng/eng.so eng/greeter.go
// 2. go build -buildmode=plugin -o chi/chi.so chi/greeter.go
// 3. go run greeter.go english
// 4. go run greeter.go chinese

type Greeter interface {
	Greet()
}

func main() {
	// determine module to load
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	var mod string
	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chinese":
		mod = "./chi/chi.so"
	default:
		fmt.Println("don't speak")
		os.Exit(1)
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexcepted type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	greeter.Greet()
}
