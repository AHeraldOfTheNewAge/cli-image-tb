package main

import (
	"errors"
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

func isValidDir(dir string) bool {
	info, err := os.Stat(dir)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return err == nil && info.IsDir()
}

func isValidImage(dir string, imagePath string) bool {
	fmt.Println(dir, imagePath)

	return true
}

func main() {
	switch {
	case len(os.Args) < 2:
		fmt.Println("Error: Please provide at least one argument! \n For example: cli-image-tb help")

		return
	case len(os.Args) < 3:
		fmt.Println("Error: Please insert the directory for the images!")

		return
	case len(os.Args) < 4:
		fmt.Println("Error: Please insert the directory for the thumbnail converted images!")

		return
	case len(os.Args) < 5:
		fmt.Println("Error: There should be at least one url to an image!")

		return
	}

	// Check the first argument, wich is either help, either the thumbnail size to get returned!
	var firstArg string = os.Args[1]

	firstArg = strings.ToLower(firstArg)

	if firstArg == "help" {
		fmt.Println("//TODO cli-image-tb [thumbnail-size] [dir] [destination-dir] [image-name]..")

		return
	}

	_, err := strconv.Atoi(firstArg)

	if err != nil {
		fmt.Println("Error: The thumbnail size should be a valid integer!")

		return
	}

	if !isValidDir(os.Args[2]) {
		fmt.Println("Error: Please insert a valid directory for the images to be converted!")

		return
	}

	if !isValidDir(os.Args[3]) {
		fmt.Println("Error: Please insert a valid directory as destionation for the converted images!")

		return
	}

	for i := 4; i < len(os.Args); i++ { // Check the images to be converted!!
		if !isValidImage(os.Args[2], os.Args[i]) {
			fmt.Println("Error: invalid image", os.Args[i])

			return
		}

		fmt.Println(os.Args[i], i)
	}
}
