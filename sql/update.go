package sql

import (
	"bytes"
	"database/sql"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type UpdateBuilder struct {
	table       string
	columns     []string
	forceUpdate bool
}

func Update(table string) *UpdateBuilder {
	return (&UpdateBuilder{
		columns: make([]string, 0),
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
	slf.columns = Map(columns, func(column string) string {
		return EscapeName(column) + "=?"
	})
	return slf
}

func (slf *UpdateBuilder) ColumnAndValue(column string, value interface{}) *UpdateBuilder {
	slf.columns = append(slf.columns, func() string {
		return EscapeName(column) + "=" + EscapeValue(value)
	}())
	return slf
}

func (slf *UpdateBuilder) builder() *bytes.Buffer {
	if "" == slf.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("UPDATE ")
	buf.WriteString(EscapeName(slf.table))
	buf.WriteString(" SET ")
	buf.WriteString(strings.Join(slf.columns, ","))
	return buf
}

func (slf *UpdateBuilder) Where() *WhereBuilder {
	return newWhere(slf.builder())
}

func (slf *UpdateBuilder) YesYesYesForceUpdate() *UpdateBuilder {
	slf.forceUpdate = true
	return slf
}

func (slf *UpdateBuilder) SQL() string {
	if slf.forceUpdate {
		return endpoint(slf.builder())
	} else {
		panic("Warning for full update, you should call 'YesYesYesForceUpdate()' to ensure.")
	}
}

func (slf *UpdateBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
