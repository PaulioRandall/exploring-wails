package sqlite

import (
	// "database/sql"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/sourcery/backend/database"
)

var (
	ErrAddingTask = trackerr.New("Failed to add task to database")
)

func (db *sqliteDB) AddTask(task database.Task) error {
	// TODO
	return ErrSQLite.ContextFor(ErrAddingTask, trackerr.ErrTodo)
}
