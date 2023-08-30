package uc

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

var pin rpio.Pin
var activated bool = false

func on(pin rpio.Pin) {
	pin.High()
}

func off(pin rpio.Pin) {
	pin.Low()
}

func activate() {
	pin = rpio.Pin(18)
	pin.Output()
	activated = true
}

func ToggleF91(command []byte) {
	instruction := string(command[:])

	switch instruction {
	case "on":
		if activated {
			on(pin)
		}
		break
	case "off":
		if activated {
			off(pin)
		}
		break
	case "activate":
		err_rpio := rpio.Open()
		if err_rpio != nil {
			fmt.Println("Could not open: " + err_rpio.Error())
			os.Exit(1)
		}
		defer rpio.Close()
		activate()
		break
	case "deactivate":
		rpio.Close()
	}
}
