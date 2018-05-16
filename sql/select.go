package sql

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type SelectBuilder struct {
	columns  []string
	table    string
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

func (sb *SelectBuilder) buffer() *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")

	if sb.distinct {
		buf.WriteString("DISTINCT ")
	}

	if len(sb.columns) == 0 {
		buf.WriteByte('*')
	} else {
		buf.WriteString(strings.Join(Map(sb.columns, EscapeName), ","))
	}

	buf.WriteString(" FROM ")
	buf.WriteString(EscapeName(sb.table))
	return buf
}

func (sb *SelectBuilder) OrderBy(columns ...string) *OrderBuilder {
	return newOrderBuilder(sb.buffer(), columns...)
}

func (sb *SelectBuilder) Where() *WhereBuilder {
	return newWhere(sb.buffer())
}

func (sb *SelectBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(sb.buffer(), limit)
}

func (sb *SelectBuilder) SQL() string {
	return endpoint(sb.buffer())
}
