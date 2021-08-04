package main

import (
	"fmt"
	"runtime"
)

// go build <filename> = to build go file to executables
// go install = to install the go file to path

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}