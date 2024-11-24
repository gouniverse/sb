package sb

import (
	"database/sql"
	"reflect"
	"strings"
)

// DatabaseDriverName finds the driver name from database
func DatabaseDriverName(db *sql.DB) string {
	driverFullName := reflect.ValueOf(db.Driver()).Type().String()

	if strings.Contains(driverFullName, DIALECT_MYSQL) {
		return DIALECT_MYSQL
	}

	if strings.Contains(driverFullName, DIALECT_POSTGRES) || strings.Contains(driverFullName, "pq") {
		return DIALECT_POSTGRES
	}

	if strings.Contains(driverFullName, DIALECT_SQLITE) {
		return DIALECT_SQLITE
	}

	if strings.Contains(driverFullName, DIALECT_MSSQL) {
		return DIALECT_MSSQL
	}

	return driverFullName
}
