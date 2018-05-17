package sqlx

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type CreateIndexBuilder struct {
	table   string
	name    string
	columns []string
	unique  bool
}

func CreateIndex(indexName string) *CreateIndexBuilder {
	return &CreateIndexBuilder{
		name:    indexName,
		columns: make([]string, 0),
	}
}

func (slf *CreateIndexBuilder) Unique() *CreateIndexBuilder {
	slf.unique = true
	return slf
}

func (slf *CreateIndexBuilder) OnTable(table string) *CreateIndexBuilder {
	slf.table = table
	return slf
}

func (slf *CreateIndexBuilder) Column(name string, desc bool) *CreateIndexBuilder {
	var column string
	if desc {
		column = EscapeName(name) + SQLSpace + "DESC"
	} else {
		column = EscapeName(column)
	}
	slf.columns = append(slf.columns, column)
	return slf
}

func (slf *CreateIndexBuilder) Columns(columns ...string) *CreateIndexBuilder {
	slf.columns = append(slf.columns, Map(columns, EscapeName)...)
	return slf
}

func (slf *CreateIndexBuilder) build() *bytes.Buffer {
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

func (slf *CreateIndexBuilder) GetSQL() string {
	return endOfSQL(slf.build())
}

func (slf *CreateIndexBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}
