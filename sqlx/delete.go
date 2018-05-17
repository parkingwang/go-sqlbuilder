package sqlx

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type DeleteBuilder struct {
	table       string
	forceDelete bool
}

func Delete(table string) *DeleteBuilder {
	return &DeleteBuilder{
		table: table,
	}
}

func (slf *DeleteBuilder) Table(table string) *DeleteBuilder {
	slf.table = table
	return slf
}

func (slf *DeleteBuilder) build() *bytes.Buffer {
	if "" == slf.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("DELETE FROM ")
	buf.WriteString(EscapeName(slf.table))
	return buf
}

func (slf *DeleteBuilder) YesYesYesForceDelete() *DeleteBuilder {
	slf.forceDelete = true
	return slf
}

func (slf *DeleteBuilder) Where(conditions SQLStatement) *WhereBuilder {
	return newWhere(slf, conditions)
}

func (slf *DeleteBuilder) Statement() string {
	return slf.build().String()
}

func (slf *DeleteBuilder) GetSQL() string {
	sqlTxt := makeSQL(slf.build())
	if slf.forceDelete {
		return sqlTxt
	} else {
		panic("Warning for FULL-DELETE, you should call 'YesYesYesForceDelete(bool)' to ensure. SQLText: " + sqlTxt)
	}
}

func (slf *DeleteBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}
