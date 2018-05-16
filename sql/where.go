package sql

import (
	"bytes"
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// Where条件连接
type CondLinker struct {
	whereBuilder *WhereBuilder
}

func (slf *CondLinker) And() *WhereBuilder {
	return slf.whereBuilder.and()
}

func (slf *CondLinker) Or() *WhereBuilder {
	return slf.whereBuilder.or()
}

func (slf *CondLinker) Limit(limit int) *LimitBuilder {
	return newLimit(slf.whereBuilder.buffer, limit)
}

func (slf *CondLinker) OrderBy(columns ...string) *OrderBuilder {
	return newOrderBuilder(slf.whereBuilder.buffer, columns...)
}

func (slf *CondLinker) SQL() string {
	return slf.whereBuilder.SQL()
}

func (slf *CondLinker) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
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

func (slf *WhereBuilder) Equal(column string) *CondLinker {
	return slf.op(column, " = ").linker
}

func (slf *WhereBuilder) EqualTo(column string, value interface{}) *CondLinker {
	return slf.opTo(column, " = ", value).linker
}

func (slf *WhereBuilder) NotEqual(column string) *CondLinker {
	return slf.op(column, " <> ").linker
}

func (slf *WhereBuilder) NotEqualTo(column string, value interface{}) *CondLinker {
	return slf.opTo(column, " <> ", value).linker
}

func (slf *WhereBuilder) GreaterThen(column string) *CondLinker {
	return slf.op(column, " > ").linker
}

func (slf *WhereBuilder) GreaterThenTo(column string, value interface{}) *CondLinker {
	return slf.opTo(column, " > ", value).linker
}

func (slf *WhereBuilder) GreaterEqualThen(column string) *CondLinker {
	return slf.op(column, " >= ").linker
}

func (slf *WhereBuilder) GreaterEqualThenTo(column string, value interface{}) *CondLinker {
	return slf.opTo(column, " >= ", value).linker
}

func (slf *WhereBuilder) LessThen(column string) *CondLinker {
	return slf.op(column, " < ").linker
}

func (slf *WhereBuilder) LessThenTo(column string, value interface{}) *CondLinker {
	return slf.opTo(column, " < ", value).linker
}

func (slf *WhereBuilder) LessEqualThen(column string) *CondLinker {
	return slf.op(column, " <= ").linker
}

func (slf *WhereBuilder) LessEqualTo(column string, value interface{}) *CondLinker {
	return slf.opTo(column, " <= ", value).linker
}

//

func (slf *WhereBuilder) and() *WhereBuilder {
	slf.buffer.WriteString(" AND ")
	return slf
}

func (slf *WhereBuilder) or() *WhereBuilder {
	slf.buffer.WriteString(" OR ")
	return slf
}

func (slf *WhereBuilder) op(column string, op string) *WhereBuilder {
	return slf.opTo(column, op, "?")
}

func (slf *WhereBuilder) opTo(column string, op string, val interface{}) *WhereBuilder {
	slf.buffer.WriteString(EscapeName(column))
	slf.buffer.WriteString(op)
	slf.buffer.WriteString(EscapeValue(val))
	return slf
}

//

func (slf *WhereBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.buffer, limit)
}

func (slf *WhereBuilder) SQL() string {
	return endpoint(slf.buffer)
}

func (slf *WhereBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
