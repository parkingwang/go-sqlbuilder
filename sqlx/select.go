package sqlx

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type SelectBuilder struct {
	columns    []string
	table      string
	fromSelect SQLStatement
	distinct   bool
}

func Select(columns ...string) *SelectBuilder {
	return &SelectBuilder{
		columns: columns,
	}
}

func (slf *SelectBuilder) Distinct() *SelectBuilder {
	slf.distinct = true
	return slf
}

func (slf *SelectBuilder) From(table string) *SelectBuilder {
	slf.table = table
	return slf
}

func (slf *SelectBuilder) FromSelect(innerSelect SQLStatement) *SelectBuilder {
	slf.fromSelect = innerSelect
	return slf
}

func (slf *SelectBuilder) build() *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")

	if slf.distinct {
		buf.WriteString("DISTINCT ")
	}

	if len(slf.columns) == 0 {
		buf.WriteByte('*')
	} else {
		buf.WriteString(joinNames(slf.columns))
	}

	buf.WriteString(" FROM ")
	if nil != slf.fromSelect {
		buf.WriteByte('(')
		buf.WriteString(slf.fromSelect.Statement())
		buf.WriteByte(')')
	} else {
		buf.WriteString(EscapeName(slf.table))
	}
	return buf
}

func (slf *SelectBuilder) Where(conditions SQLStatement) *WhereBuilder {
	return newWhere(slf, conditions)
}

func (slf *SelectBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf, limit)
}

func (slf *SelectBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderBy(slf, columns...)
}

func (slf *SelectBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf, columns...)
}

func (slf *SelectBuilder) Statement() string {
	return slf.build().String()
}

func (slf *SelectBuilder) GetSQL() string {
	return endOfSQL(slf.build())
}

func (slf *SelectBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}
