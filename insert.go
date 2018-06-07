package sqlbuilder

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type InsertBuilder struct {
	ctx     *SQLContext
	table   string
	columns *List
	values  *List
	ignore  bool
}

func newInsertBuilder(ctx *SQLContext, table string) *InsertBuilder {
	return &InsertBuilder{
		ctx:     ctx,
		table:   table,
		columns: newItems(),
		values:  newItems(),
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
	slf.columns.Add(column)
	slf.values.Add(slf.ctx.placeHolder)
	// others
	for _, col := range otherColumns {
		slf.columns.Add(col)
		slf.values.Add(slf.ctx.placeHolder)
	}
	return slf
}

// Values 依次地顺序地替换Column列表对应的数值占位符。
func (slf *InsertBuilder) Values(value interface{}, otherValues ...interface{}) *InsertBuilder {
	if len(otherValues)+1 != slf.columns.Count() {
		panic("Count of values not match columns")
	}
	// one
	slf.values.SetAt(value, 0)
	for i, newVal := range otherValues {
		slf.values.SetAt(newVal, i+1)
	}
	return slf
}

// SetValueOfColumn 设置指定名称Column对应的值
func (slf *InsertBuilder) SetValueOfColumn(column string, newValue interface{}) *InsertBuilder {
	slf.columns.Range(func(index int, val interface{}) bool {
		col := val.(string)
		if column == col {
			slf.values.SetAt(newValue, index)
			return false
		}
		return true
	})

	return slf
}

////

func (slf *InsertBuilder) addColumnValue(column string, value interface{}) {
	slf.columns.Add(column)
	slf.values.Add(value)
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
	buf.WriteString(slf.ctx.escapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(slf.ctx.joinNames(slf.columns.AvailableStrItems()))
	buf.WriteByte(')')
	buf.WriteString(" VALUES ")
	buf.WriteByte('(')
	buf.WriteString(slf.ctx.joinValues(slf.values.AvailableItems()))
	buf.WriteByte(')')
	return buf
}

func (slf *InsertBuilder) ToSQL() string {
	return sqlEndpoint(slf.compile())
}

func (slf *InsertBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.db)
}
