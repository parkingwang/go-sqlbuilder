package sqlx

import (
	"bytes"
	"database/sql"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type GroupByBuilder struct {
	buffer *bytes.Buffer
}

func newGroupBy(existsSQL string, columns ...string) *GroupByBuilder {
	lb := &GroupByBuilder{
		buffer: new(bytes.Buffer),
	}
	lb.buffer.WriteString(existsSQL)
	lb.buffer.WriteString(" GROUP BY ")
	lb.buffer.WriteString(strings.Join(Map(columns, EscapeName), ", "))
	return lb
}

func (slf *GroupByBuilder) SQL() string {
	return slf.buffer.String()
}

func (slf *GroupByBuilder) MakeSQL() string {
	return makeSQL(slf.buffer)
}

func (slf *GroupByBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}

func (slf *GroupByBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.SQL(), limit)
}

func (slf *GroupByBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.SQL(), columns...)
}
