package gsb

import (
	"bytes"
	"fmt"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func EscapeName(name string) string {
	if len(name) == 0 {
		panic("Empty name")
	}
	return SQLNameEscape + name + SQLNameEscape
}

func EscapeValue(val interface{}) string {
	if strValue, ok := val.(string); ok {
		if SQLPlaceHolder == strValue {
			return strValue
		} else {
			return SQLStringValueEscape + strValue + SQLStringValueEscape
		}
	} else {
		return fmt.Sprintf("%v", val)
	}
}

func Map0(items []interface{}, mapper func(interface{}) string) []string {
	newItems := make([]string, len(items))
	for i, v := range items {
		newItems[i] = mapper(v)
	}
	return newItems
}

func Map(items []string, mapper func(string) string) []string {
	newItems := make([]string, len(items))
	for i, v := range items {
		newItems[i] = mapper(v)
	}
	return newItems
}

func brackets(name string) string {
	if len(name) == 0 {
		panic("Empty name")
	}
	return "(" + name + ")"
}

func joinNames(items []string) string {
	return strings.Join(Map(items, EscapeName), SQLComma)
}

func joinValues(values []interface{}) string {
	return strings.Join(Map0(values, EscapeValue), SQLComma)
}

func op(name string, op string, value interface{}) string {
	return EscapeName(name) + op + EscapeValue(value)
}

func endOfSQL(buffer *bytes.Buffer) string {
	buffer.WriteByte(';')
	return buffer.String()
}
