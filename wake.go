package mac_wake_display

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation -framework IOKit
// #include <Foundation/Foundation.h>
// #include <IOKit/pwr_mgt/IOPMLib.h>
// #include <IOKit/IOKitLib.h>
// #include <stdio.h>
// #include <lowlevel.h>
import (
	"C"
)

func WakeDisplay() {
	C.wakeDisplay()
}
