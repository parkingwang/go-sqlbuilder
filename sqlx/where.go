package sqlx

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type WhereBuilder struct {
	buffer     *bytes.Buffer
	conditions SQLStatement
}

func (slf *WhereBuilder) Statement() string {
	slf.buffer.WriteString(slf.conditions.Statement())
	return slf.buffer.String()
}

func (slf *WhereBuilder) GetSQL() string {
	slf.buffer.WriteString(slf.conditions.Statement())
	return endOfSQL(slf.buffer)
}

func (slf *WhereBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf, limit)
}

func (slf *WhereBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderBy(slf, columns...)
}

func (slf *WhereBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf, columns...)
}

func (slf *WhereBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}

//

func newWhere(preStatement SQLStatement, conditions SQLStatement) *WhereBuilder {
	wb := &WhereBuilder{
		buffer:     new(bytes.Buffer),
		conditions: conditions,
	}
	if nil != preStatement {
		wb.buffer.WriteString(preStatement.Statement())
	}
	wb.buffer.WriteString(" WHERE ")
	return wb
}
