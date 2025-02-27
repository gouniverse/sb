package sb

// func NewBuilder(dialect string) *Builder {
// 	var columnSQLGenerator ColumnSQLGenerator
// 	switch dialect {
// 	case DIALECT_MYSQL:
// 		columnSQLGenerator = MySQLColumnSQLGenerator{}
// 	case DIALECT_POSTGRES:
// 		columnSQLGenerator = PostgreSQLColumnSQLGenerator{}
// 	case DIALECT_SQLITE:
// 		columnSQLGenerator = SQLiteColumnSQLGenerator{}
// 	case DIALECT_MSSQL:
// 		columnSQLGenerator = MSSQLColumnSQLGenerator{}
// 	default:
// 		panic("unsupported dialect: " + dialect)
// 	}

// 	return &Builder{
// 		Dialect:            dialect,
// 		sql:                map[string]any{},
// 		sqlColumns:         []Column{},
// 		sqlGroupBy:         []GroupBy{},
// 		sqlLimit:           0,
// 		sqlOffset:          0,
// 		sqlOrderBy:         []OrderBy{},
// 		sqlTableName:       "",
// 		sqlViewName:        "",
// 		sqlViewColumns:     []string{},
// 		sqlViewSQL:         "",
// 		sqlWhere:           []Where{},
// 		columnSQLGenerator: columnSQLGenerator,
// 	}
// }
