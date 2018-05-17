package sqlx

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestNewWhere(t *testing.T) {
	sql := newWhereTest(Equal("username")).Compile()
	checkSQLMatches(sql, " WHERE `username` = ?", t)
}

func TestNewWhereGroup(t *testing.T) {
	sql := newWhereTest(Group(Equal("username").And().Equal("password"))).Compile()
	checkSQLMatches(sql, " WHERE (`username` = ? AND `password` = ?)", t)
}

func TestNewWhereAnd(t *testing.T) {
	sql := newWhereTest(Equal("username").And().Equal("password")).
		Compile()
	checkSQLMatches(sql, " WHERE `username` = ? AND `password` = ?", t)
}

func TestNewWhereLimit(t *testing.T) {
	sql := newWhereTest(Equal("username").And().Equal("password")).Limit(10).
		Compile()
	checkSQLMatches(sql, " WHERE `username` = ? AND `password` = ? LIMIT 10", t)
}

func TestNewWhereOrderBy(t *testing.T) {
	sql := newWhereTest(Equal("username")).OrderBy("id").ASC().
		Compile()
	checkSQLMatches(sql, " WHERE `username` = ? ORDER BY `id` ASC", t)
}
