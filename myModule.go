package main

import (
	"fmt"
	"runtime"
)

func Version() {
	fmt.Println("Using go version ", runtime.Version())
}
