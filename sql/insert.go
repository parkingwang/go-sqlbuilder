package sql

import (
	"bytes"
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

func (ib *InsertBuilder) Table(table string) *InsertBuilder {
	ib.table = table
	return ib
}

func (ib *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	for _, col := range columns {
		ib.columns = append(ib.columns, col)
		ib.values = append(ib.values, "?")
	}
	return ib
}

func (ib *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	// check columns and values
	if len(ib.columns) != len(ib.values) {
		panic("length of columns and values NOT MATCH")
	}
	for i, newVal := range values {
		ib.values[i] = newVal
	}
	return ib
}

func (ib *InsertBuilder) builder() *bytes.Buffer {
	if "" == ib.table {
		panic("table not found, you should call 'Table(table)' method to set it")
	}

	buf := new(bytes.Buffer)
	buf.WriteString("INSERT INTO ")
	buf.WriteString(EscapeName(ib.table))
	buf.WriteByte('(')
	buf.WriteString(strings.Join(Map(ib.columns, EscapeName), ","))
	buf.WriteByte(')')
	buf.WriteString(" VALUES ")
	buf.WriteByte('(')
	buf.WriteString(strings.Join(Map0(ib.values, EscapeValue), ","))
	buf.WriteByte(')')
	return buf
}

func (ib *InsertBuilder) SQL() string {
	return endpoint(ib.builder())
}
