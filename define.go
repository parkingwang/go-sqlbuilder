package gsb

import (
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

const SQLPlaceHolder = "?"
const SQLComma = ", "
const SQLSpace = " "
const SQLNameEscape = "`"
const SQLStringValueEscape = "'"

type SQLStatement interface {
	Compile() string
}

type SQLGenerator interface {
	ToSQL() string
}

type Execute interface {
	Execute(db *sql.DB) *Executor
}
