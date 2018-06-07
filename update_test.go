package sqlbuilder

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestUpdateShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestUpdateShouldPanic should have panicked!")
		}
	}()
	sb := NewContext()
	sb.Update("db.t_user").Columns("username").ToSQL()
}

func TestUpdate(t *testing.T) {
	sb := NewContext()
	sql := sb.Update("db.t_user").
		Columns("username").
		AddColumnValue("age", 18).
		YesImSureUpdateTable().
		Compile()
	checkSQLMatches(sql, "UPDATE `db.t_user` SET `username`=?, `age`=18", t)
}

func TestUpdateBuilder_Where(t *testing.T) {
	sb := NewContext()
	sql := sb.Update("db.t_user").
		Columns("username").
		Where(sb.GEt("age").
			Or().LtTo("height", 50)).
		ToSQL()
	checkSQLMatches(sql, "UPDATE `db.t_user` SET `username`=? WHERE `age` >= ? OR `height` < 50;", t)
}
