package sb

// // MySQLMockDriver wraps a sqlmock.Driver to make it look like MySQL driver to reflection
// type mysqlMockDriver struct {
// 	mock sqlmock.Sqlmock
// 	db   *sql.DB
// }

// func (m *mysqlMockDriver) Open(name string) (driver.Conn, error) {
// 	conn, err := m.db.Driver().Open(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &mockConn{Conn: conn}, nil
// }

// type mockConn struct {
// 	driver.Conn
// }

// func (c *mockConn) Begin() (driver.Tx, error) {
// 	return c.Conn.Begin()
// }

// func (c *mockConn) Close() error {
// 	return c.Conn.Close()
// }

// func (c *mockConn) Prepare(query string) (driver.Stmt, error) {
// 	return c.Conn.Prepare(query)
// }

// type mockQueryableContext struct {
// 	context.Context
// 	db *sql.DB
// }

// func (m *mockQueryableContext) Queryable() *sql.DB {
// 	return m.db
// }

// func mockMysql(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	// Create our custom driver wrapper
// 	customDriver := &mysqlMockDriver{
// 		mock: mock,
// 		db:   db,
// 	}

// 	// Register it with a unique name
// 	driverName := "mysql_test_driver"
// 	sql.Register(driverName, customDriver)

// 	// Open a DB with our custom driver
// 	testDB, err := sql.Open(driverName, "sqlmock_db")
// 	if err != nil {
// 		t.Fatalf("failed to open mock database: %v", err)
// 	}

// 	// Verify our driver wrapper works as expected
// 	driverFullName := reflect.ValueOf(testDB.Driver()).Type().String()
// 	if !strings.Contains(driverFullName, "mysql") {
// 		t.Fatalf("expected driver to be MySQL, got %s", driverFullName)
// 	}
// 	return testDB, mock
// }

// func TestTableColumnExists(t *testing.T) {
// 	// Create a new mock database with MySQL driver
// 	db, mock := mockMysql(t)
// 	t.Cleanup(func() {
// 		db.Close()
// 	})

// 	// Set up expectation for Ping to ensure connection is established
// 	mock.ExpectPing()

// 	t.Logf("Expecting ping...")

// 	// Ping the database to establish a connection
// 	if err := db.Ping(); err != nil {
// 		t.Fatalf("ping failed: %v", err)
// 	}

// 	t.Logf("Ping successful...")

// 	// Set up expectation for MySQL query
// 	mock.ExpectQuery("SELECT 1 FROM information_schema.COLUMNS WHERE TABLE_NAME = ? AND COLUMN_NAME = ?").
// 		WithArgs("test_table", "test_column").
// 		WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

// 	ctx := database.Context(context.Background(), db)

// 	t.Logf("Checking if column exists...")

