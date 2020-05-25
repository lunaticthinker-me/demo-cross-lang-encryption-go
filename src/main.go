package main

import (
	"fmt"

	"github.com/templ-project/go/src/pkg/greet"
	"github.com/templ-project/go/src/pkg/ver"
)

func main() {
	fmt.Println(greet.Hello("World"))
	fmt.Println()
	ver.ShowVersion()
}
