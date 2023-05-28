package main

import (
	//"fmt"
	"context"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/sourcery/backend/database"
	"github.com/PaulioRandall/sourcery/backend/files"
)

var (
	ErrDatabaseNotOpen = trackerr.Track("Database not open")
)

type DB interface {
	Close() error
	AddTask(task database.Task) error
}

type App struct {
	ctx context.Context
	db  DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		a.db.Close()
		a.db = nil
	}
}

func (a *App) ListFilesInDir(dir string) ([]files.ReadOnlyFile, error) {
	return files.ListFilesInDir(dir)
}

func (a *App) ToAbsPath(path string) (string, error) {
	return files.ToAbsPath(path)
}

func (a *App) OpenDatabase(file string) error {
	if a.db != nil {
		return nil
	}

	var e error
	a.db, e = database.Open(file)
	return e
}

func (a *App) CloseDatabase() error {
	if a.db == nil {
		return nil
	}

	e := a.db.Close()
	if e != nil {
		return e
	}

	a.db = nil
	return nil
}

func (a *App) AddTask(task database.Task) error {
	if a.db == nil {
		return ErrDatabaseNotOpen
	}

	return a.db.AddTask(task)
}
