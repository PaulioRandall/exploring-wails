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

func New(file string) (*sqliteDB, error) {
	// TODO
	return nil, ErrSQLite.ContextFor(ErrCreatingDB, trackerr.ErrTodo)
}

func (db *sqliteDB) createTables(task database.Task) error {
	// TODO
	return ErrSQLite.ContextFor(ErrCreatingTables, trackerr.ErrTodo)
}
