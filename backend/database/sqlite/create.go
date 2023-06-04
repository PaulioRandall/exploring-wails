package sqlite

import (
	// 	"database/sql"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/sourcery/backend/database"
	"github.com/PaulioRandall/sourcery/backend/files"
)

var (
	ErrCreatingDB     = trackerr.New("Could not create database")
	ErrCreatingTables = trackerr.New("Could not create database tables")
)

func Create(file string) error {
	exists, e := files.DoesFileExist(file)
	if e != nil {
		return ErrCreatingDB.CausedBy(e)
	}

	if exists {
		return ErrCreatingDB.Because("File already exists")
	}

	return ErrSQLite.CausedBy(trackerr.ErrTodo, ErrCreatingDB)
}

func (db *sqliteDB) createTables(task database.Task) error {
	// TODO
	return ErrSQLite.CausedBy(trackerr.ErrTodo, ErrCreatingTables)
}
