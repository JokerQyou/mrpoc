package main

import (
	"database/sql"
    "fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/mattn/go-sqlite3"
    "os"
)

type log struct{}

func (l *log) Verbose() bool {
    return true
}

func (l *log) Printf(f string, v ...interface{}) {
    fmt.Printf(f, v...)
}

func main() {
    rice.Debug = true

    box, err := rice.FindBox("./sql")
    if err != nil {
        panic(err)
    }

    sourceDriver, err := httpfs.New(box.HTTPBox(), "")
    if err != nil {
        panic(err)
    }

    //db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
    if _, err := os.Stat("./test.db"); err == nil {
        fmt.Println("Removing test.db")
        err = os.Remove("./test.db")
        if err != nil {
            panic(err)
        }
    }
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        panic(err)
    }

    sqlDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
    if err != nil {
        panic(err)
    }

    m, err := migrate.NewWithInstance("embedFs", sourceDriver, "sqlite", sqlDriver)
    if err != nil {
        panic(err)
    }

    m.Log = &log{}

    if err = m.Up(); err != nil && err != migrate.ErrNoChange {
        panic(err)
    }

    fmt.Println("Migration finished")
}
