// Simple program that displays the state of the specified joystick
//
//	go run joysticktest.go 2
//
// displays state of joystick id 2
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/0xcafed00d/joystick"
)

func readJoystick(js joystick.Joystick) {
	jinfo, err := js.Read()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Println("Buttons:")
	for button := 0; button < js.ButtonCount(); button++ {
		if jinfo.Buttons&(1<<uint32(button)) != 0 {
			fmt.Print("X")
		} else {
			fmt.Print(".")
		}
	}

	fmt.Println()
	for axis := 0; axis < js.AxisCount(); axis++ {
		fmt.Printf("Axis %2d Value: %7d\n", axis, jinfo.AxisData[axis])
	}

	return
}

func main() {
	jsid := 0
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		jsid = i
	}

	js, jserr := joystick.Open(jsid)

	if jserr != nil {
		fmt.Println(jserr)
		return
	}

	ticker := time.NewTicker(time.Millisecond * 40)

	for range ticker.C {
		fmt.Printf("Joystick Name: %s Axis Count: %d Button Count: %d\n", js.Name(), js.AxisCount(), js.ButtonCount())
		readJoystick(js)
	}
}
