package sql

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestInsertInto(t *testing.T) {
	sql := InsertInto("t_users").
		Columns("username", "password").
		SQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`,`password`) VALUES (?,?);", t)
}

func TestInsertIntoValued(t *testing.T) {
	sql := InsertInto("t_users").
		Columns("username", "password").
		Values("yoojia", "123456").
		SQL()
	checkSQLMatches(sql, "INSERT INTO `t_users`(`username`,`password`) VALUES ('yoojia','123456');", t)
}
