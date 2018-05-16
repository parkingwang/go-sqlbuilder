package sql

import "bytes"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//
type SQL interface {
	SQL() string
}

func endpoint(buffer *bytes.Buffer) string {
	buffer.WriteByte(';')
	return buffer.String()
}
