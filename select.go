package sqlbuilder

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type SelectBuilder struct {
	ctx        *SQLContext
	columns    []string
	table      string
	fromSelect SQLStatement
	distinct   bool
}

func newSelectBuilder(ctx *SQLContext, column string, otherColumns ...string) *SelectBuilder {
	columns := make([]string, len(otherColumns)+1)
	columns[0] = column
	for i, col := range otherColumns {
		columns[i+1] = col
	}
	return &SelectBuilder{
		ctx:      ctx,
		columns:  columns,
		distinct: false,
	}
}

func (slf *SelectBuilder) Distinct() *SelectBuilder {
	slf.distinct = true
	return slf
}

func (slf *SelectBuilder) From(table string) *SelectBuilder {
	slf.table = table
	return slf
}

func (slf *SelectBuilder) FromSelect(innerSelect SQLStatement) *SelectBuilder {
	slf.fromSelect = innerSelect
	return slf
}

func (slf *SelectBuilder) compile() *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")
	if slf.distinct {
		buf.WriteString("DISTINCT ")
	}
	buf.WriteString(slf.ctx.joinNames(slf.columns))
	buf.WriteString(" FROM ")
	if nil != slf.fromSelect {
		buf.WriteByte('(')
		buf.WriteString(slf.fromSelect.Compile())
		buf.WriteByte(')')
	} else {
		buf.WriteString(slf.ctx.escapeName(slf.table))
	}
	return buf
}

func (slf *SelectBuilder) Where(conditions SQLStatement) *WhereBuilder {
	return newWhereBuilder(slf.ctx, slf.Compile(), conditions)
}

func (slf *SelectBuilder) Limit(limit int) *LimitBuilder {
	return newLimitBuilder(slf.ctx, slf.Compile(), limit)
}

func (slf *SelectBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *SelectBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupByBuilder(slf.ctx, slf.Compile(), columns...)
}

func (slf *SelectBuilder) Compile() string {
	return slf.compile().String()
}

func (slf *SelectBuilder) ToSQL() string {
	return sqlEndpoint(slf.compile())
}

func (slf *SelectBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.db)
}
