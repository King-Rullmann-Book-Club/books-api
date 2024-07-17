package books

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Transactor interface {
	GetBookById(id uint) (Book, error)

	Close()
}

var t Transactor

type transactor struct {
	db *sql.DB
}

func (s *transactor) Close() {
	s.db.Close()
}

func buildTransactor() Transactor {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}

	// migrations initalization
	sqlStmt := `create table if not exists migrations (id integer not null primary key autoincrement, name text unique);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	// add new migraions here in name, migration pairs
	migrations := []string{
		`books_1`,
		`create table books (id integer not null primary key autoincrement, title text);`,

		// TODO remove once data insertion is in place
		`books_2`,
		`insert into books (title) values("Perdido Street Station");`,
		`books_3`,
		`insert into books (title) values("The City & the City");`,
	}
	if len(migrations)%2 != 0 {
		log.Fatal("Invalid number of migrations")
	}
	// TODO load migrations from files

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	var migration_name string = ""
	for i, v := range migrations {
		if i%2 == 0 { // this is the name
			migration_name = v
		} else {
			rows, err := tx.Query(`select name from migrations where name = ?;`, migration_name)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			if rows.Next() {
				continue // migration alread ran, skip running statment
			}
			log.Default().Printf("Running migration %v with statment '%v'", migration_name, v)
			_, err = tx.Exec(v)
			if err != nil {
				log.Fatal(err)
			}
			_, err = tx.Exec("insert into migrations (name) values(?);", migration_name)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Could not run migrations %v", err)
	}
	return &transactor{db}
}

func NewTransactor() Transactor {
	if t != nil {
		return t
	}
	t = buildTransactor()
	return t
}

func CloseTransactor() {
	if t != nil {
		t.Close()
	}
	t = nil
}
