package gsb

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestDelete(t *testing.T) {
	sb := New()
	sql := sb.Delete("db.t_user").
		YesImSureDeleteTable().
		ToSQL()
	checkSQLMatches(sql, "DELETE FROM `db.t_user`;", t)
}

func TestDeleteShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestDeleteShouldPanic should have panicked!")
		}
	}()
	sb := New()
	sb.Delete("db.t_user").ToSQL()
}

func TestDeleteBuilder_Where(t *testing.T) {
	sb := New()
	sql := sb.Delete("t_users").
		Where(Equal("username").And().EqualTo("password", "123456")).
		ToSQL()
	checkSQLMatches(sql, "DELETE FROM `t_users` WHERE `username` = ? AND `password` = '123456';", t)
}
