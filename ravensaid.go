// #cgo CFLAGS: -I/path/to/headers
// #cgo LDFLAGS: -L/path/to/library -lravensaid
// #include <ravensaid.h>
// #include <stdlib.h>
package main

import (
	"C"
	"unsafe"
)

type RavensaidState C.RavensaidState

func RavensaidInit(path string) *RavensaidState {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return (*RavensaidState)(C.ravensaid_init(cPath))
}

func RavensaidFree(state *RavensaidState) {
	C.ravensaid_free((*C.RavensaidState)(state))
}

func Ravensaid(state *RavensaidState, message string) int {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	return int(C.ravensaid((*C.RavensaidState)(state), cMessage))
}
