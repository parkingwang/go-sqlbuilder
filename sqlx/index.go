package sqlx

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type IndexBuilder struct {
	table   string
	name    string
	columns []string
	unique  bool
}

func CreateIndex(indexName string) *IndexBuilder {
	return &IndexBuilder{
		name:    indexName,
		columns: make([]string, 0),
	}
}

func (slf *IndexBuilder) Unique() *IndexBuilder {
	slf.unique = true
	return slf
}

func (slf *IndexBuilder) OnTable(table string) *IndexBuilder {
	slf.table = table
	return slf
}

func (slf *IndexBuilder) Column(name string, desc bool) *IndexBuilder {
	var column string
	if desc {
		column = EscapeName(name) + SQLSpace + "DESC"
	} else {
		column = EscapeName(column)
	}
	slf.columns = append(slf.columns, column)
	return slf
}

func (slf *IndexBuilder) Columns(columns ...string) *IndexBuilder {
	slf.columns = append(slf.columns, Map(columns, EscapeName)...)
	return slf
}

func (slf *IndexBuilder) build() *bytes.Buffer {
	if "" == slf.table {
		panic("table not found, you should call 'Table(table)' method to set it")
	}

	buf := new(bytes.Buffer)
	buf.WriteString("CREATE ")
	if slf.unique {
		buf.WriteString("UNIQUE ")
	}
	buf.WriteString("INDEX ")
	buf.WriteString(EscapeName(slf.name))
	buf.WriteString(" ON ")
	buf.WriteString(EscapeName(slf.table))
	buf.WriteByte('(')
	// 在输入时已经转义
	buf.WriteString(strings.Join(slf.columns, SQLComma))
	buf.WriteByte(')')
	return buf
}

func (slf *IndexBuilder) GetSQL() string {
	return makeSQL(slf.build())
}

func (slf *IndexBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}
