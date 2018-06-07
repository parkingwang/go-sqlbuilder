package sqlbuilder

import (
	"fmt"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// SQL Builder
//

const SQLComma = ", "
const SQLSpace = " "

type SQLContext struct {
	prepare     SQLPrepare
	placeHolder string
	nameEscape  string
	valueEscape string
}

func NewContext() *SQLContext {
	return NewContextWith(nil)
}

func NewContextWith(prepare SQLPrepare) *SQLContext {
	return &SQLContext{
		prepare:     prepare,
		placeHolder: "?",
		nameEscape:  "`",
		valueEscape: "'",
	}
}

func (slf *SQLContext) SetPlaceHolder(ph string) *SQLContext {
	slf.placeHolder = ph
	return slf
}

func (slf *SQLContext) SetNameEscapeChar(char string) *SQLContext {
	slf.nameEscape = char
	return slf
}

func (slf *SQLContext) SetValueEscapeChar(char string) *SQLContext {
	slf.valueEscape = char
	return slf
}

func (slf *SQLContext) Insert(table string) *InsertBuilder {
	return newInsertBuilder(slf, table)
}

func (slf *SQLContext) Delete(table string) *DeleteBuilder {
	return newDeleteBuilder(slf, table)
}

func (slf *SQLContext) Update(table string) *UpdateBuilder {
	return newUpdateBuilder(slf, table)
}

func (slf *SQLContext) Select(column string, otherColumns ...string) *SelectBuilder {
	return newSelectBuilder(slf, column, otherColumns...)
}

func (slf *SQLContext) CreateIndex(indexName string) *CreateIndexBuilder {
	return newCreateIndex(slf, indexName)
}

func (slf *SQLContext) DropIndex(indexName string) *DropIndexBuilder {
	return newDropIndexBuilder(slf, indexName)
}

func (slf *SQLContext) CreateTable(tableName string) *CreateTableBuilder {
	return newCreateTableBuilder(slf, tableName)
}

//// Conditions

// 等于
func (slf *SQLContext) Eq(column string) *ConditionRelation {
	return newConditionBuilder(slf).Eq(column)
}

// 等于指定值
func (slf *SQLContext) EqTo(column string, value interface{}) *ConditionRelation {
	return newConditionBuilder(slf).EqTo(column, value)
}

// 不等于
func (slf *SQLContext) NEq(column string) *ConditionRelation {
	return newConditionBuilder(slf).NEq(column)
}

// 不等于指定值
func (slf *SQLContext) NEqTo(column string, value interface{}) *ConditionRelation {
	return newConditionBuilder(slf).NEqTo(column, value)
}

// 大于
func (slf *SQLContext) Gt(column string) *ConditionRelation {
	return newConditionBuilder(slf).Gt(column)
}

// 大于指定值
func (slf *SQLContext) GtTo(column string, value interface{}) *ConditionRelation {
	return newConditionBuilder(slf).GtTo(column, value)
}

// 大于或等于
func (slf *SQLContext) GEt(column string) *ConditionRelation {
	return newConditionBuilder(slf).GEt(column)
}

// 大于或等于指定值
func (slf *SQLContext) GEtTo(column string, value interface{}) *ConditionRelation {
	return newConditionBuilder(slf).GEtTo(column, value)
}

// 小于
func (slf *SQLContext) Lt(column string) *ConditionRelation {
	return newConditionBuilder(slf).Lt(column)
}

// 小于指定值
func (slf *SQLContext) LtTo(column string, value interface{}) *ConditionRelation {
	return newConditionBuilder(slf).LtTo(column, value)
}

// 小于或者等于
func (slf *SQLContext) LEt(column string) *ConditionRelation {
	return newConditionBuilder(slf).LEt(column)
}

// 小于或者等于指定值
func (slf *SQLContext) LEtTo(column string, value interface{}) *ConditionRelation {
	return newConditionBuilder(slf).LEtTo(column, value)
}

////

func (slf *SQLContext) escapeName(name string) string {
	if len(slf.nameEscape) <= 0 {
		return name
	}

	if len(name) == 0 {
		panic("Empty name")
	}

	if "*" == name {
		return name
	}

	// : count(id) as count
	if strings.IndexAny(name, " ") > 0 {
		return name
	}

	return slf.nameEscape + name + slf.nameEscape
}

func (slf *SQLContext) escapeValue(val interface{}) string {
	if strValue, ok := val.(string); ok {
		if slf.placeHolder == strValue {
			return strValue
		} else {
			return slf.valueEscape + strValue + slf.valueEscape
		}
	} else {
		return fmt.Sprintf("%v", val)
	}
}

////

func (slf *SQLContext) joinNames(items []string) string {
	return strings.Join(MapStr(items, slf.escapeName), SQLComma)
}

func (slf *SQLContext) joinValues(values []interface{}) string {
	return strings.Join(MapAny(values, slf.escapeValue), SQLComma)
}

func (slf *SQLContext) wrapOP(name string, op string, value interface{}) string {
	return slf.escapeName(name) + op + slf.escapeValue(value)
}

func MapStr(items []string, mapper func(string) string) []string {
	out := make([]string, len(items))
	for i, v := range items {
		out[i] = mapper(v)
	}
	return out
}

func MapAny(items []interface{}, mapper func(interface{}) string) []string {
	out := make([]string, len(items))
	for i, v := range items {
		out[i] = mapper(v)
	}
	return out
}
