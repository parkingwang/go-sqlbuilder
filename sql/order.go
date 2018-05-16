package sql

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

////

type OrderBuilder struct {
	buffer *bytes.Buffer
}

func newOrderBuilder(writer *bytes.Buffer, columns...string) *OrderBuilder {
	if len(columns) < 1 {
		panic("columns is required for ORDER BY keyword")
	}
	ob := &OrderBuilder{
		buffer: writer,
	}
	ob.buffer.WriteString(" ORDER BY ")
	ob.buffer.WriteString(strings.Join(Map(columns, EscapeColumn), ","))
	return ob
}

func (ob *OrderBuilder) Column(column string) *OrderBuilder {
	ob.buffer.WriteString(", ")
	ob.buffer.WriteString(EscapeColumn(column))
	return ob
}

func (ob *OrderBuilder) ASC() *OrderBuilder {
	ob.buffer.WriteString(" ASC")
	return ob
}

func (ob *OrderBuilder) DESC() *OrderBuilder {
	ob.buffer.WriteString(" DESC")
	return ob
}

func (ob *OrderBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(ob.buffer, limit)
}

func (ob *OrderBuilder) SQL() string {
	return endpoint(ob.buffer)
}