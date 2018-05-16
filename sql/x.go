package sql

import "fmt"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func EscapeName(column string) string {
	if len(column) == 0 {
		panic("Empty column name")
	}
	return "`" + column + "`"
}

func EscapeValue(val interface{}) string {
	if str, ok := val.(string); ok {
		if "?" == str {
			return str
		} else {
			return "'" + str + "'"
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
