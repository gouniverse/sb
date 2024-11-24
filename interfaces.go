package sb

import (
	"context"
	"database/sql"
)

type BuilderInterface interface {
	Column(column Column) BuilderInterface
	Create() string
	CreateIfNotExists() string
	CreateIndex(indexName string, columnName ...string) string
	Delete() string
	Drop() string
	DropIfExists() string
	Insert(columnValuesMap map[string]string) string
	GroupBy(groupBy GroupBy) BuilderInterface
	Limit(limit int64) BuilderInterface
	Offset(offset int64) BuilderInterface
	OrderBy(columnName string, sortDirection string) BuilderInterface

	Select(columns []string) string
	Table(name string) BuilderInterface
	Update(columnValues map[string]string) string
	View(name string) BuilderInterface
	ViewColumns(columns []string) BuilderInterface
	ViewSQL(sql string) BuilderInterface
	Where(where Where) BuilderInterface

	TableRename(oldTableName string, newTableName string) (string, error)
	TableColumnAdd(tableName string, column Column) (string, error)
	TableColumnRename(tableName, oldColumnName, newColumnName string) (string, error)
}

type DatabaseInterface interface {
	// DB the database connection
	DB() *sql.DB

	// Type the database type, i.e. "mssql", "mysql", "postgres", "sqlite"
	Type() string

	// BeginTransaction starts a transaction
	BeginTransaction() (err error)

	// BeginTransactionWithContext starts a transaction with context
	BeginTransactionWithContext(ctx context.Context, opts *sql.TxOptions) (err error)

	// Close closes the database
	Close() (err error)

	// CommitTransaction commits the transaction
	CommitTransaction() (err error)

	// DebugEnable enables or disables debug
	DebugEnable(debug bool)

	// ExecInTransaction executes a function in a transaction
	ExecInTransaction(fn func(d *Database) error) (err error)

	// Exec executes a query
	Exec(sqlStr string, args ...any) (sql.Result, error)

	IsMssql() bool
	IsMysql() bool
	IsPostgres() bool
	IsSqlite() bool
	SqlLog() []map[string]string
	SqlLogEmpty()
	SqlLogLen() int
	SqlLogEnable(enable bool)
	SqlLogShrink(leaveLast int)
	Open() (err error)
	Query(sqlStr string, args ...any) (*sql.Rows, error)
	RollbackTransaction() (err error)
	SelectToMapAny(sqlStr string, args ...any) ([]map[string]any, error)
	SelectToMapString(sqlStr string, args ...any) ([]map[string]string, error)

	// Tx the transaction
	Tx() *sql.Tx
}
