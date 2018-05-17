package sqlx

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

////

type TableBuilder struct {
	table         string
	columns       map[string]string // 类似：{username: VARCHAR(255) NOT NULL}
	constraints   []string          // 约束列表
	charset       string            // 表字符编码
	autoIncrement int               // 自增编号起步值
}

func CreateTable(table string) *TableBuilder {
	return &TableBuilder{
		table:         table,
		columns:       make(map[string]string),
		constraints:   make([]string, 0),
		charset:       "utf8",
		autoIncrement: 0,
	}
}

func (slf *TableBuilder) Column(name string) *ColumnTypeBuilder {
	return newColumnType(slf, name)
}

func (slf *TableBuilder) SetCharset(charset string) *TableBuilder {
	slf.charset = charset
	return slf
}

func (slf *TableBuilder) SetAutoIncrement(increment int) *TableBuilder {
	slf.autoIncrement = increment
	return slf
}

func (slf *TableBuilder) addColumn(name string, defines string) {
	slf.columns[name] = defines
}

func (slf *TableBuilder) addConstraint(constraint string) {
	slf.constraints = append(slf.constraints, constraint)
}

func (slf *TableBuilder) build() *bytes.Buffer {
	columns := make([]string, 0)
	for name, defines := range slf.columns {
		columns = append(columns, EscapeName(name)+defines)
	}

	buf := new(bytes.Buffer)
	buf.WriteString("CREATE TABLE ")
	buf.WriteString(EscapeName(slf.table))
	buf.WriteByte('(')
	buf.WriteString(strings.Join(append(columns, slf.constraints...), SQLComma))
	buf.WriteByte(')')
	buf.WriteString(" DEFAULT CHARSET=")
	buf.WriteString(slf.charset)
	buf.WriteString(" AUTO_INCREMENT=")
	buf.WriteString(fmt.Sprintf("%d", slf.autoIncrement))
	return buf
}

func (slf *TableBuilder) GetSQL() string {
	return makeSQL(slf.build())
}

func (slf *TableBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.GetSQL(), db)
}
