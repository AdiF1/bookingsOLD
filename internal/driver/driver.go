package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v4"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConns = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates db pool for Postgres
func ConnectSQL (dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConns)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDb(d)
	if err != nil {
		return nil, err
	} 
	return dbConn, nil
}

// testDb tries to ping the db
func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new db for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping() ; err != nil {
		return nil, err
	}

	return db, nil
}