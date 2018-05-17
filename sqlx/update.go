package sqlx

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type UpdateBuilder struct {
	table        string
	columnValues []string
	forceUpdate  bool
}

func Update(table string) *UpdateBuilder {
	return (&UpdateBuilder{
		columnValues: make([]string, 0),
	}).Table(table)
}

func (slf *UpdateBuilder) Table(table string) *UpdateBuilder {
	slf.table = table
	return slf
}

func (slf *UpdateBuilder) Columns(columns ...string) *UpdateBuilder {
	if len(columns) == 0 {
		panic("Columns is required for update")
	}
	slf.columnValues = Map(columns, func(column string) string {
		return EscapeName(column) + "=?"
	})
	return slf
}

func (slf *UpdateBuilder) ColumnAndValue(column string, value interface{}) *UpdateBuilder {
	slf.columnValues = append(slf.columnValues, func() string {
		return EscapeName(column) + "=" + EscapeValue(value)
	}())
	return slf
}

func (slf *UpdateBuilder) compile() *bytes.Buffer {
	if "" == slf.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("UPDATE ")
	buf.WriteString(EscapeName(slf.table))
	buf.WriteString(" SET ")
	// 此处Columns不需要转义处理
	buf.WriteString(strings.Join(slf.columnValues, SQLComma))
	return buf
}

func (slf *UpdateBuilder) YesYesYesForceUpdate() *UpdateBuilder {
	slf.forceUpdate = true
	return slf
}

func (slf *UpdateBuilder) Where(conditions SQLStatement) *WhereBuilder {
	return newWhere(slf, conditions)
}

func (slf *UpdateBuilder) Compile() string {
	return slf.compile().String()
}

func (slf *UpdateBuilder) GetSQL() string {
	sqlTxt := endOfSQL(slf.compile())
	if slf.forceUpdate {
		return sqlTxt
	} else {
		panic("Warning for FULL-UPDATE, you should call 'YesYesYesForceUpdate(bool)' to ensure. SQLText: " + sqlTxt)
	}
}

func (slf *UpdateBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}
