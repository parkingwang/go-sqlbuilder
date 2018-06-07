package sqlbuilder

import (
	"bytes"
	"strconv"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type LimitBuilder struct {
	ctx    *SQLContext
	buffer *bytes.Buffer
}

func newLimitBuilder(ctx *SQLContext, pre string, limit int) *LimitBuilder {
	lb := &LimitBuilder{
		ctx:    ctx,
		buffer: new(bytes.Buffer),
	}
	lb.buffer.WriteString(pre)
	lb.buffer.WriteString(" LIMIT ")
	lb.buffer.WriteString(strconv.Itoa(limit))
	return lb
}

func (slf *LimitBuilder) Offset(offset int) *LimitBuilder {
	slf.buffer.WriteString(" OFFSET ")
	slf.buffer.WriteString(strconv.Itoa(offset))
	return slf
}

func (slf *LimitBuilder) Compile() string {
	return slf.buffer.String()
}

func (slf *LimitBuilder) ToSQL() string {
	return sqlEndpoint(slf.buffer)
}

func (slf *LimitBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *LimitBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *LimitBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.prepare)
}
