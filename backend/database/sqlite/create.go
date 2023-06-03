package sqlite

import (
	// 	"database/sql"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/sourcery/backend/database"
)

var (
	ErrCreatingDB     = trackerr.New("Could not create database")
	ErrCreatingTables = trackerr.New("Could not create database tables")
)

func Create(file string) error {
	// TODO
	return ErrSQLite.CausedBy(trackerr.ErrTodo, ErrCreatingDB)
}

func (db *sqliteDB) createTables(task database.Task) error {
	// TODO
	return ErrSQLite.CausedBy(trackerr.ErrTodo, ErrCreatingTables)
}
