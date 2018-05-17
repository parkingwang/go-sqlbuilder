package sqlx

import (
	"bytes"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type ColumnTypeBuilder struct {
	table *TableBuilder

	name   string
	buffer *bytes.Buffer
}

func newColumnType(table *TableBuilder, name string) *ColumnTypeBuilder {
	ctb := &ColumnTypeBuilder{
		table:  table,
		name:   name,
		buffer: new(bytes.Buffer),
	}
	return ctb
}

// -128 到 127 常规。0 到 255 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) TinyInt(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("TINYINT", size)
	return slf
}

// -32768 到 32767 常规。0 到 65535 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) SmallInt(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("SMALLINT", size)
	return slf
}

// -8388608 到 8388607 普通。0 to 16777215 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) MediumInt(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("MEDIUMINT", size)
	return slf
}

// -9223372036854775808 到 9223372036854775807 常规。0 到 18446744073709551615 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) BigInt(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("BIGINT", size)
	return slf
}

// -2147483648 到 2147483647 常规。0 到 4294967295 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) Int(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("INT", size)
	return slf
}

// 带有浮动小数点的小数字。在括号中规定最大位数。在 d 参数中规定小数点右侧的最大位数。
func (slf *ColumnTypeBuilder) Float(size int, d int) *ColumnTypeBuilder {
	slf.addKeyWithSize2("FLOAT", size, d)
	return slf
}

// 带有浮动小数点的大数字。在括号中规定最大位数。在 d 参数中规定小数点右侧的最大位数。
func (slf *ColumnTypeBuilder) Double(size int, d int) *ColumnTypeBuilder {
	slf.addKeyWithSize2("Double", size, d)
	return slf
}

// 作为字符串存储的 DOUBLE 类型，允许固定的小数点。
func (slf *ColumnTypeBuilder) Decimal(size int, d int) *ColumnTypeBuilder {
	slf.addKeyWithSize2("DECIMAL", size, d)
	return slf
}

// 保存固定长度的字符串（可包含字母、数字以及特殊字符）。在括号中指定字符串的长度。最多 255 个字符。
func (slf *ColumnTypeBuilder) Char(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("CHAR", size)
	return slf
}

// 保存可变长度的字符串（可包含字母、数字以及特殊字符）。在括号中指定字符串的最大长度。最多 255 个字符。
func (slf *ColumnTypeBuilder) VarChar(size int) *ColumnTypeBuilder {
	slf.addKeyWithSize("VARCHAR", size)
	return slf
}

// 存放最大长度为 255 个字符的字符串。
func (slf *ColumnTypeBuilder) TinyText() *ColumnTypeBuilder {
	slf.addKey("TINYTEXT")
	return slf
}

// 存放最大长度为 65,535 个字符的字符串。
func (slf *ColumnTypeBuilder) Text() *ColumnTypeBuilder {
	slf.addKey("TEXT")
	return slf
}

// 存放最大长度为 16,777,215 个字符的字符串。
func (slf *ColumnTypeBuilder) MediumText() *ColumnTypeBuilder {
	slf.addKey("MEDIUMTEXT")
	return slf
}

// 存放最大长度为 4,294,967,295 个字符的字符串。
func (slf *ColumnTypeBuilder) LongText() *ColumnTypeBuilder {
	slf.addKey("LONGTEXT")
	return slf
}

// 用于 BLOBs (Binary Large OBjects)。存放最多 65,535 字节的数据。
func (slf *ColumnTypeBuilder) Blob() *ColumnTypeBuilder {
	slf.addKey("BLOB")
	return slf
}

// 用于 BLOBs (Binary Large OBjects)。存放最多 16,777,215 字节的数据。
func (slf *ColumnTypeBuilder) MediumBlob() *ColumnTypeBuilder {
	slf.addKey("MEDIUMBLOB")
	return slf
}

// 用于 BLOBs (Binary Large OBjects)。存放最多 4,294,967,295 字节的数据。
func (slf *ColumnTypeBuilder) LongBlob() *ColumnTypeBuilder {
	slf.addKey("LONGBLOB")
	return slf
}

// SET 最多只能包含 64 个列表项，不过 SET 可存储一个以上的值。
func (slf *ColumnTypeBuilder) Set() *ColumnTypeBuilder {
	slf.addKey("SET")
	return slf
}

