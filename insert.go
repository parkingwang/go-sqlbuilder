package gsb

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type InsertBuilder struct {
	ctx     *SQLBuilder
	table   string
	columns []string
	values  []interface{}
	cursor  uint8
	ignore  bool
}

func newInsertBuilder(ctx *SQLBuilder, table string) *InsertBuilder {
	return &InsertBuilder{
		ctx:     ctx,
		table:   table,
		cursor:  0,
		columns: make([]string, SQLDefaultColumns, SQLMaxColumns),
		values:  make([]interface{}, SQLDefaultColumns, SQLMaxColumns),
		ignore:  false,
	}
}

// Ignore 设置SQL语法的Ignore选项。忽略SQL运行时错误，慎用。
func (slf *InsertBuilder) Ignore() *InsertBuilder {
	slf.ignore = true
	return slf
}

// Columns 配置Column列表。默认情况下，为每个Column创建一个数值占位符。
func (slf *InsertBuilder) Columns(column string, otherColumns ...string) *InsertBuilder {
	// one
	slf.setColumnValue(column, slf.ctx.placeHolder, 0)
	slf.cursor = 0
	// others
	for i, col := range otherColumns {
		slf.setColumnValue(col, slf.ctx.placeHolder, i+1)
		slf.cursor++
	}
	return slf
}

// Values 依次地顺序地替换Column列表对应的数值占位符。
func (slf *InsertBuilder) Values(value interface{}, otherValues ...interface{}) *InsertBuilder {
	if uint8(len(otherValues) /*+1 index: -1*/) != slf.cursor {
		panic("Count of values not match columns")
	}
	// one
	slf.values[0] = value
	for i, newVal := range otherValues {
		slf.values[i+1] = newVal
	}
	return slf
}

// SetValueOfColumn 设置指定名称Column对应的值
func (slf *InsertBuilder) SetValueOfColumn(column string, value interface{}) *InsertBuilder {
	for i, col := range slf.columns {
		if uint8(i) > slf.cursor {
			return slf
		}
		if column == col {
			slf.setColumnValue(column, value, i)
		}
	}

	return slf
}

////

func (slf *InsertBuilder) setColumnValue(column string, value interface{}, idx int) {
	slf.columns[idx] = column
	slf.values[idx] = value
}

func (slf *InsertBuilder) compile() *bytes.Buffer {
	if "" == slf.table {
		panic("table not found, you should call 'Table(table)' method to set it")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("INSERT")
	if slf.ignore {
		buf.WriteString(" IGNORE ")
	}
	buf.WriteString(" INTO ")
	buf.WriteString(slf.ctx.EscapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(slf.ctx.JoinNames(slf.columns[:slf.cursor+1]))
	buf.WriteByte(')')
	buf.WriteString(" VALUES ")
	buf.WriteByte('(')
	buf.WriteString(slf.ctx.JoinValues(slf.values[:slf.cursor+1]))
	buf.WriteByte(')')
	return buf
}

func (slf *InsertBuilder) ToSQL() string {
	return endOfSQL(slf.compile())
}

func (slf *InsertBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.ToSQL(), prepare)
}
