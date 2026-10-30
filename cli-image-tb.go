package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var assetsDir = "cheetsheets"

func isCheatSheetsFolder() bool {
	info, err := os.Stat(assetsDir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Error: cheetsheet tool broken! Cheat-sheets folder does not exist.")

			return false
		}

		fmt.Println("Error: cheetsheet tool broken!", err)

		return false
	}

	if !info.IsDir() {
		fmt.Println("Error: cheetsheet tool broken! The cheat sheets should be stored in a folder, how is this not a folder!?")

		return false
	}

	return true
}

func printAvailable() {
	if !isCheatSheetsFolder() {
		return
	}

	entries, err := os.ReadDir(assetsDir)

	if err != nil {
		fmt.Println("Error: cheetsheet tool broken! Cannot read in the cheat sheets folder", err)
	}

	fmt.Println("Available cheat sheets: \n")

	for _, entry := range entries {
		if entry.IsDir() { // Why would a folder be here?!.. anyway..
			continue
		}

		fmt.Println("*", strings.ToLower(strings.Replace(entry.Name(), ".txt", "", 1)))
	}

	fmt.Println("")
}

func printCheatSheet(cheatSheetName string) {
	if !isCheatSheetsFolder() {
		return
	}

	data, err := os.ReadFile(filepath.Join(assetsDir, cheatSheetName+".txt"))

	if err != nil {
		fmt.Println("Error: Cannot read the cheatsheet!", err)

		return
	}

	const (
		ColorReset = "\033[0m"
		ColorCyan  = "\033[36m"
	)

	fmt.Printf("%s==== %s ====%s\n", ColorCyan, cheatSheetName+" cheat-sheet", ColorReset)
	fmt.Println("\n" + string(data))
	fmt.Printf("\n%s================%s\n", ColorCyan, ColorReset)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide at least one argument! \n For example: cli-image-tb help")

		return
	}

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	var firstArg string = os.Args[1]

	firstArg = strings.ToLower(firstArg)

	if firstArg == "help" {
		fmt.Println("//TODO")

		return
	}

	_, err := strconv.Atoi(firstArg)

	if err != nil {
		fmt.Println("Error: The thumbnail size should be a valid integer!")

		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Error: Please insert the directory for the images!")
	}

	if len(os.Args) < 4 {
		fmt.Println("Error: There should be at least one url to an image!")
	}

	return

	printCheatSheet(firstArg)
}
