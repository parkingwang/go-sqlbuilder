package sqlx

import (
	"bytes"
	"strconv"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type columnDefine struct {
	name    string
	defines string
}

type TableBuilder struct {
	table         string
	columns       []columnDefine
	constraints   []string
	charset       string
	autoIncrement int
	ifNotExists   bool
}

func CreateTable(table string) *TableBuilder {
	return &TableBuilder{
		table:         table,
		columns:       make([]columnDefine, 0),
		constraints:   make([]string, 0),
		charset:       "utf8",
		autoIncrement: 0,
		ifNotExists:   true,
	}
}

func (slf *TableBuilder) IfNotExists(ifIs bool) *TableBuilder {
	slf.ifNotExists = ifIs
	return slf
}

func (slf *TableBuilder) SetCharset(charset string) *TableBuilder {
	slf.charset = charset
	return slf
}

func (slf *TableBuilder) SetAutoIncrement(increment int) *TableBuilder {
	slf.autoIncrement = increment
	return slf
}

func (slf *TableBuilder) Column(name string) *ColumnTypeBuilder {
	return newColumnType(slf, name)
}

func (slf *TableBuilder) addColumn(name string, defines string) {
	for _, d := range slf.columns {
		if d.name == name {
			panic("Duplicated column define, name: " + name)
		}
	}
	slf.columns = append(slf.columns, columnDefine{
		name:    name,
		defines: defines,
	})
}

func (slf *TableBuilder) addConstraint(constraint string) {
	slf.constraints = append(slf.constraints, constraint)
}

func (slf *TableBuilder) build() *bytes.Buffer {
	columns := make([]string, 0)
	for _, define := range slf.columns {
		columns = append(columns, EscapeName(define.name)+define.defines)
	}

	buf := new(bytes.Buffer)
	buf.WriteString("CREATE TABLE ")
	if slf.ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	buf.WriteString(EscapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(strings.Join(append(columns, slf.constraints...), SQLComma))
	buf.WriteByte(')')
	buf.WriteString(" DEFAULT CHARSET=")
	buf.WriteString(slf.charset)
	buf.WriteString(" AUTO_INCREMENT=")
	buf.WriteString(strconv.Itoa(slf.autoIncrement))
	return buf
}

func (slf *TableBuilder) GetSQL() string {
	return makeSQL(slf.build())
}
