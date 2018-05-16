package sql

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// Where条件连接
type CondLinker struct {
	whereBuilder *WhereBuilder
}

func (cb *CondLinker) And() *WhereBuilder {
	return cb.whereBuilder.and()
}

func (cb *CondLinker) Or() *WhereBuilder {
	return cb.whereBuilder.or()
}

func (cb *CondLinker) SQL() string {
	return cb.whereBuilder.SQL()
}


////

// Where条件构建器
type WhereBuilder struct {
	buffer *bytes.Buffer
	linker *CondLinker
}

func newWhere(statement *bytes.Buffer) *WhereBuilder {
	wb := &WhereBuilder{
		buffer: statement,
	}
	wb.linker = &CondLinker{
		whereBuilder: wb,
	}
	wb.buffer.WriteString(" WHERE ")
	return wb
}

func (wb *WhereBuilder) Equal(column string) *CondLinker {
	return wb.op(column, "=").linker
}

func (wb *WhereBuilder) EqualTo(column string, value interface{}) *CondLinker {
	return wb.opTo(column, "=", value).linker
}

func (wb *WhereBuilder) NotEqual(column string) *CondLinker {
	return wb.op(column, "<>").linker
}

func (wb *WhereBuilder) NotEqualTo(column string, value interface{}) *CondLinker {
	return wb.opTo(column, "<>", value).linker
}

func (wb *WhereBuilder) GreaterThen(column string) *CondLinker {
	return wb.op(column, ">").linker
}

func (wb *WhereBuilder) GreaterThenTo(column string, value interface{}) *CondLinker {
	return wb.opTo(column, ">", value).linker
}

func (wb *WhereBuilder) GreaterEqualThen(column string) *CondLinker {
	return wb.op(column, ">=").linker
}

func (wb *WhereBuilder) GreaterEqualThenTo(column string, value interface{}) *CondLinker {
	return wb.opTo(column, ">=", value).linker
}

func (wb *WhereBuilder) LessThen(column string) *CondLinker {
	return wb.op(column, "<").linker
}

func (wb *WhereBuilder) LessThenTo(column string, value interface{}) *CondLinker {
	return wb.opTo(column, "<", value).linker
}

func (wb *WhereBuilder) LessEqualThen(column string) *CondLinker {
	return wb.op(column, "<=").linker
}

func (wb *WhereBuilder) LessEqualTo(column string, value interface{}) *CondLinker {
	return wb.opTo(column, "<=", value).linker
}

//

func (wb *WhereBuilder) and() *WhereBuilder {
	wb.buffer.WriteString(" AND ")
	return wb
}

func (wb *WhereBuilder) or() *WhereBuilder {
	wb.buffer.WriteString(" OR ")
	return wb
}

func (wb *WhereBuilder) op(column string, op string) *WhereBuilder {
	return wb.opTo(column, op, "?")
}

func (wb *WhereBuilder) opTo(column string, op string, val interface{}) *WhereBuilder {
	wb.buffer.WriteString(EscapeColumn(column))
	wb.buffer.WriteString(op)
	wb.buffer.WriteString(EscapeValue(val))
	return wb
}

//

func (wb *WhereBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(wb.buffer, limit)
}

func (wb *WhereBuilder) SQL() string {
	return endpoint(wb.buffer)
}
