package gsb

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestGroup(t *testing.T) {
	ctx := NewContext()
	sql := Group(ctx.Eq("username").And().NEq("age")).
		And().
		Group(ctx.Eq("nick_name").Or().NEq("height")).
		Compile()
	checkSQLMatches(sql, "(`username` = ? AND `age` <> ?) AND (`nick_name` = ? OR `height` <> ?)", t)
}

func TestGroup1(t *testing.T) {
	ctx := NewContext()
	sql := Group(ctx.Eq("username").And().NEq("age")).
		And().
		Group(ctx.Eq("nick_name").Or().NEq("height")).
		Or().
		Group(ctx.Eq("age").And().Eq("weight")).
		Compile()
	checkSQLMatches(sql, "(`username` = ? AND `age` <> ?) AND (`nick_name` = ? OR `height` <> ?) OR (`age` = ? AND `weight` = ?)", t)
}
