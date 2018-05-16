package sqlx

import (
	"bytes"
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//
type Statement interface {
	SQL() string
}

type MakeSQL interface {
	MakeSQL() string
}

type Execute interface {
	Execute(db *sql.DB) *Executor
}

func makeSQL(buffer *bytes.Buffer) string {
	buffer.WriteByte(';')
	return buffer.String()
}
