package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestCreateIndex(t *testing.T) {
	sbc := NewContext()
	sql := sbc.CreateIndex("PersonIndex").
		Unique().
		OnTable("t_users").
		Columns("LastName", "FirstName").
		Column("Age", true).
		ToSQL()

	checkSQLMatches(sql, "CREATE UNIQUE INDEX `PersonIndex` ON `t_users`(`LastName`, `FirstName`, `Age` DESC);", t)
}
