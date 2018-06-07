package gsb

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
const SQLMaxColumns = 255
const SQLDefaultColumns = 61

type SQLBuilder struct {
	placeHolder string
	nameEscape  string
	valueEscape string
}

func New() *SQLBuilder {
	return &SQLBuilder{
		placeHolder: "?",
		nameEscape:  "`",
		valueEscape: "'",
	}
}

func (slf *SQLBuilder) SetPlaceHolder(ph string) *SQLBuilder {
	slf.placeHolder = ph
	return slf
}

func (slf *SQLBuilder) SetNameEscapeChar(char string) *SQLBuilder {
	slf.nameEscape = char
	return slf
}

func (slf *SQLBuilder) SetValueEscapeChar(char string) *SQLBuilder {
	slf.valueEscape = char
	return slf
}

func (slf *SQLBuilder) Insert(table string) *InsertBuilder {
	return newInsertBuilder(slf, table)
}

func (slf *SQLBuilder) Delete(table string) *DeleteBuilder {
	return newDeleteBuilder(slf, table)
}

////

func (slf *SQLBuilder) EscapeName(name string) string {
	if len(name) == 0 {
		panic("Empty name")
	}
	return slf.nameEscape + name + slf.nameEscape
}

func (slf *SQLBuilder) EscapeValue(val interface{}) string {
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

func (slf *SQLBuilder) JoinNames(items []string) string {
	return strings.Join(MapStr(items, slf.EscapeName), SQLComma)
}

func (slf *SQLBuilder) JoinValues(values []interface{}) string {
	return strings.Join(MapAny(values, slf.EscapeValue), SQLComma)
}

func (slf *SQLBuilder) WrapOP(name string, op string, value interface{}) string {
	return slf.EscapeName(name) + op + slf.EscapeValue(value)
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
