package database

import (
	"fmt"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (t Task) String() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Text)
}
