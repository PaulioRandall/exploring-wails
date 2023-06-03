package sqlite

import (
	"database/sql"
	"strings"

	"github.com/PaulioRandall/go-trackerr"
)

var (
	ErrOpeningDB = trackerr.New("Could not open database")
	ErrInvalidDB = trackerr.New("Not a valid database")
	ErrClosingDB = trackerr.New("Could not close database")
)

func Open(file string) (*sqliteDB, error) {
	fileURL := file + "?" + strings.Join(sqlitePragma, "&")

	conn, e := sql.Open("sqlite3", fileURL)
	if e != nil {
		return nil, ErrSQLite.CausedBy(e, ErrOpeningDB)
	}

	db := &sqliteDB{
		conn: conn,
	}

	e = db.validate()
	if e != nil {
		db.Close()
		return nil, ErrSQLite.CausedBy(e)
	}

	// TODO: query to check that it is actually a SQLite database
	//       because SQLite will open anything and won't tell you its not a DB
	//       until you try to do something.

	return db, nil
}

func (db *sqliteDB) validate() error {
	// TODO
	return ErrInvalidDB
}

func (db *sqliteDB) Close() error {
	e := db.conn.Close()
	if e != nil {
		return ErrSQLite.CausedBy(e, ErrClosingDB)
	}
	return nil
}
