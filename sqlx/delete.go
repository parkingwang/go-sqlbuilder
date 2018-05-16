package sqlx

import (
	"bytes"
	"database/sql"
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

func (slf *DeleteBuilder) builder() *bytes.Buffer {
	if "" == slf.table {
		panic("Table name not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("DELETE FROM ")
	buf.WriteString(EscapeName(slf.table))
	return buf
}

func (slf *DeleteBuilder) Where() *WhereBuilder {
	return newWhere(slf.builder())
}

func (slf *DeleteBuilder) YesYesYesForceDelete() *DeleteBuilder {
	slf.forceDelete = true
	return slf
}

func (slf *DeleteBuilder) SQL() string {
	if slf.forceDelete {
		return endpoint(slf.builder())
	} else {
		panic("Warning for full delete, you should call 'YesYesForceDelete()' to ensure.")
	}
}

func (slf *DeleteBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
