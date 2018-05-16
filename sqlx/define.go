package sqlx

import (
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

const SQLPlaceHolder = "?"
const SQLComma = ", "
const SQLNameEscape = "`"
const SQLStringValueEscape = "'"

type Statement interface {
	SQL() string
}

type MakeSQL interface {
	MakeSQL() string
}

type Execute interface {
	Execute(db *sql.DB) *Executor
}
