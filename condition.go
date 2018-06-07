package gsb

import (
	"bytes"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// 条件之间的关联关系
type ConditionRelation struct {
	ctx    *SQLContext
	buffer *bytes.Buffer
}

func newConditionBuilder(ctx *SQLContext) *ConditionRelation {
	return &ConditionRelation{
		ctx:    ctx,
		buffer: new(bytes.Buffer),
	}
}

// And
func (slf *ConditionRelation) And() *ConditionRelation {
	slf.buffer.WriteString(" AND ")
	return slf
}

// Or
func (slf *ConditionRelation) Or() *ConditionRelation {
	slf.buffer.WriteString(" OR ")
	return slf
}

func (slf *ConditionRelation) Compile() string {
	return slf.buffer.String()
}

//// Conditions

// 等于
func (slf *ConditionRelation) Eq(column string) *ConditionRelation {
	return slf.EqTo(column, slf.ctx.placeHolder)
}

// 等于指定值
func (slf *ConditionRelation) EqTo(column string, value interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " = ", value))
	return slf
}

// 不等于
func (slf *ConditionRelation) NEq(column string) *ConditionRelation {
	return slf.NEqTo(column, slf.ctx.placeHolder)
}

// 不等于指定值
func (slf *ConditionRelation) NEqTo(column string, value interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " <> ", value))
	return slf
}

// 大于
func (slf *ConditionRelation) Gt(column string) *ConditionRelation {
	return slf.GtTo(column, slf.ctx.placeHolder)
}

// 大于指定值
func (slf *ConditionRelation) GtTo(column string, value interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " > ", value))
	return slf
}

// 大于或等于
func (slf *ConditionRelation) GEt(column string) *ConditionRelation {
	return slf.GEtTo(column, slf.ctx.placeHolder)
}

// 大于或等于指定值
func (slf *ConditionRelation) GEtTo(column string, value interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " >= ", value))
	return slf
}

// 小于
func (slf *ConditionRelation) Lt(column string) *ConditionRelation {
	return slf.LEtTo(column, slf.ctx.placeHolder)
}

// 小于指定值
func (slf *ConditionRelation) LtTo(column string, value interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " < ", value))
	return slf
}

// 小于或者等于
func (slf *ConditionRelation) LEt(column string) *ConditionRelation {
	return slf.LEtTo(column, slf.ctx.placeHolder)
}

// 小于或者等于指定值
func (slf *ConditionRelation) LEtTo(column string, value interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " <= ", value))
	return slf
}

// Like
func (slf *ConditionRelation) Like(column string, pattern string) *ConditionRelation {
	slf.buffer.WriteString(slf.op(column, " LIKE ", pattern))
	return slf
}

// Between
func (slf *ConditionRelation) Between(column string, start interface{}, end interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.opv(column, " BETWEEN ", slf.ctx.escapeValue(start)+" AND "+slf.ctx.escapeValue(end)))
	return slf
}

// 在指定集合中
func (slf *ConditionRelation) In(column string, items ...interface{}) *ConditionRelation {
	slf.buffer.WriteString(slf.opv(column, " IN ", brackets(slf.ctx.joinValues(items))))
	return slf
}

////

func (slf *ConditionRelation) opv(name string, op string, value string) string {
	return slf.ctx.escapeName(name) + op + value
}

func (slf *ConditionRelation) op(name string, op string, value interface{}) string {
	return slf.ctx.escapeName(name) + op + slf.ctx.escapeValue(value)
}
