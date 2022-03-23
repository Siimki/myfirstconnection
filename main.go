package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// This is a simple project for all of you to try out
	// idea is to test connecting, editing and submitting
	// changes to GitHub via Git | Maybe i need to do
	// something else for that to happen but lets hope not
	//   ~ kaneetz | v.1

	// Start a scanner
	scanner := bufio.NewScanner(os.Stdin)
	// print out first "Hello World"
	fmt.Println("Enter some text:")
	// scanner is doing the scanning thing
	scanner.Scan()
	// scanner input is saved as a string
	inpTxt := scanner.Text()
	// It knows what you typed! | Be nice :P
	fmt.Printf("You entered: %v", inpTxt)

}
        // Adding comment is enough, right?
