package main

////https://github.com/go-sql-driver/mysql/
///https://github.com/go-sql-driver/mysql/wiki/Examples

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DbService is ...
type DbService interface {
	open()
	list() (bool, error)
	close()
}

type dbService struct{}

func (dbService) list() (bool, error) {
	log.Println("db service called")
	db, err := sql.Open("mysql", "root:mysqlgolang@/golangpoc")
	defer db.Close()
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	rows, err := db.Query("SELECT * FROM squarenum")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return true, nil
}

func (dbService) open() {
	log.Println("db service open....")
	db, err := sql.Open("mysql", "root:mysqlgolang@/golangpoc")
	defer db.Close()
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func (dbService) close() {
	log.Println("db service close....")
}
