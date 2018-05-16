package sqlx

import (
	"bytes"
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type WhereBuilder struct {
	buffer     *bytes.Buffer
	conditions Statement
}

func (slf *WhereBuilder) SQL() string {
	slf.buffer.WriteString(slf.conditions.SQL())
	return slf.buffer.String()
}

func (slf *WhereBuilder) MakeSQL() string {
	slf.buffer.WriteString(slf.conditions.SQL())
	return makeSQL(slf.buffer)
}

func (slf *WhereBuilder) Limit(limit int) *LimitBuilder {
	return newLimit(slf.SQL(), limit)
}

func (slf *WhereBuilder) OrderBy(columns ...string) *OrderByBuilder {
	return newOrderByBuilder(slf.SQL(), columns...)
}

func (slf *WhereBuilder) GroupBy(columns ...string) *GroupByBuilder {
	return newGroupBy(slf.SQL(), columns...)
}

func (slf *WhereBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.MakeSQL(), db)
}

//

func newWhereWith(existsSQL string, conditions Statement) *WhereBuilder {
	wb := &WhereBuilder{
		buffer:     new(bytes.Buffer),
		conditions: conditions,
	}
	if len(existsSQL) > 0 {
		wb.buffer.WriteString(existsSQL)
	}
	wb.buffer.WriteString(" WHERE ")
	return wb
}

func newWhere(conditions Statement) *WhereBuilder {
	return newWhereWith("", conditions)
}
