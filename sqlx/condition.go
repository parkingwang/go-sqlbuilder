package sqlx

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type Condition struct {
	buffer *bytes.Buffer
}

func newCondition() *Condition {
	return &Condition{
		buffer: new(bytes.Buffer),
	}
}

func (slf *Condition) And() *Condition {
	slf.buffer.WriteString(" AND ")
	return slf
}

func (slf *Condition) Or() *Condition {
	slf.buffer.WriteString(" OR ")
	return slf
}

func (slf *Condition) SQL() string {
	return slf.buffer.String()
}

//// Conditions

func Equal(column string) *Condition {
	return newCondition().Equal(column)
}

func (slf *Condition) Equal(column string) *Condition {
	return slf.EqualTo(column, SQLPlaceHolder)
}

func EqualTo(column string, value interface{}) *Condition {
	return newCondition().EqualTo(column, value)
}

func (slf *Condition) EqualTo(column string, value interface{}) *Condition {
	slf.buffer.WriteString(opEscape(column, " = ", value))
	return slf
}

func NotEqual(column string) *Condition {
	return newCondition().NotEqual(column)
}

func (slf *Condition) NotEqual(column string) *Condition {
	return slf.NotEqualTo(column, SQLPlaceHolder)
}

func NotEqualTo(column string, value interface{}) *Condition {
	return newCondition().NotEqualTo(column, value)
}

func (slf *Condition) NotEqualTo(column string, value interface{}) *Condition {
	slf.buffer.WriteString(opEscape(column, " <> ", value))
	return slf
}

func GreaterThen(column string) *Condition {
	return newCondition().GreaterThen(column)
}

func (slf *Condition) GreaterThen(column string) *Condition {
	return slf.GreaterThenTo(column, SQLPlaceHolder)
}

func GreaterThenTo(column string, value interface{}) *Condition {
	return newCondition().GreaterThenTo(column, value)
}

func (slf *Condition) GreaterThenTo(column string, value interface{}) *Condition {
	slf.buffer.WriteString(opEscape(column, " > ", value))
	return slf
}

func GreaterEqualThen(column string) *Condition {
	return newCondition().GreaterEqualThen(column)
}

func (slf *Condition) GreaterEqualThen(column string) *Condition {
	return slf.GreaterEqualThenTo(column, SQLPlaceHolder)
}

func GreaterEqualThenTo(column string, value interface{}) *Condition {
	return newCondition().GreaterEqualThenTo(column, value)
}

func (slf *Condition) GreaterEqualThenTo(column string, value interface{}) *Condition {
	slf.buffer.WriteString(opEscape(column, " >= ", value))
	return slf
}

func LessThen(column string) *Condition {
	return newCondition().LessThen(column)
}

func (slf *Condition) LessThen(column string) *Condition {
	return slf.LessEqualThenTo(column, SQLPlaceHolder)
}

func LessThenTo(column string, value interface{}) *Condition {
	return newCondition().LessThenTo(column, value)
}

func (slf *Condition) LessThenTo(column string, value interface{}) *Condition {
	slf.buffer.WriteString(opEscape(column, " < ", value))
	return slf
}

func LessEqualThen(column string) *Condition {
	return newCondition().LessEqualThen(column)
}

func (slf *Condition) LessEqualThen(column string) *Condition {
	return slf.LessEqualThenTo(column, SQLPlaceHolder)
}

func LessEqualThenTo(column string, value interface{}) *Condition {
	return newCondition().LessEqualThenTo(column, value)
}

func (slf *Condition) LessEqualThenTo(column string, value interface{}) *Condition {
	slf.buffer.WriteString(opEscape(column, " <= ", value))
	return slf
}

//

func (slf *Condition) Like(column string, pattern string) *Condition {
	slf.buffer.WriteString(opEscape(column, " LIKE ", pattern))
	return slf
}

func (slf *Condition) Between(column string, start interface{}, end interface{}) *Condition {
	slf.buffer.WriteString(opIgnore(column, " BETWEEN ", EscapeValue(start)+" AND "+EscapeValue(end)))
	return slf
}

func (slf *Condition) In(column string, items ...interface{}) *Condition {
	slf.buffer.WriteString(opIgnore(column, " IN ", wrapBrackets(joinValues(items))))
	return slf
}
