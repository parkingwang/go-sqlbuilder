package sqlbuilder

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type DropIndexBuilder struct {
	ctx   *SQLContext
	name  string
	table string
}

func newDropIndexBuilder(ctx *SQLContext, indexName string) *DropIndexBuilder {
	return &DropIndexBuilder{
		ctx:  ctx,
		name: indexName,
	}
}

func (slf *DropIndexBuilder) OnTable(table string) *DropIndexBuilder {
	slf.table = table
	return slf
}

func (slf *DropIndexBuilder) compile() *bytes.Buffer {
	if "" == slf.table {
		panic("table not found, you should call 'OnTable(table)' method to set it")
	}
	//ALTER TABLE table_name DROP INDEX index_name
	buf := new(bytes.Buffer)
	buf.WriteString("ALTER TABLE ")
	buf.WriteString(slf.ctx.escapeName(slf.table))
	buf.WriteString(" DROP INDEX ")
	buf.WriteString(slf.ctx.escapeName(slf.name))
	return buf
}

func (slf *DropIndexBuilder) ToSQL() string {
	return sqlEndpoint(slf.compile())
}

func (slf *DropIndexBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.prepare)
}
