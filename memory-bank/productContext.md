# Product Context: Simplified SQL Builder (SB)

## Problem
Many Go developers require a way to programmatically generate SQL queries, but existing SQL builder libraries can be overly complex or offer features beyond what's needed for basic tasks. This leads to a steeper learning curve and increased project dependencies.

## Solution
SB provides a simplified and lightweight SQL builder that focuses on the most common SQL operations. It offers a fluent interface for constructing queries, making it easy to generate SQL code in a readable and maintainable way. The database wrapper simplifies transaction management, reducing boilerplate code.

## How it Should Work
- The library should provide functions to build SQL queries for creating tables, dropping tables, inserting data, deleting data, selecting data, creating views, and dropping views.
- The API should be intuitive and easy to use, with clear naming conventions.
- The generated SQL should be compatible with MySQL, PostgreSQL, and SQLite.
- The database wrapper should seamlessly integrate with the standard `database/sql` package, allowing developers to use existing database connections.
- Transaction management should be simplified through `BeginTransaction`, `CommitTransaction`, and `RollbackTransaction` methods.

## User Experience Goals
- Developers can quickly learn and use the library without extensive documentation.
- SQL queries can be generated with minimal code.
- Transaction management is straightforward and less error-prone.
- The library integrates smoothly into existing Go projects.
