package sqlx

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestGroup(t *testing.T) {
	sql := Group(Equal("username").And().NotEqual("age")).
		And().
		Group(Equal("nick_name").Or().NotEqual("height")).
		SQL()
	checkSQLMatches(sql, "(`username` = ? AND `age` <> ?) AND (`nick_name` = ? OR `height` <> ?)", t)
}

func TestGroup1(t *testing.T) {
	sql := Group(Equal("username").And().NotEqual("age")).
		And().
		Group(Equal("nick_name").Or().NotEqual("height")).
		Or().
		Group(Equal("age").And().Equal("weight")).
		SQL()
	checkSQLMatches(sql, "(`username` = ? AND `age` <> ?) AND (`nick_name` = ? OR `height` <> ?) OR (`age` = ? AND `weight` = ?)", t)
}
