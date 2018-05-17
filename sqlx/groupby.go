package sqlx

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type GroupByBuilder struct {
	buffer *bytes.Buffer
}

func newGroupBy(preStatement SQLStatement, columns ...string) *GroupByBuilder {
	gbb := &GroupByBuilder{
		buffer: new(bytes.Buffer),
	}
	gbb.buffer.WriteString(preStatement.Statement())
	gbb.buffer.WriteString(" GROUP BY ")
	gbb.buffer.WriteString(joinNames(columns))
	return gbb
}

func (slf *GroupByBuilder) Statement() string {
	return slf.buffer.String()
}

func (slf *GroupByBuilder) GetSQL() string {
	return endOfSQL(slf.buffer)
}

func (slf *GroupByBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}

func (slf *GroupByBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf, limit)
}

func (slf *GroupByBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderBy(slf, columns...)
}
