package mala

import (
	"database/sql"
	"io/ioutil"
	"testing"

	_ "github.com/lib/pq"
)

func initDB(t *testing.T) *sql.DB {
	// postgres://user:password@host/db_name?sslmode=disable&port=5432
	bytes, err := ioutil.ReadFile("z_test.conf")
	if err != nil {
		t.Fatal(err)
	}
	db, err := sql.Open("postgres", string(bytes))
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestBeginner(t *testing.T) {
	db := initDB(t)
	t.Log(db)
}
