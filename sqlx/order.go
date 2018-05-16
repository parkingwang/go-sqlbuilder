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

type OrderBuilder struct {
	buffer *bytes.Buffer
}

func newOrderBuilder(existsSQL string, columns ...string) *OrderBuilder {
	if len(columns) == 0 {
		panic("Columns is required for ORDER BY keyword")
	}
	ob := &OrderBuilder{
		buffer: new(bytes.Buffer),
	}
	ob.buffer.WriteString(existsSQL)
	ob.buffer.WriteString(" ORDER BY ")
	ob.buffer.WriteString(strings.Join(Map(columns, EscapeName), ","))
	return ob
}

func (slf *OrderBuilder) Column(column string) *OrderBuilder {
	slf.buffer.WriteString(", ")
	slf.buffer.WriteString(EscapeName(column))
	return slf
}

func (slf *OrderBuilder) ASC() *OrderBuilder {
	slf.buffer.WriteString(" ASC")
	return slf
}

func (slf *OrderBuilder) DESC() *OrderBuilder {
	slf.buffer.WriteString(" DESC")
	return slf
}

func (slf *OrderBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.SQL(), limit)
}

func (slf *OrderBuilder) SQL() string {
	return slf.buffer.String()
}

func (slf *OrderBuilder) MakeSQL() string {
	return makeSQL(slf.buffer)
}

func (slf *OrderBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
