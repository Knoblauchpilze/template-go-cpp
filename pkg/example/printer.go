package example

import "fmt"

// #include <stdlib.h>
// #cgo CFLAGS: -I../../cpp/include
// #include <toy-lib/lib_c_wrapper.h>
// #cgo LDFLAGS: -L../../cpp/lib -ltoy-lib
import "C"

func FormatHelloWorld(app string) string {
	return fmt.Sprintf("Hello world from %s!\n", app)
}

func PrintHelloWorldCpp() {
	C.lib_c_wrapper_say_hello()
}
