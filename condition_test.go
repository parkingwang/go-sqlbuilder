package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestCondition_And(t *testing.T) {
	sqc := NewContext()
	sql := newConditionBuilder(sqc).
		Eq("username").
		And().NEq("age").
		Compile()
	checkSQLMatches(sql, "`username` = ? AND `age` <> ?", t)
}

func TestCondition_Between(t *testing.T) {
	sqc := NewContext()
	sql := newConditionBuilder(sqc).
		Between("age", 18, 32).
		Compile()
	checkSQLMatches(sql, "`age` BETWEEN 18 AND 32", t)
}

func TestCondition_In(t *testing.T) {
	sqc := NewContext()
	sql := newConditionBuilder(sqc).
		In("name", "yoojia", "yoojiachen", "yoojiachen@gmail.com").
		Compile()
	checkSQLMatches(sql, "`name` IN ('yoojia', 'yoojiachen', 'yoojiachen@gmail.com')", t)
}
