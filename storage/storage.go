package storage

import (
	"bufio"
	"fmt"
	"os"

	"encoding/json"

	"github.com/EnesDemirtas/steely-will/core"
)

const filename = "steely_will.json"

type Storage struct {
	Habits []core.Habit `json:"habits"`
}

func New() *Storage {
	return &Storage{
		Habits: []core.Habit{},
	}
}

func (s *Storage) Store(h core.Habit) error {
	file, err := s.CreateOrOpen()
	if err != nil {
		return err
	}
	defer file.Close()

	habitStr, err := json.Marshal(&h)
	if err != nil {
		return err
	}

	habitStr = append(habitStr, '\n')
	
	if _, err := file.Write(habitStr); err != nil {
		return err
	}
	return nil
}

func (s *Storage) LoadAll() ([]core.Habit, error) {
	file, err := s.CreateOrOpen()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	var habits []core.Habit

	for _, line := range fileLines {
		var habit core.Habit
		if err := json.Unmarshal([]byte(line), &habit); err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}

	return habits, nil
}

func (s *Storage) CreateOrOpen() (*os.File, error) {
	fileMode := os.O_APPEND|os.O_CREATE|os.O_RDWR

	file, err := os.OpenFile(filename, fileMode, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	
	return file, nil
}