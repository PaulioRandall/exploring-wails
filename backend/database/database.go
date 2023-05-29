package database

import (
	"fmt"

	"github.com/PaulioRandall/go-trackerr"
)

var (
	ErrOpeningDatabase = trackerr.New("Could not open database")
	ErrClosingDatabase = trackerr.New("Could not close database")

	ErrAddingTask = trackerr.New("Failed to add task to database")
)

type database struct {
}

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (t Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Text)
}

func Open(file string) (*database, error) {
	// TODO
	return nil, ErrOpeningDatabase
}

func (db *database) Close() error {
	// TODO
	return ErrClosingDatabase
}

func (db *database) AddTask(task Task) error {
	// TODO
	return ErrAddingTask
}
