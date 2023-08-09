package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func led_on(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LEDs on\n")
	io.WriteString(w, "Gundam f91 LEDs on\n")
	pin := get_pin()
	pin.Output()
	pin.Low()
}

func led_off(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LEDs off\n")                   // Goes to console
	io.WriteString(w, "Gundam f91 LEDs off\n") // Goes to webpage
	pin := get_pin()
	pin.High()
	rpio.Close()
}

func get_pin() rpio.Pin {
	pin := rpio.Pin(18)
	return pin
}

func main() {
	err_rpio := rpio.Open()
	if err_rpio != nil {
		fmt.Println("Could not open: " + err_rpio.Error())
		os.Exit(1)
	}
	defer rpio.Close()

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
