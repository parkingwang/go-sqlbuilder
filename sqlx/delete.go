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

func (slf *DeleteBuilder) Where(conditions Statement) *WhereBuilder {
	return newWhereWith(slf.SQL(), conditions)
}

func (slf *DeleteBuilder) SQL() string {
	return slf.build().String()
}

func (slf *DeleteBuilder) MakeSQL() string {
	sqlTxt := makeSQL(slf.build())
	if slf.forceDelete {
		return sqlTxt
	} else {
		panic("Warning for FULL-DELETE, you should call 'YesYesYesForceDelete()' to ensure. SQL: " + sqlTxt)
	}
}

func (slf *DeleteBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.MakeSQL(), db)
}
