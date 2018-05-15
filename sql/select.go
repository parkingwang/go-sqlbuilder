package sql

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type SelectBuilder struct {
	columns []string
	table 	string
	distinct bool
}

func Select(columns ...string) *SelectBuilder {
	return &SelectBuilder{
		columns: columns,
	}
}

func (sb *SelectBuilder) Distinct() *SelectBuilder {
	sb.distinct = true
	return sb
}

func (sb *SelectBuilder) From(table string) *SelectBuilder {
	sb.table = table
	return sb
}

func (sb *SelectBuilder) build() *bytes.Buffer {
	if "" == sb.table {
		panic("table not found, you should call 'From(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")

	if sb.distinct {
		buf.WriteString("DISTINCT ")
	}

	if len(sb.columns) == 0 {
		buf.WriteByte('*')
	}else{
		buf.WriteString(strings.Join(Map(sb.columns, Escape), ","))
	}

	buf.WriteString(" FROM ")
	buf.WriteString(Escape(sb.table))
	return buf
}

func (sb *SelectBuilder) Where() *WhereBuilder {
	return newWhere(sb.build())
}

func (sb *SelectBuilder) SQL() string {
	buf := sb.build()
	buf.WriteByte(';')
	return buf.String()
}


