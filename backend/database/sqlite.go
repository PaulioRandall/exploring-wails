package database

import (
	"database/sql"
	"strings"

	"github.com/PaulioRandall/go-trackerr"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrSQLite = trackerr.New("SQLite database error")

	ErrCreatingDB = trackerr.New("Could not create database")
	ErrOpeningDB  = trackerr.New("Could not open database")
	ErrClosingDB  = trackerr.New("Could not close database")

	ErrAddingTask = trackerr.New("Failed to add task to database")

	// See (SQLite) https://www.sqlite.org/pragma.html
	// See (go-sqlite3) https://github.com/mattn/go-sqlite3#connection-string
	sqlitePragma = []string{
		// 5x the default page & cache sizes to speed things up (I'm on desktop)
		"_page_size=5120",
		"_cache_size=10000",

		// Turn off stuff that slows inserts down, just while developing.
		"_journal=MEMORY",
		//"_foreign_keys=OFF",
		//"_ignore_check_constraints=ON",
		"_sync=OFF",
	}
)

type sqliteDB struct {
	conn *sql.DB
}

func NewSQLite(file string) (*sqliteDB, error) {
	return nil, ErrSQLite.CausedBy(ErrCreatingDB.CausedBy(trackerr.ErrTodo))
}

func OpenSQLite(file string) (*sqliteDB, error) {
	fileURL := file + "?" + strings.Join(sqlitePragma, "&")

	conn, e := sql.Open("sqlite3", fileURL)
	if e != nil {
		return nil, ErrSQLite.CausedBy(ErrOpeningDB.CausedBy(e))
	}

	db := &sqliteDB{
		conn: conn,
	}

	// TODO: query to check that it is actually a SQLite database
	//       because SQLite will open anything and won't tell you its not a DB
	//       until you try to do something.

	return db, nil
}

func (db *sqliteDB) Close() error {
	e := db.conn.Close()
	if e != nil {
		return ErrSQLite.CausedBy(ErrClosingDB.CausedBy(e))
	}
	return nil
}

func (db *sqliteDB) AddTask(task Task) error {
	// TODO
	return ErrSQLite.CausedBy(ErrAddingTask.CausedBy(trackerr.ErrTodo))
}
