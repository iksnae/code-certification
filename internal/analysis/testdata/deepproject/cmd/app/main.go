package main

import (
	"fmt"

	"example.com/deepproject/pkg/greet"
)

func main() {
	fmt.Println(greet.Hello("world"))
	fmt.Println(greet.Goodbye("world"))
	fmt.Println(greet.Format("direct"))
	run()
}

func run() {
	fmt.Println(greet.Hello("again"))
}
