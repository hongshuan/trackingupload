//File: main.go

package main

import (
    "bookstore/models"
    "fmt"
    "net/http"
)

func main() {
    models.InitDB("postgres://user:pass@localhost/bookstore")

    http.HandleFunc("/books", booksIndex)
    http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := models.AllBooks()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}

//File: models/db.go

package models

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}

//File: models/books.go

package models

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func AllBooks() ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}

// Dependency injection
/*
type Env struct {
    db *sql.DB
    logger *log.Logger
    templates *template.Template
}
*/

// File: main.go

package main

import (
    "bookstore/models"
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

type Env struct {
    db *sql.DB
}

func main() {
    db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }
    env := &Env{db: db}

    http.HandleFunc("/books", env.booksIndex)
    http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := models.AllBooks(env.db)
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}

//File: models/db.go

package models

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}

//File: models/books.go

package models

import "database/sql"

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func AllBooks(db *sql.DB) ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}

//File: main.go

package main

import (
    "bookstore/models"
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

type Env struct {
    db *sql.DB
}

func main() {
    db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }
    env := &Env{db: db}

    http.Handle("/books", booksIndex(env))
    http.ListenAndServe(":3000", nil)
}

func booksIndex(env *Env) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            http.Error(w, http.StatusText(405), 405)
            return
        }
        bks, err := models.AllBooks(env.db)
        if err != nil {
            http.Error(w, http.StatusText(500), 500)
            return
        }
        for _, bk := range bks {
            fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
        }
    })
}

//File: main.go

package main

import (
    "fmt"
    "log"
    "net/http"
    "bookstore/models"
)

type Env struct {
    db models.Datastore
}

func main() {
    db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }

    env := &Env{db}

    http.HandleFunc("/books", env.booksIndex)
    http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := env.db.AllBooks()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}

//File: models/db.go

package models

import (
    _ "github.com/lib/pq"
    "database/sql"
)

type Datastore interface {
    AllBooks() ([]*Book, error)
}

type DB struct {
    *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}

//File: models/books.go

package models

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func (db *DB) AllBooks() ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}

package main

import (
    "bookstore/models"
    "context"
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

type ContextInjector struct {
    ctx context.Context
    h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func main() {
    db, err := models.NewDB("postgres://user:pass!@localhost/bookstore")
    if err != nil {
        log.Panic(err)
    }
    ctx := context.WithValue(context.Background(), "db", db)

    http.Handle("/books", &ContextInjector{ctx, http.HandlerFunc(booksIndex)})
    http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }

    db, ok := r.Context().Value("db").(*sql.DB)
    if !ok {
        http.Error(w, "could not get database connection pool from context", 500)
        return
    }

    bks, err := models.AllBooks(db)
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
    }
}

//File: models/db.go

package models

import (
    "database/sql"

    _ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}

//File: models/books.go

package models

import (
    "database/sql"
)

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

func AllBooks(db *sql.DB) ([]*Book, error) {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}
