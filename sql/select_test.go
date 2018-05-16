package sql

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestSelectAll(t *testing.T) {
	sql := Select().From("t_users").
		SQL()
	checkSQLMatches(sql, "SELECT * FROM `t_users`;", t)
}

func TestSelect(t *testing.T) {
	sql := Select("id", "username").
		From("t_users").
		SQL()
	checkSQLMatches(sql, "SELECT `id`,`username` FROM `t_users`;", t)
}

func TestSelectWhere(t *testing.T) {
	sql := Select("id", "username").
		From("t_users").
		Where().Equal("password").
		Or().EqualTo("password", "*").
		SQL()
	checkSQLMatches(sql, "SELECT `id`,`username` FROM `t_users` WHERE `password` = ? OR `password` = '*';", t)
}

func TestSelectWhereOrder(t *testing.T) {
	sql := Select("id", "username").
		From("t_users").
		Where().Equal("password").
		OrderBy("id").ASC().
		SQL()
	checkSQLMatches(sql, "SELECT `id`,`username` FROM `t_users` WHERE `password` = ? ORDER BY `id` ASC;", t)
}

func TestSelectWhereLimit(t *testing.T) {
	sql := Select("id", "username").
		From("t_users").
		Where().Equal("password").
		Limit(10).Offset(200).
		SQL()
	checkSQLMatches(sql, "SELECT `id`,`username` FROM `t_users` WHERE `password` = ? LIMIT 10 OFFSET 200;", t)
}
