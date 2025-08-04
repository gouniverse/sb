# Project Brief: Simplified SQL Builder (SB)

Last updated at: 2025-07-22

## Goal
Develop and maintain a simplified SQL builder library in Go, focusing on essential SQL functionalities and ease of use.

## Core Requirements
- Provide a fluent interface for constructing SQL queries (CREATE, DROP, ALTER, TRUNCATE, INSERT, UPDATE, DELETE, SELECT, VIEW creation/deletion).
- Include a database wrapper to handle transactions transparently.
- Support MySQL, PostgreSQL, and SQLite dialects.
- Maintain comprehensive documentation within the memory bank.

## Target Audience
Go developers who need a lightweight SQL builder for basic database operations and prefer simplicity over extensive features.

## Success Metrics
- Library is easy to install and use (positive user feedback, low issue count related to usability).
- Core SQL building functionalities are implemented and well-tested.
- Database wrapper effectively manages transactions.
- Project is well-documented and understandable.

## Out of Scope
- Complex SQL query building features (e.g., advanced JOINs, subqueries, window functions).
- ORM-like functionalities (though basic CRUD operations are supported).
- Support for database schema migrations (though basic table/column operations are supported).
- Extensive database driver support beyond MySQL, PostgreSQL, and SQLite (MSSQL is also partially supported).
