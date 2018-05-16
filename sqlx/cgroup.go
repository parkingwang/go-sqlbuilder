package sqlx

import "bytes"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type ConditionGroup struct {
	buffer     *bytes.Buffer // 生成语句的缓存
	conditions SQLStatement  // 由外部指定的，用于使用括号包装起来的条件语句
}

func Group(conditions SQLStatement) *ConditionGroup {
	return newGroup(conditions)
}

func (slf *ConditionGroup) Group(conditions SQLStatement) *ConditionGroup {
	slf.conditions = conditions
	return slf
}

func (slf *ConditionGroup) And() *ConditionGroup {
	return newGroupWith(slf, " AND ")
}

func (slf *ConditionGroup) Or() *ConditionGroup {
	return newGroupWith(slf, " OR ")
}

func (slf *ConditionGroup) Statement() string {
	slf.buffer.WriteString(wrapBrackets(slf.conditions.Statement()))
	return slf.buffer.String()
}

//

func newGroup(conditions SQLStatement) *ConditionGroup {
	return &ConditionGroup{
		buffer:     new(bytes.Buffer),
		conditions: conditions,
	}
}

func newGroupWith(preStatement SQLStatement, op string) *ConditionGroup {
	group := newGroup(nil)
	group.buffer.WriteString(preStatement.Statement())
	group.buffer.WriteString(op)
	return group
}
