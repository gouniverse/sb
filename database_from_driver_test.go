package sb

import (
	"testing"

	// _ "github.com/glebarez/go-sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func TestNewDatabaseFromDriver(t *testing.T) {
	db, err := NewDatabaseFromDriver("sqlite3", "file:test_newdatabase_from_driver.db?cache=shared&mode=memory")
	if err != nil {
		t.Fatal("Error must be NIL but got: ", err.Error())
	}
	if db == nil {
		t.Fatal("Database MUST NOT BE NIL")
	}
	if db.DB() == nil {
		t.Fatal("Database db field MUST NOT BE NIL")
	}
	if db.Tx() != nil {
		t.Fatal("Database tx field MUST BE NIL")
	}
}
