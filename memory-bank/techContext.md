# Technical Context: Simplified SQL Builder (SB)

## Technologies Used
- Go programming language
- `database/sql` package (Go's standard database library)
- MySQL, PostgreSQL, and SQLite database drivers

## Development Setup
- Go development environment
- A suitable IDE or text editor
- Access to MySQL, PostgreSQL, or SQLite databases for testing

## Technical Constraints
- The library should have minimal dependencies to reduce project size and complexity.
- The generated SQL should be compatible with the target database dialects.
- The database wrapper should not introduce significant performance overhead.
- Transaction management should be reliable and consistent across different database drivers.

## Dependencies
- `database/sql` (Go's standard database library)
- Specific database drivers (e.g., `github.com/go-sql-driver/mysql` for MySQL)
