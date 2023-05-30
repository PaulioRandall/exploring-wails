package main

import (
	"context"

	"github.com/PaulioRandall/go-trackerr"

	"github.com/PaulioRandall/sourcery/backend/database"
	"github.com/PaulioRandall/sourcery/backend/files"
)

type DB interface {
	Close() error
	AddTask(task database.Task) error
}

type App struct {
	ctx context.Context
	db  DB
}

var (
	ErrDatabaseNotOpen = trackerr.Track("Database not open")
)

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
	return squashResult(files.ListFilesInDir(dir))
}

func (a *App) ToAbsPath(path string) (string, error) {
	return files.ToAbsPath(path)
}

func (a *App) OpenDatabase(file string) (any, error) {
	if a.db != nil {
		return nil, nil
	}

	var e error
	a.db, e = database.OpenSQLite(file)
	if e != nil {
		return nil, squashIfError(e)
	}

	return nil, nil
}

func (a *App) CloseDatabase() (any, error) {
	if a.db == nil {
		return nil, nil
	}

	e := a.db.Close()
	if e != nil {
		return nil, squashIfError(e)
	}

	a.db = nil
	return nil, nil
}

func (a *App) AddTask(task database.Task) (any, error) {
	if a.db == nil {
		return nil, ErrDatabaseNotOpen
	}

	return nil, a.db.AddTask(task)
}

func squashResult[T any](v T, e error) (T, error) {
	if e != nil {
		return v, trackerr.Squash(e)
	}
	return v, nil
}

func squashIfError(e error) error {
	if e != nil {
		return trackerr.Squash(e)
	}
	return nil
}
