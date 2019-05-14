package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type MoodEntry struct {
	Rating uint64
	Date   time.Time
}

func (m MoodEntry) String() string {
	return fmt.Sprintf("%s: %d/10",
		m.Date.Format("Mon 2/1/2006"),
		m.Rating)
}

func (m MoodEntry) CSVFormat() []string {
	return []string{m.Date.String(), fmt.Sprintf("%d", m.Rating)}
}

func (m MoodEntry) Save(moodFile string) error {
	file, err := os.OpenFile(moodFile,
		os.O_RDWR|os.O_APPEND|os.O_CREATE,
		0755)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	records := m.CSVFormat()
	err = writer.Write(records)
	return err
}
