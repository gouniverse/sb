package sb

import (
	"context"
	"database/sql"
	"testing"

	"github.com/samber/lo"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gouniverse/base/database"
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func initMySQLWithTable(tableName string, columns []Column) (db *sql.DB, err error) {
	db, err = database.Open(database.Options().
		SetDatabaseType(database.DATABASE_TYPE_MYSQL).
		SetDatabaseHost("localhost").
		SetDatabasePort("33306").
		SetDatabaseName("test").
		SetUserName("test").
		SetPassword("test"))

	if err != nil {
		return nil, err
	}

	err = TableDropIfExists(db, tableName)

	if err != nil {
		return nil, err
	}

	err = TableCreate(db, tableName, columns)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initSQLiteWithTable(tableName string, columns []Column) (db *sql.DB, err error) {
	db, err = database.Open(database.Options().
		SetDatabaseType(database.DATABASE_TYPE_SQLITE).
		SetDatabaseName(":memory:"))

	if err != nil {
		return nil, err
	}

	err = TableDropIfExists(db, tableName)

	if err != nil {
		return nil, err
	}

	err = TableCreate(db, tableName, columns)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestTableColumnsMySQL(t *testing.T) {
	columns := _TestTableColumns_columns()

	db, err := initMySQLWithTable("test_table_columns", columns)

	defer db.Close()

	if err != nil {
		t.Fatal("Error must be NIL but got: ", err.Error())
	}

	columns, err = TableColumns(context.Background(), db, "test_table_columns", true)

	if err != nil {
		t.Fatal("Error must be NIL but got: ", err.Error())
	}

	if len(columns) != 9 {
		t.Fatal("Error must be 9 but got: ", len(columns))
	}

	expecteds := []struct {
		columnName   string
		columnType   string
		isPrimaryKey bool
		length       int
		decimals     int
		isNullable   bool
		Default      string
		isUnique     bool
	}{
		{"id", COLUMN_TYPE_STRING, true, 40, 0, false, "", false},
		{"title", COLUMN_TYPE_STRING, false, 100, 0, false, "", true},
		{"image", COLUMN_TYPE_BLOB, false, 0, 0, false, "", false},
		{"price", COLUMN_TYPE_DECIMAL, false, 0, 0, false, "", false},
		{"price_custom", COLUMN_TYPE_DECIMAL, false, 12, 10, false, "", false},
		{"short_description", COLUMN_TYPE_TEXT, false, 0, 0, false, "", false},
		{"long_description", COLUMN_TYPE_TEXT, false, 0, 0, false, "", false},
		{"created_at", COLUMN_TYPE_DATETIME, false, 0, 0, false, "", false},
		{"deleted_at", COLUMN_TYPE_DATETIME, false, 0, 0, true, "", false},
	}

	for _, expected := range expecteds {
		column, found := lo.Find(columns, func(column Column) bool {
			return column.Name == expected.columnName
		})

		if !found {
			t.Fatal("Error column '"+expected.columnName+"' must be found but got: ", found)
		}

		if column.Type != expected.columnType {
			t.Fatal("Error column '"+expected.columnName+"' type must be '"+expected.columnType+"' but got: ", column.Type)
		}
	}

	t.Log(columns)
}

func TestTableColumnsSQLite(t *testing.T) {
	columns := _TestTableColumns_columns()

	db, err := initSQLiteWithTable("test_table_columns", columns)

	defer db.Close()

	if err != nil {
		t.Fatal("Error must be NIL but got: ", err.Error())
	}

	columns, err = TableColumns(context.Background(), db, "test_table_columns", true)

	if err != nil {
		t.Fatal("Error must be NIL but got: ", err.Error())
	}

	if len(columns) != 9 {
		t.Fatal("Error must be 9 but got: ", len(columns))
	}

	expecteds := []struct {
		columnName   string
		columnType   string
		isPrimaryKey bool
		length       int
		decimals     int
		isNullable   bool
		Default      string
		isUnique     bool
	}{
		{"id", COLUMN_TYPE_STRING, true, 40, 0, false, "", false},
		{"title", COLUMN_TYPE_STRING, false, 100, 0, false, "", true},
		{"image", COLUMN_TYPE_BLOB, false, 0, 0, false, "", false},
		{"price", COLUMN_TYPE_DECIMAL, false, 0, 0, false, "", false},
		{"price_custom", COLUMN_TYPE_DECIMAL, false, 12, 10, false, "", false},
		{"short_description", COLUMN_TYPE_TEXT, false, 0, 0, false, "", false},
		{"long_description", COLUMN_TYPE_TEXT, false, 0, 0, false, "", false},
		{"created_at", COLUMN_TYPE_DATETIME, false, 0, 0, false, "", false},
		{"deleted_at", COLUMN_TYPE_DATETIME, false, 0, 0, true, "", false},
	}

	for _, expected := range expecteds {
		column, found := lo.Find(columns, func(column Column) bool {
			return column.Name == expected.columnName
		})

		if !found {
			t.Fatal("Error column '"+expected.columnName+"' must be found but got: ", found)
		}

		if column.Type != expected.columnType {
			t.Fatal("Error column '"+expected.columnName+"' type must be '"+expected.columnType+"' but got: ", column.Type)
		}
	}

	t.Log(columns)
}

func _TestTableColumns_columns() []Column {
	columns := []Column{
		{
			Name:       "id",
			Type:       COLUMN_TYPE_STRING,
			Length:     40,
			PrimaryKey: true,
		},
		{
			Name:   "title",
			Type:   COLUMN_TYPE_STRING,
			Length: 100,
			Unique: true,
		},
		{
			Name: "image",
			Type: COLUMN_TYPE_BLOB,
		},
		{
			Name: "price",
			Type: COLUMN_TYPE_DECIMAL,
		},
		{
			Name:     "price_custom",
			Type:     COLUMN_TYPE_DECIMAL,
			Length:   12,
			Decimals: 10,
		},
		{
			Name: "short_description",
			Type: COLUMN_TYPE_TEXT,
		},
		{
			Name: "long_description",
			Type: COLUMN_TYPE_TEXT,
		},
		{
			Name: "created_at",
			Type: COLUMN_TYPE_DATETIME,
		},
		{
			Name:     "deleted_at",
			Type:     COLUMN_TYPE_DATETIME,
			Nullable: true,
		},
	}

	return columns
}