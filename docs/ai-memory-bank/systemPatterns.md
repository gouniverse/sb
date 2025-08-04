# System Patterns: Simplified SQL Builder (SB)

Last updated at: 2025-07-22

## Architecture
The library implements a fluent builder pattern for constructing type-safe SQL queries. The `Builder` struct maintains the query state, with method chaining for defining query components (tables, columns, conditions, etc.). The architecture is designed for extensibility and database dialect independence.

Key architectural components:
- `Builder`: Main entry point for constructing queries
- `Database`: Wrapper around `sql.DB` with enhanced functionality
- `Column`: Represents table columns with their properties
- `Where`: Handles query conditions
- Dialect-specific SQL generation for MySQL, PostgreSQL, and SQLite

## Key Technical Decisions
- **Fluent Interface**: Method chaining for building queries in a readable, expressive way
- **Dialect Support**: Native support for MySQL, PostgreSQL, and SQLite with consistent API
- **Minimal Dependencies**: Primarily relies on Go standard library, with minimal external dependencies
- **Transaction Management**: Comprehensive transaction support with `Begin`, `Commit`, and `Rollback` methods
- **Type Safety**: Strong typing for SQL components to prevent runtime errors
- **Performance**: Optimized for minimal overhead over raw SQL queries

## Design Patterns
- **Builder Pattern**: Core pattern for constructing complex SQL queries through method chaining
- **Facade Pattern**: Simplified interface over `database/sql` package
- **Strategy Pattern**: Different SQL generation strategies for various database dialects
- **Template Method**: Used in query building to handle common SQL patterns
- **Dependency Injection**: Database drivers and other dependencies are injected for testability

## Component Relationships
```
Builder
├── Table
│   ├── Column
│   │   ├── Data Type
│   │   ├── Constraints
│   │   └── Default Values
│   └── Indexes
├── Where (Conditions)
│   ├── AND/OR logic
│   ├── Comparison operators
│   └── Parameter binding
└── Query Operations
    ├── SELECT
    ├── INSERT
    ├── UPDATE
    ├── DELETE
    └── Schema Operations

Database
├── Connection Management
├── Transaction Handling
├── Query Execution
└── Result Processing
    └── Rows Scanning
```
