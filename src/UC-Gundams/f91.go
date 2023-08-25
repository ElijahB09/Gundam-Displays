package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func on(pin rpio.Pin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LEDs on\n")
		io.WriteString(w, "Gundam f91 LEDs on\n")
		pin.High()
	}
}

func off(pin rpio.Pin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LEDs off\n")                   // Goes to console
		io.WriteString(w, "Gundam f91 LEDs off\n") // Goes to webpage
		pin.Low()
	}
}

func main() {
	err_rpio := rpio.Open()
	if err_rpio != nil {
		fmt.Println("Could not open: " + err_rpio.Error())
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()

	http.HandleFunc("/f91/on", on(pin))
	http.HandleFunc("/f91/off", off(pin))

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
