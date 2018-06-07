package sqlbuilder

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type GroupByBuilder struct {
	ctx    *SQLContext
	buffer *bytes.Buffer
}

func newGroupByBuilder(ctx *SQLContext, pre string, columns ...string) *GroupByBuilder {
	gbb := &GroupByBuilder{
		ctx:    ctx,
		buffer: new(bytes.Buffer),
	}
	gbb.buffer.WriteString(pre)
	gbb.buffer.WriteString(" GROUP BY ")
	gbb.buffer.WriteString(ctx.joinNames(columns))
	return gbb
}

func (slf *GroupByBuilder) Compile() string {
	return slf.buffer.String()
}

func (slf *GroupByBuilder) ToSQL() string {
	return sqlEndpoint(slf.buffer)
}

func (slf *GroupByBuilder) Limit(limit int) *LimitBuilder {
	return newLimitBuilder(slf.ctx, slf.Compile(), limit)
}

func (slf *GroupByBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *GroupByBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.db)
}
