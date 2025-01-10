package main

import (
	"fmt"
	"os"
)

var VERSION = os.Getenv("HIVEBOX_VERSION")

func main() {
	printVersion()
}

func printVersion() {
	fmt.Println(VERSION)
}
