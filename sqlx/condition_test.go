package sqlx

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestCondition_And(t *testing.T) {
	sql := newCondition().
		Equal("username").
		And().NotEqual("age").
		Statement()
	checkSQLMatches(sql, "`username` = ? AND `age` <> ?", t)
}

func TestCondition_Between(t *testing.T) {
	sql := newCondition().
		Between("age", 18, 32).
		Statement()
	checkSQLMatches(sql, "`age` BETWEEN 18 AND 32", t)
}

func TestCondition_In(t *testing.T) {
	sql := newCondition().
		In("name", "yoojia", "yoojiachen", "yoojiachen@gmail.com").
		Statement()
	checkSQLMatches(sql, "`name` IN ('yoojia', 'yoojiachen', 'yoojiachen@gmail.com')", t)
}
