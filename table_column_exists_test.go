package sb

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gouniverse/base/database"
)

func TestTableColumnExists(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := database.Context(context.Background(), db)
	defer db.Close()

	type args struct {
		ctx        database.QueryableContext
		tableName  string
		columnName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "MySQL - Column Exists",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "test_column",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "PostgreSQL - Column Exists",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "test_column",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "SQLite - Column Exists",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "test_column",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "MySQL - Column Does Not Exist",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "nonexistent_column",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "PostgreSQL - Column Does Not Exist",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "nonexistent_column",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "SQLite - Column Does Not Exist",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "nonexistent_column",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Empty Table Name",
			args: args{
				ctx:        ctx,
				tableName:  "",
				columnName: "test_column",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Empty Column Name",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Nil Queryable",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "test_column",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Unsupported Database Type",
			args: args{
				ctx:        ctx,
				tableName:  "test_table",
				columnName: "test_column",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TableColumnExists(tt.args.ctx, tt.args.tableName, tt.args.columnName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TableColumnExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TableColumnExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tableColumnExists_MySQL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := database.Context(context.Background(), db)

	mock.ExpectQuery("SELECT 1 FROM information_schema\\.COLUMNS WHERE TABLE_NAME = \\? AND COLUMN_NAME = \\?").
		WithArgs("test_table", "test_column").
		WillReturnRows(sqlmock.NewRows([]string{"1"}))

	exists, err := TableColumnExists(ctx, "test_table", "test_column")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("expected true, got false")
	}

	mock.ExpectQuery("SELECT 1 FROM information_schema\\.COLUMNS WHERE TABLE_NAME = \\? AND COLUMN_NAME = \\?").
		WithArgs("test_table", "nonexistent_column").
		WillReturnRows(sqlmock.NewRows([]string{}))

	exists, err = TableColumnExists(ctx, "test_table", "nonexistent_column")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if exists {
		t.Errorf("expected false, got true")
	}

}

func Test_tableColumnExists_PostgreSQL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := database.Context(context.Background(), db)

	mock.ExpectQuery("SELECT EXISTS \\(SELECT 1 FROM information_schema\\.columns WHERE table_name = $1 AND column_name = $2\\)").
		WithArgs("test_table", "test_column").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}))

	exists, err := TableColumnExists(ctx, "test_table", "test_column")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("expected true, got false")
	}

	mock.ExpectQuery("SELECT EXISTS \\(SELECT 1 FROM information_schema\\.columns WHERE table_name = $1 AND column_name = $2\\)").
		WithArgs("test_table", "nonexistent_column").
		WillReturnRows(sqlmock.NewRows([]string{}))

	exists, err = TableColumnExists(ctx, "test_table", "nonexistent_column")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if exists {
		t.Errorf("expected false, got true")
	}
}

func Test_tableColumnExists_SQLite(t *testing.T) {
	db, err := initSQLiteWithTable("test_table", []Column{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := database.Context(context.Background(), db)

	exists, err := TableColumnExists(ctx, "test_table", "test_column")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("expected true, got false")
	}

	exists, err = TableColumnExists(ctx, "test_table", "nonexistent_column")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if exists {
		t.Errorf("expected false, got true")
	}
}
