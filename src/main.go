package main

/*
#cgo CFLAGS: -x objective-c -fobjc-arc -include MG-Prefix.pch
#cgo LDFLAGS: -framework Cocoa -framework JavaScriptCore -framework WebKit -framework AppKit

extern int NSApplicationMain(int, const char *[]);
*/
import "C"

import (
	"reflect"
	"unsafe"
)

func main() {
	C.NSApplicationMain(0, nil)
}

//export handleURL
func handleURL(url *C.char, rawBuffer unsafe.Pointer, bufferLength *C.int) {
	bufferLen := int(*bufferLength)
	hdr := reflect.SliceHeader{
		Data: uintptr(rawBuffer),
		Len:  bufferLen,
		Cap:  bufferLen,
	}
	slice := *(*[]byte)(unsafe.Pointer(&hdr))

	output := "<html><head><style>html{background-color:white;color:green;}</style></head><body>Hello, " + C.GoString(url) + "</body></html>"
	outputData := []byte(output)
	if len(outputData) > bufferLen {
		outputData = outputData[:bufferLen]
	}
	copy(slice, outputData)
	*bufferLength = C.int(len(outputData))
}
