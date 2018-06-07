package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestSelectAll(t *testing.T) {
	sb := NewContext()
	sql := sb.Select("*").From("t_users").
		ToSQL()
	checkSQLMatches(sql, "SELECT * FROM `t_users`;", t)
}

func TestSelect(t *testing.T) {
	sb := NewContext()
	sql := sb.Select("id", "username").
		From("t_users").
		ToSQL()
	checkSQLMatches(sql, "SELECT `id`, `username` FROM `t_users`;", t)
}

func TestSelectWhere(t *testing.T) {
	sb := NewContext()
	sql := sb.Select("id", "username").
		From("t_users").
		Where(sb.Eq("password").
			Or().EqTo("password", "*")).
		ToSQL()
	checkSQLMatches(sql, "SELECT `id`, `username` FROM `t_users` WHERE `password` = ? OR `password` = '*';", t)
}

func TestSelectWhereOrder(t *testing.T) {
	sb := NewContext()
	sql := sb.Select("id", "username").
		From("t_users").
		Where(sb.Eq("password")).
		OrderBy("id").ASC().
		ToSQL()
	checkSQLMatches(sql, "SELECT `id`, `username` FROM `t_users` WHERE `password` = ? ORDER BY `id` ASC;", t)
}

func TestSelectWhereLimit(t *testing.T) {
	sb := NewContext()
	sql := sb.Select("id", "username").
		From("t_users").
		Where(sb.Eq("password")).
		Limit(10).Offset(200).
		ToSQL()
	checkSQLMatches(sql, "SELECT `id`, `username` FROM `t_users` WHERE `password` = ? LIMIT 10 OFFSET 200;", t)
}

func TestSelectWhereInnerSelect(t *testing.T) {
	sb := NewContext()
	sql := sb.Select("id", "username").
		FromSelect(sb.Select("*").From("t_users_bak").Where(sb.NEq("name"))).
		Where(sb.Eq("password")).
		Limit(10).Offset(200).
		ToSQL()
	checkSQLMatches(sql, "SELECT `id`, `username` FROM (SELECT * FROM `t_users_bak` WHERE `name` <> ?) WHERE `password` = ? LIMIT 10 OFFSET 200;", t)
}
