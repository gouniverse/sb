package sb

import (
	"testing"
	"time"
)

func TestDialects(t *testing.T) {
	if DIALECT_MSSQL != "mssql" {
		t.Fatal(`DIALECT_MSSQL must be "mssql"`)
	}

	if DIALECT_MYSQL != "mysql" {
		t.Fatal(`DIALECT_MYSQL must be "mysql"`)
	}

	if DIALECT_POSTGRES != "postgres" {
		t.Fatal(`DIALECT_POSTGRES must be "postgres"`)
	}

	if DIALECT_SQLITE != "sqlite" {
		t.Fatal(`DIALECT_SQLITE must be "sqlite"`)
	}
}

func TestColumnAttributes(t *testing.T) {
	if COLUMN_ATTRIBUTE_AUTO != "auto" {
		t.Fatal(`COLUMN_ATTRIBUTE_AUTO must be "auto"`)
	}

	if COLUMN_ATTRIBUTE_DECIMALS != "decimals" {
		t.Fatal(`COLUMN_ATTRIBUTE_DECIMALS must be "decimals"`)
	}

	if COLUMN_ATTRIBUTE_LENGTH != "length" {
		t.Fatal(`COLUMN_ATTRIBUTE_LENGTH must be "length"`)
	}

	if COLUMN_ATTRIBUTE_NULLABLE != "nullable" {
		t.Fatal(`COLUMN_ATTRIBUTE_NULLABLE must be "nullable"`)
	}

	if COLUMN_ATTRIBUTE_PRIMARY != "primary" {
		t.Fatal(`COLUMN_ATTRIBUTE_PRIMARY must be "primary"`)
	}

}

func TestColumnTypes(t *testing.T) {
	if COLUMN_TYPE_BLOB != "blob" {
		t.Fatal(`COLUMN_TYPE_BLOB must be "blob"`)
	}

	if COLUMN_TYPE_DATE != "date" {
		t.Fatal(`COLUMN_TYPE_DATE must be "date"`)
	}

	if COLUMN_TYPE_DATETIME != "datetime" {
		t.Fatal(`COLUMN_TYPE_DATETIME must be "datetime"`)
	}

	if COLUMN_TYPE_DECIMAL != "decimal" {
		t.Fatal(`COLUMN_TYPE_DECIMAL must be "decimal"`)
	}

	if COLUMN_TYPE_FLOAT != "float" {
		t.Fatal(`COLUMN_TYPE_FLOAT must be "float"`)
	}

	if COLUMN_TYPE_INTEGER != "integer" {
		t.Fatal(`COLUMN_TYPE_INTEGER must be "integer"`)
	}

	if COLUMN_TYPE_STRING != "string" {
		t.Fatal(`COLUMN_TYPE_STRING must be "string"`)
	}

	if COLUMN_TYPE_TEXT != "text" {
		t.Fatal(`COLUMN_TYPE_TEXT must be "text"`)
	}

	if COLUMN_TYPE_LONGTEXT != "longtext" {
		t.Fatal(`COLUMN_TYPE_LONGTEXT must be "longtext"`)
	}

}

func TestCommon(t *testing.T) {
	if YES != "yes" {
		t.Fatal(`YES must be "yes"`)
	}

	if NO != "no" {
		t.Fatal(`NO must be "no"`)
	}

}

func TestNullDate(t *testing.T) {
	if NULL_DATE != "0002-01-01" {
		t.Fatal(`NULL_DATE must be "0002-01-01", found: `, NULL_DATE)
	}

	layout := "2006-01-02"

	tm, err := time.Parse(layout, NULL_DATE)

	if err != nil {
		t.Fatal(err)
	}

	if tm.IsZero() {
		t.Fatal(`NULL_DATE must not be zero`)
	}
}

func TestNullDateTime(t *testing.T) {
	if NULL_DATETIME != "0002-01-01 00:00:00" {
		t.Fatal(`NULL_DATETIME must be "0002-01-01 00:00:00", found: `, NULL_DATETIME)
	}

	layout := "2006-01-02 15:04:05"

	tm, err := time.Parse(layout, NULL_DATETIME)

	if err != nil {
		t.Fatal(err)
	}

	if tm.IsZero() {
		t.Fatal(`NULL_DATETIME must not be zero`)
	}
}
