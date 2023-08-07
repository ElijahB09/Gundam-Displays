package ucgundams

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("Could not open: " + err.Error())
		os.Exit(1)
	}

	defer rpio.Close()

	pin := rpio.Pin(18)

	pin.Output()

	for i := 0; i < 5; i++ {
		pin.Low()
	}
	pin.High()
}
