package sql

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

func (ub *DeleteBuilder) Table(table string) *DeleteBuilder {
	ub.table = table
	return ub
}

func (ub *DeleteBuilder) builder() *bytes.Buffer {
	if "" == ub.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("DELETE FROM ")
	buf.WriteString(EscapeName(ub.table))
	return buf
}

func (ub *DeleteBuilder) Where() *WhereBuilder {
	return newWhere(ub.builder())
}

func (ub *DeleteBuilder) YesYesYesForceDelete() *DeleteBuilder {
	ub.forceDelete = true
	return ub
}

func (ub *DeleteBuilder) SQL() string {
	if ub.forceDelete {
		return endpoint(ub.builder())
	} else {
		panic("Warning for full delete, you should call 'YesYesForceDelete()' to ensure.")
	}
}
