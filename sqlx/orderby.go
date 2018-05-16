package sqlx

import (
	"bytes"
	"database/sql"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

////

type OrderByBuilder struct {
	buffer *bytes.Buffer
}

func newOrderByBuilder(existsSQL string, columns ...string) *OrderByBuilder {
	if len(columns) == 0 {
		panic("Columns is required for ORDER BY keyword")
	}
	ob := &OrderByBuilder{
		buffer: new(bytes.Buffer),
	}
	ob.buffer.WriteString(existsSQL)
	ob.buffer.WriteString(" ORDER BY ")
	ob.buffer.WriteString(strings.Join(Map(columns, EscapeName), ","))
	return ob
}

func (slf *OrderByBuilder) Column(column string) *OrderByBuilder {
	slf.buffer.WriteString(", ")
	slf.buffer.WriteString(EscapeName(column))
	return slf
}

func (slf *OrderByBuilder) ASC() *OrderByBuilder {
	slf.buffer.WriteString(" ASC")
	return slf
}

func (slf *OrderByBuilder) DESC() *OrderByBuilder {
	slf.buffer.WriteString(" DESC")
	return slf
}

func (slf *OrderByBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.SQL(), limit)
}

func (slf *OrderByBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf.SQL(), columns...)
}

func (slf *OrderByBuilder) SQL() string {
	return slf.buffer.String()
}

func (slf *OrderByBuilder) MakeSQL() string {
	return makeSQL(slf.buffer)
}

func (slf *OrderByBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.MakeSQL(), db)
}
