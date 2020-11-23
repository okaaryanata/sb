package db

import (
	"database/sql"
	"log"
	"os"

	// init mysql
	_ "github.com/go-sql-driver/mysql"
)

//Connection hold connection and Transaction to mysql
type Connection struct {
	Server *sql.DB
	Tx     *sql.Tx
}

// Connect is function to connect to database mysql
func Connect() Connection {
	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE"))
	if err != nil {
		log.Print(err.Error())
	}
	return Connection{Server: db}
}

// Close is function to execute close
func (c Connection) Close() {
	c.Server.Close()
}

// Begin is function to execute begin
func (c *Connection) Begin() {
	s := c.Server
	tx, err := s.Begin()
	if err != nil {
		log.Print(err.Error())
		return
	}
	c.Tx = tx
}

// Commit is function to execute commit
func (c Connection) Commit() {
	err := c.Tx.Commit()
	if err != nil {
		log.Print(err.Error())
		return
	}
}

// Exec is function to execute exec (insert data) and return last insert id
func (c Connection) Exec(stmnt string, args ...interface{}) (int, error) {
	r, err := c.Tx.Exec(stmnt, args...)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Rollback is function to execute rollback if there is error
func (c Connection) Rollback() error {
	err := c.Tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

// Query is function to execute query sql rows
func (c Connection) Query(stmnt string, args ...interface{}) (*sql.Rows, error) {
	r, err := c.Tx.Query(stmnt)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// QueryRow is function to execute query and get single row
func (c Connection) QueryRow(stmnt string, args ...interface{}) *sql.Row {
	r := c.Tx.QueryRow(stmnt)
	return r
}
