package sql

import (
	"bytes"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type LimitBuilder struct {
	buffer *bytes.Buffer
}

func newLimit(buffer *bytes.Buffer, limit int) *LimitBuilder {
	lb := &LimitBuilder{
		buffer: buffer,
	}
	lb.buffer.WriteString(" LIMIT ")
	lb.buffer.WriteString(fmt.Sprintf("%d", limit))
	return lb
}

func (lb *LimitBuilder) Offset(offset int) string {
	lb.buffer.WriteString(" OFFSET ")
	lb.buffer.WriteString(fmt.Sprintf("%d", offset))
	return endpoint(lb.buffer)
}

func (lb *LimitBuilder) SQL() string {
	return endpoint(lb.buffer)
}
