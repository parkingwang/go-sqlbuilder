package sql

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//
type SQL interface {
	SQL() string
}

type OrderBy interface {
	OrderBy(columns...string)
}