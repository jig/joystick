// Simple program that displays the state of the specified joystick
//
//	go run joysticktest.go 2
//
// displays state of joystick id 2
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/0xcafed00d/joystick"
)

var (
	lastAxis  [10]int
	lastMoved = time.Now()
)

func readJoystick(js joystick.Joystick) error {
	jinfo, err := js.Read()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
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
		if lastAxis[axis] != jinfo.AxisData[axis] {
			lastMoved = time.Now()
			fmt.Printf("Axis %2d Value: %7d *\n", axis, jinfo.AxisData[axis])
			lastAxis[axis] = jinfo.AxisData[axis]
		} else {
			fmt.Printf("Axis %2d Value: %7d\n", axis, jinfo.AxisData[axis])
		}
	}
	if time.Since(lastMoved) > 1*time.Minute {
		fmt.Printf("Disconnected: %v\n", time.Since(lastMoved))
		return errors.New("Joystick disconnected")
	}
	fmt.Printf("Last change: %v\n", time.Since(lastMoved))
	return nil
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
		if err := readJoystick(js); err != nil {
			log.Fatal(err)
		}
	}
}
