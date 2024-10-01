package storage

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	// needed so that driver is installed

	_ "github.com/mattn/go-sqlite3"
)

type Transactor interface {
	// columnNames are columns to be retrieved, columnDests are output values that the column results are put into
	GetRecordById(tableName string, id uint, columnNames []string, columnDests ...any) error

	Close()
}

type transactor struct {
	db         *sql.DB
	tableRegex *regexp.Regexp
}

type fieldNameError struct {
	fieldName string
}

func (e fieldNameError) Error() string {
	return fmt.Sprintf("%s is a invalid field name", e.fieldName)
}

func (s *transactor) Close() {
	s.db.Close()
}

func (s *transactor) GetRecordById(tableName string, id uint, columnNames []string, columnDests ...any) error {
	// crude SQL injection protection
	if !s.tableRegex.Match([]byte(tableName)) {
		return fieldNameError{tableName}
	}
	var columnQs []string
	for _, x := range columnNames {
		if !s.tableRegex.Match([]byte(x)) {
			return fieldNameError{x}
		}
		columnQs = append(columnQs, x)
	}
	row := s.db.QueryRow(`select `+strings.Join(columnQs, ", ")+` from `+tableName+`  where id = ? limit 1;`, id)
	err := row.Scan(columnDests...)
	if err != nil {
		log.Default().Printf("GetRecordById erro %v", err)
		return err
	}
	return nil
}

func NewTransactor() Transactor {
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
	return &transactor{db, regexp.MustCompile("[a-zA-Z0-9_]+")}
}

func CloseTransactor(t Transactor) {
    t.Close()
}
