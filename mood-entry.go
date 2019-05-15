package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const dateFormat string = "2 1 2006"

var BadRecord error = errors.New("Bad record.")

type MoodEntry struct {
	Rating uint64
	Date   time.Time
}

func FromCSVRecord(record []string) (*MoodEntry, error) {
	if len(record) != 2 {
		return nil, BadRecord
	}
	time, err := time.Parse(dateFormat, record[0])
	if err != nil {
		return nil, err
	}
	rating, err := strconv.ParseUint(record[1], 10, 64)
	if err != nil {
		return nil, err
	}
	entry := &MoodEntry{
		Rating: rating,
		Date:   time,
	}
	return entry, nil
}

func (m MoodEntry) String() string {
	return fmt.Sprintf("%s: %d/10",
		m.Date.Format(dateFormat),
		m.Rating)
}

func (m MoodEntry) CSVFormat() []string {
	return []string{m.Date.Format(dateFormat), fmt.Sprintf("%d", m.Rating)}
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

func LoadCSV(moodFile string) ([]MoodEntry, error) {
	var entries []MoodEntry
	file, err := os.OpenFile(moodFile,
		os.O_RDONLY,
		0755)
	if err != nil {
		return entries, err
	}

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return entries, err
		}
		entry, err := FromCSVRecord(record)
		if err != nil {
			return entries, err
		}
		entries = append(entries, *entry)
	}
	return entries, nil
}
