package models

import (
	"database/sql"
	"fmt"

	"github.com/fabioberger/dockerize-tutorial/config"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

var Db *gorp.DbMap

func Init() {
	// connect to postgres database
	// TODO: use a password here if needed
	fmt.Println("[database] Connecting to database...")
	db, err := sql.Open(config.Database.Driver, config.Database.Open)
	if err != nil {
		panic(err)
	}
	fmt.Println("[database] Connected successfully.")

	// construct a gorp DbMap
	Db = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'reminders' and
	// specifying that the Id property is an auto incrementing PK
	Db.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	if config.Env == "test" {
		// if we're in the test environment, clear the database on startup
		Db.TruncateTables()
	}
}
