package database

import (
	"fmt"
	"testing"
	"database/sql"
)

func testDatabaseConnection(t *testing.T) error {
	connStr := "cabontech dbname=taskmanager sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	
	fmt.Println("Database connection is working.")
	return nil
}

func main() {
	err := testDatabaseConnection(nil)
	if err != nil {
		fmt.Printf("Database connection test failed: %v\n", err)
		return
	}
}

