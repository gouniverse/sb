# System Patterns: Simplified SQL Builder (SB)

## Architecture
The library follows a builder pattern for constructing SQL queries. The `Builder` struct holds the state of the query being built, and methods are chained together to define the query's components (table, columns, where clauses, etc.).

The database wrapper uses the standard Go `database/sql` package for database interactions. It provides methods for executing queries and managing transactions.

## Key Technical Decisions
- Use a fluent interface for the SQL builder to improve readability and ease of use.
- Support multiple database dialects through a dialect-specific code generation or templating approach.
- Implement the database wrapper as a thin layer on top of the `database/sql` package to minimize overhead.
- Provide a consistent API for transaction management across different database drivers.

## Design Patterns
- **Builder Pattern:** Used for constructing SQL queries.
- **Facade Pattern:** The database wrapper provides a simplified interface to the `database/sql` package.

## Component Relationships
- `Builder` -> `Table` -> `Column` -> `Where`
- `Database` -> `sql.DB` (Go's standard database connection)
