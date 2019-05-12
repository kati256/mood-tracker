package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

func CreateSaveDirectory() error {
	err := os.Mkdir(dataDirectory, 0700)
	if err != nil {
		return err
	}
	moodFile := filepath.Join(dataDirectory, "moods.csv")
	file, err := os.Create(moodFile)
	file.Close()
	return err
}

type MoodEntry struct {
	Rating uint64
	Date   time.Time
}

func (m MoodEntry) String() string {
	return fmt.Sprintf("%s: %d/10",
		m.Date.Format("Mon 2/1/2006"),
		m.Rating)
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
		err := CreateSaveDirectory()
		if err != nil {
			fmt.Printf("Failed to create directory structure.\n%v\n", err)
			os.Exit(-1)
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please rate your mood in scale of 1 to 10: ")
	ratestr, err := reader.ReadString('\n')
	ratestr = strings.TrimSpace(ratestr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
	rate, err := strconv.ParseUint(ratestr, 10, 8)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
	dailyMood := MoodEntry{
		Rating: rate,
		Date:   time.Now(),
	}
	fmt.Println(dailyMood.String())
}
