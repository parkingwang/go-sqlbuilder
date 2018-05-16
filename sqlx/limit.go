package sqlx

import (
	"bytes"
	"database/sql"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type LimitBuilder struct {
	buffer *bytes.Buffer
}

func newLimit(existsSQL string, limit int) *LimitBuilder {
	lb := &LimitBuilder{
		buffer: new(bytes.Buffer),
	}
	lb.buffer.WriteString(existsSQL)
	lb.buffer.WriteString(" LIMIT ")
	lb.buffer.WriteString(fmt.Sprintf("%d", limit))
	return lb
}

func (slf *LimitBuilder) Offset(offset int) *LimitBuilder {
	slf.buffer.WriteString(" OFFSET ")
	slf.buffer.WriteString(fmt.Sprintf("%d", offset))
	return slf
}

func (slf *LimitBuilder) SQL() string {
	return slf.buffer.String()
}

func (slf *LimitBuilder) MakeSQL() string {
	return makeSQL(slf.buffer)
}

func (slf *LimitBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.MakeSQL(), db)
}

func (slf *LimitBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.SQL(), columns...)
}

func (slf *LimitBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf.SQL(), columns...)
}
