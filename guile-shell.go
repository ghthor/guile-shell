package main

// #cgo pkg-config: guile-2.0
// #include "guile-shell.h"
import "C"

import (
	"fmt"
	"time"
)

//export LogNow
func LogNow() {
	fmt.Println(time.Now())
}

func main() {
	fmt.Println("starting guile shell")
	C.LaunchGuile()
}
