package main

/*
#cgo CFLAGS: -x objective-c -fobjc-arc -include MG-Prefix.pch
#cgo LDFLAGS: -framework Cocoa -framework JavaScriptCore -framework WebKit -framework AppKit

extern int NSApplicationMain(int, const char *[]);
*/
import "C"

func main() {
	C.NSApplicationMain(0, nil)
}
