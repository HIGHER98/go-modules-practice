package myModule

import (
	"fmt"
	"runtime"
)

func Version() {
	fmt.Println("Using go version ", runtime.Version())
}
