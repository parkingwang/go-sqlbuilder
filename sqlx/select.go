package sqlx

import (
	"bytes"
	"database/sql"
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
	buf.WriteString(EscapeName(slf.table))
	return buf
}

func (slf *SelectBuilder) Where(conditions Statement) *WhereBuilder {
	return newWhereWith(slf.SQL(), conditions)
}

func (slf *SelectBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.SQL(), limit)
}

func (slf *SelectBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.SQL(), columns...)
}

func (slf *SelectBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf.SQL(), columns...)
}

func (slf *SelectBuilder) SQL() string {
	return slf.build().String()
}

func (slf *SelectBuilder) MakeSQL() string {
	return makeSQL(slf.build())
}

func (slf *SelectBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.MakeSQL(), db)
}
