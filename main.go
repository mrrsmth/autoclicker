package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo" // This library will help to click
	"github.com/robotn/gohook"  // This library will help to listen for the keyboard
)

func main() {
	fmt.Println("Auto Clicker")
	autoClickerActive := false // Variable for checking if auto clicker is active or not
	go func() {                // An infinite loop that runs concurrently and it will check whether if the auto clicker is active or not, if so, then it will click
		for range time.Tick(time.Millisecond * 5) { // Using time.Tick instead of time.Sleep to simplify the code
			if autoClickerActive {
				robotgo.Click("left", false) // "left" tells that I need it to left click. false tells that we need single click and not a double click
			}
		}
	}()
	fmt.Println("Press F6 to Enable/Disable the Auto Clicker")
	for keyEvent := range hook.Start() { // Using range to read events from the hook
		if keyEvent.Kind > hook.KeyDown && keyEvent.Kind < hook.KeyUp { // Checking for a key press
			if keyEvent.Rawcode == 117 { // Checking if the rawcode matches with the rawcode of the F6 key (117)
				autoClickerActive = !autoClickerActive // Inverting autoClickerActive's value
				if autoClickerActive {
					fmt.Println("Enabled")
				} else {
					fmt.Println("Disabled")
				}
			}
		}
	}
}
