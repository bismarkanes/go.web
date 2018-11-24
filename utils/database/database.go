package database

import (
  "fmt"
  "github.com/jmoiron/sqlx"
)

var (
  // dBConn is the connection handle for the database
  dBConn *sqlx.DB
)

// SetupDatabases initialize database
func SetupDatabases(dialect string, username string, password string, hostname string, port string, database string) error {
  db, err := sqlx.Open(dialect, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8", username, password, hostname, port, database))
  if err != nil {
    return err
  }

  dBConn = db
  return nil
}

// GetHandle get database handle
func GetHandle() *sqlx.DB {
  return dBConn
}
