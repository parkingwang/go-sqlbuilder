package sqlx

import (
	"bytes"
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type InsertBuilder struct {
	table   string
	columns []string
	values  []interface{}
}

func Insert(table string) *InsertBuilder {
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
		slf.values = append(slf.values, SQLPlaceHolder)
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

func (slf *InsertBuilder) build() *bytes.Buffer {
	if "" == slf.table {
		panic("table not found, you should call 'Table(table)' method to set it")
	}

	buf := new(bytes.Buffer)
	buf.WriteString("INSERT INTO ")
	buf.WriteString(EscapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(joinNames(slf.columns))
	buf.WriteByte(')')
	buf.WriteString(" VALUES ")
	buf.WriteByte('(')
	buf.WriteString(joinValues(slf.values))
	buf.WriteByte(')')
	return buf
}

func (slf *InsertBuilder) GetSQL() string {
	return makeSQL(slf.build())
}

func (slf *InsertBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.GetSQL(), db)
}
