package sqlite

import (
	"database/sql"

	"github.com/PaulioRandall/go-trackerr"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrSQLite = trackerr.New("SQLite database error")

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
