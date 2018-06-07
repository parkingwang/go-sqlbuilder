package gsb

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func brackets(name string) string {
	if len(name) == 0 {
		panic("Empty name")
	}
	return "(" + name + ")"
}

func sqlEndpoint(buffer *bytes.Buffer) string {
	buffer.WriteByte(';')
	return buffer.String()
}
