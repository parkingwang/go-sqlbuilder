package sql

import (
	"bytes"
	"database/sql"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type InsertBuilder struct {
	table   string
	columns []string
	values  []interface{}
}

func InsertInto(table string) *InsertBuilder {
	return &InsertBuilder{
		table:   table,
		columns: make([]string, 0),
		values:  make([]interface{}, 0),
	}
}

func (slf *InsertBuilder) Table(table string) *InsertBuilder {
	slf.table = table
	return slf
}

func (slf *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	for _, col := range columns {
		slf.columns = append(slf.columns, col)
		slf.values = append(slf.values, "?")
	}
	return slf
}

func (slf *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	// check columns and values
	if len(slf.columns) != len(slf.values) {
		panic("length of columns and values NOT MATCH")
	}
	for i, newVal := range values {
		slf.values[i] = newVal
	}
	return slf
}

func (slf *InsertBuilder) builder() *bytes.Buffer {
	if "" == slf.table {
		panic("table not found, you should call 'Table(table)' method to set it")
	}

	buf := new(bytes.Buffer)
	buf.WriteString("INSERT INTO ")
	buf.WriteString(EscapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(strings.Join(Map(slf.columns, EscapeName), ","))
	buf.WriteByte(')')
	buf.WriteString(" VALUES ")
	buf.WriteByte('(')
	buf.WriteString(strings.Join(Map0(slf.values, EscapeValue), ","))
	buf.WriteByte(')')
	return buf
}

func (slf *InsertBuilder) SQL() string {
	return endpoint(slf.builder())
}

func (slf *InsertBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
