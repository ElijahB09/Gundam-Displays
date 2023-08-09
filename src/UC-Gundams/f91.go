package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	http.HandleFunc("/f91/on", led_on)
	http.HandleFunc("/f91/off", led_off)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func led_on(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LEDs on\n")
	err := rpio.Open()
	if err != nil {
		fmt.Println("Could not open: " + err.Error())
		os.Exit(1)
	}
	defer rpio.Close()
	pin := get_pin()
	pin.Output()
	pin.Low()
}

func led_off(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LEDs off\n")
	pin := get_pin()
	pin.High()
	rpio.Close()
}

func get_pin() rpio.Pin {
	pin := rpio.Pin(18)
	return pin
}
