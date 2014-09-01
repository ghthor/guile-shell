package main

// #cgo pkg-config: guile-2.0
// #include "guile-shell.h"
// #include "charArray.h"
import "C"

import (
	"fmt"
	"os"
	"time"
)

//export LogNow
func LogNow() {
	fmt.Println(time.Now())
}

func main() {
	fmt.Println("starting guile shell")

	args := os.Args[1:]

	cargs := C.MakeCharArray(C.int(len(args)))
	defer C.FreeCharArray(cargs, C.int(len(args)))

	for i, s := range args {
		C.SetArrayString(cargs, C.CString(s), C.int(i))
	}

	C.LaunchGuile(C.int(len(args)), cargs)
}
