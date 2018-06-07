package sqlbuilder

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestInsertInto(t *testing.T) {
	sb := NewContext()
	sql := sb.Insert("t_users").
		Columns("username", "password").
		SetValueOfColumn("password", 123).
		ToSQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`, `password`) VALUES (?, 123);", t)
}

func TestInsertInto1(t *testing.T) {
	sb := NewContext()
	sql := sb.Insert("t_users").
		Columns("username", "password").
		Values("yoojia", "1234").
		ToSQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`, `password`) VALUES ('yoojia', '1234');", t)
}

func TestInsertIntoValued(t *testing.T) {
	sb := NewContext()
	sql := sb.Insert("t_users").
		Columns("username", "password").
		Values("yoojia", "123456").
		ToSQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`, `password`) VALUES ('yoojia', '123456');", t)
}
