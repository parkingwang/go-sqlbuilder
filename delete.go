package gsb

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type DeleteBuilder struct {
	ctx    *SQLBuilder
	table  string
	ensure bool // 删除全表时，需要强制设置一个标记位。
}

func newDeleteBuilder(ctx *SQLBuilder, table string) *DeleteBuilder {
	return &DeleteBuilder{
		ctx:    ctx,
		table:  table,
		ensure: false,
	}
}

func (slf *DeleteBuilder) Table(table string) *DeleteBuilder {
	slf.table = table
	return slf
}

func (slf *DeleteBuilder) compile() *bytes.Buffer {
	if "" == slf.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("DELETE FROM ")
	buf.WriteString(slf.ctx.EscapeName(slf.table))
	return buf
}

func (slf *DeleteBuilder) YesImSureDeleteTable() *DeleteBuilder {
	slf.ensure = true
	return slf
}

func (slf *DeleteBuilder) Where(conditions SQLStatement) *WhereBuilder {
	return newWhere(slf, conditions)
}

func (slf *DeleteBuilder) Compile() string {
	return slf.compile().String()
}

func (slf *DeleteBuilder) ToSQL() string {
	sqlTxt := endOfSQL(slf.compile())
	if slf.ensure {
		return sqlTxt
	} else {
		panic("Warning for FULL-DELETE the table, you must call 'YesImSureDeleteTable(bool)' to ensure. SQLText: " + sqlTxt)
	}
}

func (slf *DeleteBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.ToSQL(), prepare)
}
