package sqlbuilder

import (
	"bytes"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type ColumnTypeBuilder struct {
	ctx    *SQLContext
	table  *CreateTableBuilder
	name   string
	buffer *bytes.Buffer
}

func newColumnType(ctx *SQLContext, table *CreateTableBuilder, name string) *ColumnTypeBuilder {
	ctb := &ColumnTypeBuilder{
		ctx:    ctx,
		table:  table,
		name:   name,
		buffer: new(bytes.Buffer),
	}
	return ctb
}

// -128 到 127 常规。0 到 255 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) TinyInt(size int) *ColumnTypeBuilder {
	return slf.keySize("TINYINT", size)
}

// -32768 到 32767 常规。0 到 65535 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) SmallInt(size int) *ColumnTypeBuilder {
	return slf.keySize("SMALLINT", size)
}

// -8388608 到 8388607 普通。0 to 16777215 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) MediumInt(size int) *ColumnTypeBuilder {
	return slf.keySize("MEDIUMINT", size)
}

// -9223372036854775808 到 9223372036854775807 常规。0 到 18446744073709551615 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) BigInt(size int) *ColumnTypeBuilder {
	return slf.keySize("BIGINT", size)
}

// -2147483648 到 2147483647 常规。0 到 4294967295 无符号*。在括号中规定最大位数。
func (slf *ColumnTypeBuilder) Int(size int) *ColumnTypeBuilder {
	return slf.keySize("INT", size)
}

// 带有浮动小数点的小数字。在括号中规定最大位数。在 d 参数中规定小数点右侧的最大位数。
func (slf *ColumnTypeBuilder) Float(size int, d int) *ColumnTypeBuilder {
	return slf.keySize2("FLOAT", size, d)
}

// 带有浮动小数点的大数字。在括号中规定最大位数。在 d 参数中规定小数点右侧的最大位数。
func (slf *ColumnTypeBuilder) Double(size int, d int) *ColumnTypeBuilder {
	return slf.keySize2("Double", size, d)
}

// 作为字符串存储的 DOUBLE 类型，允许固定的小数点。
func (slf *ColumnTypeBuilder) Decimal(size int, d int) *ColumnTypeBuilder {
	return slf.keySize2("DECIMAL", size, d)
}

// 保存固定长度的字符串（可包含字母、数字以及特殊字符）。在括号中指定字符串的长度。最多 255 个字符。
func (slf *ColumnTypeBuilder) Char(size int) *ColumnTypeBuilder {
	return slf.keySize("CHAR", size)
}

// 保存可变长度的字符串（可包含字母、数字以及特殊字符）。在括号中指定字符串的最大长度。最多 255 个字符。
func (slf *ColumnTypeBuilder) VarChar(size int) *ColumnTypeBuilder {
	return slf.keySize("VARCHAR", size)
}

// 存放最大长度为 255 个字符的字符串。
func (slf *ColumnTypeBuilder) TinyText() *ColumnTypeBuilder {
	return slf.key("TINYTEXT")
}

// 存放最大长度为 65,535 个字符的字符串。
func (slf *ColumnTypeBuilder) Text() *ColumnTypeBuilder {
	return slf.key("TEXT")
}

// 存放最大长度为 16,777,215 个字符的字符串。
func (slf *ColumnTypeBuilder) MediumText() *ColumnTypeBuilder {
	return slf.key("MEDIUMTEXT")
}

// 存放最大长度为 4,294,967,295 个字符的字符串。
func (slf *ColumnTypeBuilder) LongText() *ColumnTypeBuilder {
	return slf.key("LONGTEXT")
}

// 用于 BLOBs (Binary Large OBjects)。存放最多 65,535 字节的数据。
func (slf *ColumnTypeBuilder) Blob() *ColumnTypeBuilder {
	return slf.key("BLOB")
}

// 用于 BLOBs (Binary Large OBjects)。存放最多 16,777,215 字节的数据。
func (slf *ColumnTypeBuilder) MediumBlob() *ColumnTypeBuilder {
	return slf.key("MEDIUMBLOB")
}

// 用于 BLOBs (Binary Large OBjects)。存放最多 4,294,967,295 字节的数据。
func (slf *ColumnTypeBuilder) LongBlob() *ColumnTypeBuilder {
	return slf.key("LONGBLOB")
}

// SET 最多只能包含 64 个列表项，不过 SET 可存储一个以上的值。
func (slf *ColumnTypeBuilder) Set() *ColumnTypeBuilder {
	return slf.key("SET")
}

// 日期。格式：YYYY-MM-DD
// 支持的范围是从 '1000-01-01' 到 '9999-12-31'
func (slf *ColumnTypeBuilder) Date() *ColumnTypeBuilder {
	return slf.key("DATE")
}

