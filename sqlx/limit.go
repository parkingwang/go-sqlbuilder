package sqlx

import (
	"bytes"
	"database/sql"
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

func (slf *LimitBuilder) Offset(offset int) *LimitBuilder {
	slf.buffer.WriteString(" OFFSET ")
	slf.buffer.WriteString(fmt.Sprintf("%d", offset))
	return slf
}

func (slf *LimitBuilder) SQL() string {
	return endpoint(slf.buffer)
}

func (slf *LimitBuilder) Execute(db *sql.DB) *Executor {
	return newExecute(slf.SQL(), db)
}
