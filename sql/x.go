package sql

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func Escape(column string) string {
	return "`" + column + "`"
}

func Map(items []string, mapper func(string) string) []string {
	newItems := make([]string, len(items))
	for i, v := range items {
		newItems[i] = mapper(v)
	}
	return newItems
}