// 日期和时间的组合。格式：YYYY-MM-DD HH:MM:SS
// 支持的范围是从 '1000-01-01 00:00:00' 到 '9999-12-31 23:59:59'
func (slf *ColumnTypeBuilder) DateTime() *ColumnTypeBuilder {
	return slf.key("DATETIME")
}

// *时间戳。TIMESTAMP 值使用 Unix 纪元('1970-01-01 00:00:00' UTC) 至今的描述来存储。格式：YYYY-MM-DD HH:MM:SS
// 支持的范围是从 '1970-01-01 00:00:01' UTC 到 '2038-01-09 03:14:07' UTC
func (slf *ColumnTypeBuilder) Timestamp() *ColumnTypeBuilder {
	return slf.key("TIMESTAMP")
}

// 时间。格式：HH:MM:SS 注释：支持的范围是从 '-838:59:59' 到 '838:59:59'
func (slf *ColumnTypeBuilder) Time() *ColumnTypeBuilder {
	return slf.key("TIME")
}

//

func (slf *ColumnTypeBuilder) Unique() *ColumnTypeBuilder {
	return slf.UniqueNamed("")
}

func (slf *ColumnTypeBuilder) UniqueNamed(name string) *ColumnTypeBuilder {
	slf.table.addUnique(name, slf.ctx.escapeName(slf.name))
	return slf
}

func (slf *ColumnTypeBuilder) PrimaryKey() *ColumnTypeBuilder {
	slf.table.addConstraint(fmt.Sprintf("PRIMARY KEY(%s)", slf.ctx.escapeName(slf.name)))
	return slf
}

func (slf *ColumnTypeBuilder) ForeignKeyNamed(fkName string, refTableName string, refColumnName string) *ColumnTypeBuilder {
	slf.table.addConstraint(fmt.Sprintf("%sFOREIGN KEY (%s) REFERENCES %s(%s)",
		namedConstraint(slf.ctx, fkName),
		slf.ctx.escapeName(slf.name),
		slf.ctx.escapeName(refTableName),
		slf.ctx.escapeName(refColumnName)))
	return slf
}

func (slf *ColumnTypeBuilder) ForeignKey(refTableName string, refColumnName string) *ColumnTypeBuilder {
	return slf.ForeignKeyNamed("", refTableName, refColumnName)
}

func (slf *ColumnTypeBuilder) NotNull() *ColumnTypeBuilder {
	return slf.key("NOT NULL")
}

func (slf *ColumnTypeBuilder) AutoIncrement() *ColumnTypeBuilder {
	return slf.key("AUTO_INCREMENT")
}

//

func (slf *ColumnTypeBuilder) DefaultNow() *ColumnTypeBuilder {
	return slf.key("DEFAULT NOW()")
}

func (slf *ColumnTypeBuilder) Default0() *ColumnTypeBuilder {
	return slf.key("DEFAULT 0")
}

func (slf *ColumnTypeBuilder) DefaultEmptyString() *ColumnTypeBuilder {
	return slf.key("DEFAULT ''")
}

func (slf *ColumnTypeBuilder) DefaultNull() *ColumnTypeBuilder {
	return slf.key("DEFAULT NULL")
}

func (slf *ColumnTypeBuilder) Default(value interface{}) *ColumnTypeBuilder {
	return slf.key("DEFAULT " + slf.ctx.escapeValue(value))
}

//

func (slf *ColumnTypeBuilder) Column(name string) *ColumnTypeBuilder {
	slf.compile()
	return newColumnType(slf.ctx, slf.table, name)
}

func (slf *ColumnTypeBuilder) ToSQL() string {
	slf.compile()
	return slf.table.ToSQL()
}

func (slf *ColumnTypeBuilder) Execute() *Executor {
	return newExecute(slf.ToSQL(), slf.ctx.prepare)
}

//

func (slf *ColumnTypeBuilder) compile() {
	slf.table.addColumn(slf.name, slf.buffer.String())
}

func (slf *ColumnTypeBuilder) keySize(key string, size int) *ColumnTypeBuilder {
	slf.buffer.WriteString(SQLSpace)
	slf.buffer.WriteString(fmt.Sprintf("%s(%d)", key, size))
	return slf
}

func (slf *ColumnTypeBuilder) keySize2(key string, size1 int, size2 int) *ColumnTypeBuilder {
	slf.buffer.WriteString(SQLSpace)
	slf.buffer.WriteString(fmt.Sprintf("%s(%d,%d)", key, size1, size2))
	return slf
}

func (slf *ColumnTypeBuilder) key(key string) *ColumnTypeBuilder {
	slf.buffer.WriteString(SQLSpace)
	slf.buffer.WriteString(key)
	return slf
}