// 	exists, err := TableColumnExists(ctx, "test_table", "test_column")
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if !exists {
// 		t.Fatalf("expected column to exist, got false")
// 	}

// 	t.Logf("Column check successful...")

// 	// Verify all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("there were unfulfilled expectations: %s", err)
// 	}

// 	t.Logf("Expectations met...")
// }

// func TestTableColumnExists(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	// Set mock driver type to MySQL for testing
// 	mock.ExpectQuery("SELECT 1 FROM information_schema\\.COLUMNS WHERE TABLE_NAME = \\? AND COLUMN_NAME = \\?").
// 		WithArgs("test_table", "test_column").
// 		WillReturnRows(sqlmock.NewRows([]string{"1"}))

// 	ctx := database.Context(context.Background(), db)
// 	defer db.Close()

// 	type args struct {
// 		ctx        database.QueryableContext
// 		tableName  string
// 		columnName string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "MySQL - Column Exists",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "test_column",
// 			},
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name: "PostgreSQL - Column Exists",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "test_column",
// 			},
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name: "SQLite - Column Exists",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "test_column",
// 			},
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name: "MySQL - Column Does Not Exist",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "nonexistent_column",
// 			},
// 			want:    false,
// 			wantErr: false,
// 		},
// 		{
// 			name: "PostgreSQL - Column Does Not Exist",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "nonexistent_column",
// 			},
// 			want:    false,
// 			wantErr: false,
// 		},
// 		{
// 			name: "SQLite - Column Does Not Exist",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "nonexistent_column",
// 			},
// 			want:    false,
// 			wantErr: false,
// 		},
// 		{
// 			name: "Empty Table Name",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "",
// 				columnName: "test_column",
// 			},
// 			want:    false,
// 			wantErr: true,
// 		},
// 		{
// 			name: "Empty Column Name",
// 			args: args{
// 				ctx:        ctx,
// 				tableName:  "test_table",
// 				columnName: "",
// 			},
// 			want:    false,
// 			wantErr: true,
// 		},
// 		{
// 			name: "Unsupported Database Type",
// 			args: args{
// 				ctx:        database.Context(context.Background(), nil),
// 				tableName:  "test_table",
// 				columnName: "test_column",
// 			},
// 			want:    false,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := TableColumnExists(tt.args.ctx, tt.args.tableName, tt.args.columnName)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("TableColumnExists() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("TableColumnExists() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_tableColumnExists_MySQL(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	ctx := database.Context(context.Background(), db)

// 	mock.ExpectQuery("SELECT 1 FROM information_schema\\.COLUMNS WHERE TABLE_NAME = \\? AND COLUMN_NAME = \\?").
// 		WithArgs("test_table", "test_column").
// 		WillReturnRows(sqlmock.NewRows([]string{"1"}))

// 	exists, err := TableColumnExists(ctx, "test_table", "test_column")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if !exists {
// 		t.Errorf("expected true, got false")
// 	}

// 	mock.ExpectQuery("SELECT 1 FROM information_schema\\.COLUMNS WHERE TABLE_NAME = \\? AND COLUMN_NAME = \\?").
// 		WithArgs("test_table", "nonexistent_column").
// 		WillReturnRows(sqlmock.NewRows([]string{}))

// 	exists, err = TableColumnExists(ctx, "test_table", "nonexistent_column")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if exists {
// 		t.Errorf("expected false, got true")
// 	}

// }

// func Test_tableColumnExists_PostgreSQL(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	ctx := database.Context(context.Background(), db)

// 	mock.ExpectQuery("SELECT EXISTS \\(SELECT 1 FROM information_schema\\.columns WHERE table_name = $1 AND column_name = $2\\)").
// 		WithArgs("test_table", "test_column").
// 		WillReturnRows(sqlmock.NewRows([]string{"exists"}))

// 	exists, err := TableColumnExists(ctx, "test_table", "test_column")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if !exists {
// 		t.Errorf("expected true, got false")
// 	}

// 	mock.ExpectQuery("SELECT EXISTS \\(SELECT 1 FROM information_schema\\.columns WHERE table_name = $1 AND column_name = $2\\)").
// 		WithArgs("test_table", "nonexistent_column").
// 		WillReturnRows(sqlmock.NewRows([]string{}))

// 	exists, err = TableColumnExists(ctx, "test_table", "nonexistent_column")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if exists {
// 		t.Errorf("expected false, got true")
// 	}
// }

// func Test_tableColumnExists_SQLite(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// Set mock driver type to SQLite for testing
// 	mock.ExpectQuery("PRAGMA table_info\\(\\?\\)").
// 		WithArgs("test_table").
// 		WillReturnRows(sqlmock.NewRows([]string{"cid", "name", "type", "notnull", "dflt_value", "pk"}))

// 	ctx := database.Context(context.Background(), db)

// 	exists, err := TableColumnExists(ctx, "test_table", "test_column")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if !exists {
// 		t.Errorf("expected true, got false")
// 	}

// 	exists, err = TableColumnExists(ctx, "test_table", "nonexistent_column")
// 	if err != nil {
// 		t.Errorf("unexpected error: %v", err)
// 	}
// 	if exists {
// 		t.Errorf("expected false, got true")
// 	}
// }
