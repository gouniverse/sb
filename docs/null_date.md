# NULL Date Decision

## Comparison of NULL Date Options

| Criterion | 0001-01-01 00:00:00 | 0002-01-01 00:00:00 |
|-----------|---------------------|---------------------|
| ✅ Widely Supported | Yes in .NET, Python, PostgreSQL | Yes, slightly safer in some systems |
| ⚠️ Safer for Legacy/SQL Systems | No (MySQL may reject it) | Yes (avoids 0001 edge cases) |
| 💡 Clearly Not Realistic Date | Yes | Yes |
| 📅 One Year After MinValue | No | Yes (fewer "min value" bugs) |
| 🔄 Risk of Collision with Defaults | High (might overlap with MinValue defaults) | Low (less likely to be mistaken for MinValue) |

## Decision

**Use `0002-01-01` as the NULL date value** for the following reasons:

1. **Better Compatibility**: Avoids issues with MySQL and other systems that may reject `0001-01-01`
2. **Clearer Intent**: Less likely to be confused with system minimum values
3. **Reduced Risk**: Lower chance of collision with default or minimum date values in various systems
4. **Still Clearly Invalid**: Remains an obviously invalid date for most real-world applications

## Implementation

```go
// NULL_DATE represents an invalid/empty date value
// Using 0002-01-01 instead of 0001-01-01 for better compatibility with various databases
// and to avoid confusion with minimum date values
const NULL_DATE = "0002-01-01"
const NULL_DATETIME = "0002-01-01 00:00:00"