// 日期。格式：YYYY-MM-DD
// 支持的范围是从 '1000-01-01' 到 '9999-12-31'
func (slf *ColumnTypeBuilder) Date() *ColumnTypeBuilder {
	slf.addKey("DATE")
	return slf
}

// 日期和时间的组合。格式：YYYY-MM-DD HH:MM:SS
// 支持的范围是从 '1000-01-01 00:00:00' 到 '9999-12-31 23:59:59'
func (slf *ColumnTypeBuilder) DateTime() *ColumnTypeBuilder {
	slf.addKey("DATETIME")
	return slf
}

// *时间戳。TIMESTAMP 值使用 Unix 纪元('1970-01-01 00:00:00' UTC) 至今的描述来存储。格式：YYYY-MM-DD HH:MM:SS
// 支持的范围是从 '1970-01-01 00:00:01' UTC 到 '2038-01-09 03:14:07' UTC
func (slf *ColumnTypeBuilder) Timestamp() *ColumnTypeBuilder {
	slf.addKey("TIMESTAMP")
	return slf
}

// 时间。格式：HH:MM:SS 注释：支持的范围是从 '-838:59:59' 到 '838:59:59'
func (slf *ColumnTypeBuilder) Time() *ColumnTypeBuilder {
	slf.addKey("TIME")
	return slf
}

//

func (slf *ColumnTypeBuilder) Unique() *ColumnTypeBuilder {
	slf.table.addConstraint(fmt.Sprintf("UNIQUE(%s)", EscapeName(slf.name)))
	return slf
}

func (slf *ColumnTypeBuilder) PrimaryKey() *ColumnTypeBuilder {
	slf.table.addConstraint(fmt.Sprintf("PRIMARY KEY(%s)", EscapeName(slf.name)))
	return slf
}

func (slf *ColumnTypeBuilder) ForeignKey(refTableName string, refColumnName string) *ColumnTypeBuilder {
	slf.table.addConstraint(fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s(%s)",
		slf.name, EscapeName(refTableName), EscapeName(refColumnName)))
	return slf
}

func (slf *ColumnTypeBuilder) NotNull() *ColumnTypeBuilder {
	slf.addKey("NOT NULL")
	return slf
}

func (slf *ColumnTypeBuilder) AutoIncrement() *ColumnTypeBuilder {
	slf.addKey("AUTO_INCREMENT")
	return slf
}

//

func (slf *ColumnTypeBuilder) DefaultNow() *ColumnTypeBuilder {
	return slf.Default("NOW()")
}

func (slf *ColumnTypeBuilder) Default0() *ColumnTypeBuilder {
	return slf.Default(0)
}

func (slf *ColumnTypeBuilder) DefaultEmptyString() *ColumnTypeBuilder {
	return slf.Default("''")
}

func (slf *ColumnTypeBuilder) DefaultNull() *ColumnTypeBuilder {
	return slf.Default("NULL")
}

func (slf *ColumnTypeBuilder) Default(value interface{}) *ColumnTypeBuilder {
	slf.addKey("DEFAULT " + EscapeValue(value))
	return slf
}

//

func (slf *ColumnTypeBuilder) Column(name string) *ColumnTypeBuilder {
	slf.columnDefineComplete()
	return newColumnType(slf.table, name)
}

func (slf *ColumnTypeBuilder) GetSQL() string {
	slf.columnDefineComplete()
	return slf.table.GetSQL()
}

func (slf *ColumnTypeBuilder) Execute(prepare SQLPrepare) *Executor {
	return newExecute(slf.GetSQL(), prepare)
}

//

func (slf *ColumnTypeBuilder) columnDefineComplete() {
	slf.table.addColumn(slf.name, slf.buffer.String())
}

func (slf *ColumnTypeBuilder) addKeyWithSize(key string, size int) {
	slf.buffer.WriteString(SQLSpace)
	slf.buffer.WriteString(fmt.Sprintf("%s(%d)", key, size))
}

func (slf *ColumnTypeBuilder) addKeyWithSize2(key string, size1 int, size2 int) {
	slf.buffer.WriteString(SQLSpace)
	slf.buffer.WriteString(fmt.Sprintf("%s(%d,%d)", key, size1, size2))
}

func (slf *ColumnTypeBuilder) addKey(key string) {
	slf.buffer.WriteString(SQLSpace)
	slf.buffer.WriteString(key)
}
