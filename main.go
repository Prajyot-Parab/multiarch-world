package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
)

func main() {

	fmt.Println("Process started.")
	// Create a new spinner
	s := spinner.New(spinner.CharSets[26], 100*time.Millisecond)

	// Set the spinner prefix text
	s.Prefix = "Loading"

	// Start the spinner
	s.Start()

	// Simulate some time-consuming task
	time.Sleep(3 * time.Second)

	// Stop the spinner
	s.Stop()

	// Check the current platform
	currentOS := runtime.GOOS
	currentArch := runtime.GOARCH
	fmt.Printf("Running on %s/%s\n", currentOS, currentArch)

	// Print a message after the spinner
	fmt.Println("Process completed.")
	fmt.Println("Welcome to the Multi-Arch World")
}
