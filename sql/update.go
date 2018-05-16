package sql

import (
	"bytes"
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

func (ub *UpdateBuilder) Table(table string) *UpdateBuilder {
	ub.table = table
	return ub
}

func (ub *UpdateBuilder) Columns(columns ...string) *UpdateBuilder {
	if len(columns) == 0 {
		panic("Columns is required for update")
	}
	ub.columns = Map(columns, func(column string) string {
		return EscapeName(column) + "=?"
	})
	return ub
}

func (ub *UpdateBuilder) ColumnAndValue(column string, value interface{}) *UpdateBuilder {
	ub.columns = append(ub.columns, func() string {
		return EscapeName(column) + "=" + EscapeValue(value)
	}())
	return ub
}

func (ub *UpdateBuilder) builder() *bytes.Buffer {
	if "" == ub.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("UPDATE ")
	buf.WriteString(EscapeName(ub.table))
	buf.WriteString(" SET ")
	buf.WriteString(strings.Join(ub.columns, ","))
	return buf
}

func (ub *UpdateBuilder) Where() *WhereBuilder {
	return newWhere(ub.builder())
}

func (ub *UpdateBuilder) YesYesYesForceUpdate() *UpdateBuilder {
	ub.forceUpdate = true
	return ub
}

func (ub *UpdateBuilder) SQL() string {
	if ub.forceUpdate {
		return endpoint(ub.builder())
	} else {
		panic("Warning for full update, you should call 'YesYesYesForceUpdate()' to ensure.")
	}
}
