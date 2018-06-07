package sqlbuilder

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type WhereBuilder struct {
	ctx        *SQLContext
	buffer     *bytes.Buffer
	conditions SQLStatement
}

func (slf *WhereBuilder) Compile() string {
	slf.buffer.WriteString(slf.conditions.Compile())
	return slf.buffer.String()
}

func (slf *WhereBuilder) ToSQL() string {
	slf.buffer.WriteString(slf.conditions.Compile())
	return sqlEndpoint(slf.buffer)
}

//
func (slf *WhereBuilder) Limit(limit int) *LimitBuilder {
	return newLimitBuilder(slf.ctx, slf.Compile(), limit)
}

func (slf *WhereBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *WhereBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *WhereBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.db)
}

func newWhereBuilder(ctx *SQLContext, pre string, conditions SQLStatement) *WhereBuilder {
	wb := &WhereBuilder{
		ctx:        ctx,
		buffer:     new(bytes.Buffer),
		conditions: conditions,
	}
	wb.buffer.WriteString(pre)
	wb.buffer.WriteString(" WHERE ")
	return wb
}
