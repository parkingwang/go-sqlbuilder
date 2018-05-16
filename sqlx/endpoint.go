package sqlx

import (
	"bytes"
	"database/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//
type SQL interface {
	SQL() string
}

type Execute interface {
	Execute(db *sql.DB) *Executor
}

func endpoint(buffer *bytes.Buffer) string {
	buffer.WriteByte(';')
	return buffer.String()
}
