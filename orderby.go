package sqlbuilder

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type OrderByBuilder struct {
	ctx    *SQLContext
	buffer *bytes.Buffer
}

func newOrderByBuilder(ctx *SQLContext, pre string, columns ...string) *OrderByBuilder {
	if len(columns) == 0 {
		panic("Columns is required for ORDER BY keyword")
	}
	ob := &OrderByBuilder{
		ctx:    ctx,
		buffer: new(bytes.Buffer),
	}
	ob.buffer.WriteString(pre)
	ob.buffer.WriteString(" ORDER BY ")
	ob.buffer.WriteString(strings.Join(MapStr(columns, ctx.escapeName), ","))
	return ob
}

func (slf *OrderByBuilder) Column(column string) *OrderByBuilder {
	slf.buffer.WriteString(", ")
	slf.buffer.WriteString(slf.ctx.escapeName(column))
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
	return newLimitBuilder(slf.ctx, slf.Compile(), limit)
}

func (slf *OrderByBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *OrderByBuilder) Compile() string {
	return slf.buffer.String()
}

func (slf *OrderByBuilder) ToSQL() string {
	return sqlEndpoint(slf.buffer)
}

func (slf *OrderByBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.prepare)
}
