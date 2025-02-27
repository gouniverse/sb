package sb

import (
	"strings"
)

type Where struct {
	Raw      string
	Column   string
	Operator string
	Type     string
	Value    string
	Children []Where
}

/**
 * Converts wheres to SQL
 * @param array $wheres
 * @return string
 */
func (b *Builder) whereToSql(wheres []Where) string {
	sql := []string{}
	for _, where := range wheres {
		if where.Raw != "" {
			sql = append(sql, where.Raw)
			continue
		}

		if where.Type == "" {
			where.Type = "AND"
		}

		if where.Column != "" {
			sqlSingle := b.whereToSqlSingle(where.Column, where.Operator, where.Value)

			if len(sql) > 0 {
				sql = append(sql, where.Type+" "+sqlSingle)
			} else {
				sql = append(sql, sqlSingle)
			}

		}
		// } else {
		// 	$_sql = array();
		// 	$all = $where['WHERE'];
		// 	for ($k = 0; k < count($all); k++) {
		// 		$w = $all[$k];
		// 		$sqlSingle = $this->whereToSqlSingle($w['COLUMN'], $w['OPERATOR'], $w['VALUE']);
		// 		if ($k == 0) {
		// 			$_sql[] = $sqlSingle;
		// 		} else {
		// 			$_sql[] = $w['TYPE'] . " " . $sqlSingle;
		// 		}
		// 	}
		// 	$_sql = (count($_sql) > 0) ? " (" . implode(" ", $_sql) . ")" : "";

		// 	if ($i == 0) {
		// 		$sql[] = $_sql;
		// 	} else {
		// 		$sql[] = $where['TYPE'] . " " . $_sql;
		// 	}
		// }
	}

	if len(sql) > 0 {
		return " WHERE " + strings.Join(sql, " ")
	}

	return ""
}

func (b *Builder) whereToSqlSingle(column, operator, value string) string {
	if operator == "==" || operator == "===" {
		operator = "="
	}
	if operator == "!=" || operator == "!==" {
		operator = "<>"
	}
	columnQuoted := b.quoteColumn(column)
	valueQuoted := b.quoteValue(value)

	sql := ""
	if b.Dialect == DIALECT_MYSQL {
		if value == "NULL" && operator == "=" {
			sql = columnQuoted + " IS NULL"
		} else if value == "NULL" && operator == "<>" {
			sql = columnQuoted + " IS NOT NULL"
		} else {
			sql = columnQuoted + " " + operator + " " + valueQuoted
		}
	}
	if b.Dialect == DIALECT_POSTGRES {
		if value == "NULL" && operator == "=" {
			sql = columnQuoted + " IS NULL"
		} else if value == "NULL" && operator == "<>" {
			sql = columnQuoted + " IS NOT NULL"
		} else {
			sql = columnQuoted + " " + operator + " " + valueQuoted
		}
	}
	if b.Dialect == DIALECT_SQLITE {
		if value == "NULL" && operator == "=" {
			sql = columnQuoted + " IS NULL"
		} else if value == "NULL" && operator == "<>" {
			sql = columnQuoted + " IS NOT NULL"
		} else {
			sql = columnQuoted + " " + operator + " " + valueQuoted
		}
	}
	return sql
}
