package sqlx

import (
	"bytes"
	"strconv"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type LimitBuilder struct {
	buffer *bytes.Buffer
}

func newLimit(preStatement SQLStatement, limit int) *LimitBuilder {
	lb := &LimitBuilder{
		buffer: new(bytes.Buffer),
	}
	lb.buffer.WriteString(preStatement.Statement())
	lb.buffer.WriteString(" LIMIT ")
	lb.buffer.WriteString(strconv.Itoa(limit))
	return lb
}

func (slf *LimitBuilder) Offset(offset int) *LimitBuilder {
	slf.buffer.WriteString(" OFFSET ")
	slf.buffer.WriteString(strconv.Itoa(offset))
	return slf
}

func (slf *LimitBuilder) Statement() string {
	return slf.buffer.String()
}

func (slf *LimitBuilder) GetSQL() string {
	return endOfSQL(slf.buffer)
}

func (slf *LimitBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}

func (slf *LimitBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderBy(slf, columns...)
}

func (slf *LimitBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf, columns...)
}
