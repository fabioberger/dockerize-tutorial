package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20141219125215(txn *sql.Tx) {
	create := `CREATE TABLE posts (
		Id serial NOT NULL,
		Title varchar UNIQUE,
		Body varchar UNIQUE,
		Timestamp int NOT NULL
	);`
	if _, err := txn.Exec(create); err != nil {
		panic(err)
	}
	index := `CREATE UNIQUE INDEX ON posts (Id);`
	if _, err := txn.Exec(index); err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20141219125215(txn *sql.Tx) {
	if _, err := txn.Exec("DROP TABLE posts;"); err != nil {
		panic(err)
	}
}
