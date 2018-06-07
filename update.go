package gsb

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type UpdateBuilder struct {
	ctx             *SQLContext
	table           string
	columnAndValues *List
	ensure          bool
}

func newUpdateBuilder(ctx *SQLContext, table string) *UpdateBuilder {
	return &UpdateBuilder{
		ctx:             ctx,
		table:           table,
		columnAndValues: newItems(),
	}
}

func (slf *UpdateBuilder) Table(table string) *UpdateBuilder {
	slf.table = table
	return slf
}

func (slf *UpdateBuilder) Columns(column string, otherColumns ...string) *UpdateBuilder {
	slf.columnAndValues.Add(escapeWith(slf.ctx, column))
	for _, col := range otherColumns {
		slf.columnAndValues.Add(escapeWith(slf.ctx, col))
	}
	return slf
}

func (slf *UpdateBuilder) AddColumnValue(column string, value interface{}) *UpdateBuilder {
	slf.columnAndValues.Add(slf.ctx.escapeName(column) + "=" + slf.ctx.escapeValue(value))
	return slf
}

func (slf *UpdateBuilder) compile() *bytes.Buffer {
	if "" == slf.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("UPDATE ")
	buf.WriteString(slf.ctx.escapeName(slf.table))
	buf.WriteString(" SET ")
	// 此处Columns不需要转义处理
	buf.WriteString(strings.Join(slf.columnAndValues.AvailableStrItems(), SQLComma))
	return buf
}

func (slf *UpdateBuilder) YesImSureUpdateTable() *UpdateBuilder {
	slf.ensure = true
	return slf
}

func (slf *UpdateBuilder) Where(conditions SQLStatement) *WhereBuilder {
	return newWhereBuilder(slf.ctx, slf.Compile(), conditions)
}

func (slf *UpdateBuilder) Compile() string {
	return slf.compile().String()
}

func (slf *UpdateBuilder) ToSQL() string {
	sqlTxt := sqlEndpoint(slf.compile())
	if slf.ensure {
		return sqlTxt
	} else {
		panic("Warning for FULL-UPDATE, you should call 'YesImSureUpdateTable(bool)' to ensure. SQLText: " + sqlTxt)
	}
}

func (slf *UpdateBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.db)
}

func escapeWith(ctx *SQLContext, column string) string {
	return ctx.escapeName(column) + "=?"
}
