package sqlx

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestDelete(t *testing.T) {
	sql := Delete("db.t_user").
		YesYesYesForceDelete().
		MakeSQL()
	checkSQLMatches(sql, "DELETE FROM `db.t_user`;", t)
}

func TestDeleteShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestDeleteShouldPanic should have panicked!")
		}
	}()

	Delete("db.t_user").MakeSQL()
}

func TestDeleteBuilder_Where(t *testing.T) {
	sql := Delete("t_users").
		Where(Equal("username").And().EqualTo("password", "123456")).
		MakeSQL()
	checkSQLMatches(sql, "DELETE FROM `t_users` WHERE `username` = ? AND `password` = '123456';", t)
}
