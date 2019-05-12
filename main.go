package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

var dataDirectory string = ""

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	dataDirectory = filepath.Join(user.HomeDir, ".local", "mood-tracker")
	saveDirExists, err := exists(dataDirectory)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !saveDirExists {
		fmt.Println("Gotta make the config")
		os.Exit(-1)
	}
	fmt.Println("All set up!")
	fmt.Println("Hi!")
}
