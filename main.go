package main

import (
	"encoding/json"
	"os/exec"
	"time"

	"github.com/pquerna/otp/totp"

	"fmt"
	"io/ioutil"
	"os"
)

const (
	filePath string = "data.json"
)

func main() {
	for {
		clearScreen()
		// Read json and loop through codes
		entries := readJSONFile(filePath)
		for _, entry := range entries {
			// Generate and display codes
			code, err := totp.GenerateCode(entry.Secret, time.Now())
			if err != nil {
				panic(err)
			}
			fmt.Println(fmt.Sprintf("%v:\t\t%v", entry.Name, code))
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

// clearScreen runs a clear screen for each time the generation is run
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("\033[2J")
}

// Entry is an individual 2FA code to generate
type Entry struct {
	Name   string
	Secret string
}
