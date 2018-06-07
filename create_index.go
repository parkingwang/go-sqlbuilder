package gsb

import (
	"bytes"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type CreateIndexBuilder struct {
	ctx     *SQLContext
	table   string
	name    string
	columns []string
	unique  bool
}

func newCreateIndex(ctx *SQLContext, indexName string) *CreateIndexBuilder {
	return &CreateIndexBuilder{
		ctx:     ctx,
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
		column = slf.ctx.escapeName(name) + SQLSpace + "DESC"
	} else {
		column = slf.ctx.escapeName(column)
	}
	slf.columns = append(slf.columns, column)
	return slf
}

func (slf *CreateIndexBuilder) Columns(columns ...string) *CreateIndexBuilder {
	slf.columns = append(slf.columns, MapStr(columns, slf.ctx.escapeName)...)
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
	buf.WriteString(slf.ctx.escapeName(slf.name))
	buf.WriteString(" ON ")
	buf.WriteString(slf.ctx.escapeName(slf.table))
	buf.WriteByte('(')
	// 在输入时已经转义
	buf.WriteString(strings.Join(slf.columns, SQLComma))
	buf.WriteByte(')')
	return buf
}

func (slf *CreateIndexBuilder) ToSQL() string {
	return sqlEndpoint(slf.build())
}

func (slf *CreateIndexBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.db)
}
