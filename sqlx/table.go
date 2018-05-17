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
	constraints   []string            // 通用的约束列表
	uniques       map[string][]string // Unique约束列表，根据名称来合并其Column。默认合并在 “” 组
	charset       string
	autoIncrement int
	ifNotExists   bool
}

func CreateTable(table string) *TableBuilder {
	return &TableBuilder{
		table:         table,
		columns:       make([]columnDefine, 0),
		constraints:   make([]string, 0),
		uniques:       make(map[string][]string),
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

func (slf *TableBuilder) addUnique(name string, column string) {
	if exists, ok := slf.uniques[name]; ok {
		slf.uniques[name] = append(exists, column)
	} else {
		slf.uniques[name] = append(make([]string, 0), column)
	}
}

func (slf *TableBuilder) addConstraint(constraint string) {
	slf.constraints = append(slf.constraints, constraint)
}

func (slf *TableBuilder) compile() *bytes.Buffer {
	// 数据列
	columns := make([]string, 0)
	for _, define := range slf.columns {
		columns = append(columns, EscapeName(define.name)+define.defines)
	}

	// 通用约束
	columns = append(columns, slf.constraints...)

	// Unique约束列
	for name, colNames := range slf.uniques {
		constraint := namedConstraint(name) + "UNIQUE (" + strings.Join(colNames, SQLComma) + ")"
		columns = append(columns, constraint)
	}

	buf := new(bytes.Buffer)
	buf.WriteString("CREATE TABLE ")
	if slf.ifNotExists {
		buf.WriteString("IF NOT EXISTS ")
	}
	buf.WriteString(EscapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(strings.Join(columns, SQLComma))
	buf.WriteByte(')')
	buf.WriteString(" DEFAULT CHARSET=")
	buf.WriteString(slf.charset)
	buf.WriteString(" AUTO_INCREMENT=")
	buf.WriteString(strconv.Itoa(slf.autoIncrement))
	return buf
}

func (slf *TableBuilder) GetSQL() string {
	return endOfSQL(slf.compile())
}

func namedConstraint(name string) string {
	if len(name) > 0 {
		return "CONSTRAINT " + EscapeName(name) + SQLSpace
	} else {
		return ""
	}
}
