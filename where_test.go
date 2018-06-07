package sqlbuilder

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestNewWhere(t *testing.T) {
	ctx := NewContext()
	sql := newWhereTest(ctx, ctx.Eq("username")).Compile()
	checkSQLMatches(sql, " WHERE `username` = ?", t)
}

func TestNewWhereGroup(t *testing.T) {
	ctx := NewContext()
	sql := newWhereTest(ctx, Group(ctx.Eq("username").And().Eq("password"))).Compile()
	checkSQLMatches(sql, " WHERE (`username` = ? AND `password` = ?)", t)
}

func TestNewWhereAnd(t *testing.T) {
	ctx := NewContext()
	sql := newWhereTest(ctx, ctx.Eq("username").And().Eq("password")).
		Compile()
	checkSQLMatches(sql, " WHERE `username` = ? AND `password` = ?", t)
}

func TestNewWhereLimit(t *testing.T) {
	ctx := NewContext()
	sql := newWhereTest(ctx, ctx.Eq("username").And().Eq("password")).Limit(10).
		Compile()
	checkSQLMatches(sql, " WHERE `username` = ? AND `password` = ? LIMIT 10", t)
}

func TestNewWhereOrderBy(t *testing.T) {
	ctx := NewContext()
	sql := newWhereTest(ctx, ctx.Eq("username")).OrderBy("id").ASC().Compile()
	checkSQLMatches(sql, " WHERE `username` = ? ORDER BY `id` ASC", t)
}
