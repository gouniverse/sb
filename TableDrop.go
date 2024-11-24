package sb

import (
	"database/sql"

	"github.com/gouniverse/base/database"
)

func TableDropSql(q database.Queryable, tableName string) string {
	databaseType := database.DatabaseType(q)

	return NewBuilder(databaseType).Table(tableName).Drop()
}

func TableDropIfExistsSql(db *sql.DB, tableName string) string {
	databaseType := database.DatabaseType(db)

	return NewBuilder(databaseType).Table(tableName).DropIfExists()
}

func TableDrop(db *sql.DB, tableName string) error {
	sqlTableDrop := TableDropSql(db, tableName)

	_, err := db.Exec(sqlTableDrop)

	return err
}

func TableDropIfExists(db *sql.DB, tableName string) error {
	sqlTableDrop := TableDropIfExistsSql(db, tableName)

	_, err := db.Exec(sqlTableDrop)

	return err
}
