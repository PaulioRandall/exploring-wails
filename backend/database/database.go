package database

import (
	"fmt"
)

// Datastore interface is exactly what it says on the tin.
//
// The API separates what can be stored from how it is stored.
type Datastore interface {
	AddTask(task Task) error
	Close()
}

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (t Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Text)
}
