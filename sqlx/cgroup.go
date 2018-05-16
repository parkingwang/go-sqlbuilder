package sqlx

import "bytes"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type ConditionGroup struct {
	buffer *bytes.Buffer // 生成语句的缓存
	cond   Statement     // 由外部指定的，用于使用括号包装起来的条件语句
}

func Group(cond Statement) *ConditionGroup {
	return newGroup(cond)
}

func (g *ConditionGroup) Group(cond Statement) *ConditionGroup {
	g.cond = cond
	return g
}

func (g *ConditionGroup) And() *ConditionGroup {
	return newGroupWith(g.SQL(), " AND ")
}

func (g *ConditionGroup) Or() *ConditionGroup {
	return newGroupWith(g.SQL(), " OR ")
}

func (g *ConditionGroup) SQL() string {
	// 写入本组SQL到缓存
	// 如果预存有SQL，则自动拼接
	g.buffer.WriteString("(" + g.cond.SQL() + ")")
	return g.buffer.String()
}

//

func newGroup(cond Statement) *ConditionGroup {
	return &ConditionGroup{
		buffer: new(bytes.Buffer),
		cond:   cond,
	}
}

func newGroupWith(existsSQL string, op string) *ConditionGroup {
	group := newGroup(nil)
	group.buffer.WriteString(existsSQL)
	group.buffer.WriteString(op)
	return group
}
