package gsb

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type DropIndexBuilder struct {
	name  string
	table string
}

func DropIndex(indexName string) *DropIndexBuilder {
	return &DropIndexBuilder{
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
	buf.WriteString(EscapeName(slf.table))
	buf.WriteString(" DROP INDEX ")
	buf.WriteString(EscapeName(slf.name))
	return buf
}

func (slf *DropIndexBuilder) GetSQL() string {
	return endOfSQL(slf.compile())
}

func (slf *DropIndexBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}
