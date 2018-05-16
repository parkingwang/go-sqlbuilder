package sql

import (
	"bytes"
	"database/sql"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type SelectBuilder struct {
	columns  []string
	table    string
	distinct bool
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

func (slf *SelectBuilder) buffer() *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")

	if slf.distinct {
		buf.WriteString("DISTINCT ")
	}

	if len(slf.columns) == 0 {
		buf.WriteByte('*')
	} else {
		buf.WriteString(strings.Join(Map(slf.columns, EscapeName), ","))
	}

	buf.WriteString(" FROM ")
	buf.WriteString(EscapeName(slf.table))
	return buf
}

func (slf *SelectBuilder) OrderBy(columns ...string) *OrderBuilder {
	return newOrderBuilder(slf.buffer(), columns...)
}

func (slf *SelectBuilder) Where() *WhereBuilder {
	return newWhere(slf.buffer())
}

func (slf *SelectBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.buffer(), limit)
}

func (slf *SelectBuilder) SQL() string {
	return endpoint(slf.buffer())
}

func (slf *SelectBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
