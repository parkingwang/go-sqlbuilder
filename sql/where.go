package sql

import "bytes"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type WhereBuilder struct {
	writer *bytes.Buffer
}

func newWhere(statement *bytes.Buffer) *WhereBuilder {
	return &WhereBuilder{
		writer: statement,
	}
}

func (wb *WhereBuilder) And() *WhereBuilder {
	wb.writer.WriteString(" AND ")
	return wb
}

func (wb *WhereBuilder) Or() *WhereBuilder {
	wb.writer.WriteString(" OR ")
	return wb
}

func (wb *WhereBuilder) Equal(column string) *WhereBuilder {
	return wb.op(column, "=")
}

func (wb *WhereBuilder) NotEqual(column string) *WhereBuilder {
	return wb.op(column, "<>")
}

func (wb *WhereBuilder) GreaterThen(column string) *WhereBuilder {
	return wb.op(column, ">")
}

func (wb *WhereBuilder) GreaterEqualThen(column string) *WhereBuilder {
	return wb.op(column, ">=")
}


func (wb *WhereBuilder) LessThen(column string) *WhereBuilder {
	return wb.op(column, "<=")
}

func (wb *WhereBuilder) LessEqualThen(column string) *WhereBuilder {
	return wb.op(column, "<=")
}


func (wb *WhereBuilder) op(column string, op string) *WhereBuilder {
	wb.writer.WriteString(Escape(column))
	wb.writer.WriteString(op)
	wb.writer.WriteByte('?')
	return wb
}

func (wb *WhereBuilder) SQL() string {
	wb.writer.WriteByte(';')
	return wb.writer.String()
}