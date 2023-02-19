package main

import (
	"fmt"

	"github.com/KnoblauchPilze/template-go-cpp/pkg/example"
)

func main() {
	fmt.Print(example.FormatHelloWorld("cli"))
	example.PrintHelloWorldCpp()
}
