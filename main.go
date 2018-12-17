package main

import (
	"encoding/json"
	"time"

	"fmt"
	"io/ioutil"

	"github.com/inancgumus/screen"
	"github.com/pquerna/otp/totp"
)

const (
	filePath string = "data.json"
	format   string = "%-40s%s\n"
)

func main() {
	// Clears the screen
	screen.Clear()
	for {
		screen.MoveTopLeft()
		// Read json and loop through codes
		entries := readJSONFile(filePath)
		for _, entry := range entries {
			// Generate and display codes
			code, err := totp.GenerateCode(entry.Secret, time.Now())
			if err != nil {
				panic(err)
			}

			display := fmt.Sprintf(format, entry.Name, code)
			fmt.Printf("%v", display)
		}

		// Wait before regenerating codes
		time.Sleep(time.Second)
	}

}

// readJSONFile reads a file and converts the JSON to an Entry type
func readJSONFile(file string) []Entry {
	plan, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var data []Entry
	err = json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}

	return data
}

// Entry is an individual 2FA code to generate
type Entry struct {
	Name   string
	Secret string
}
