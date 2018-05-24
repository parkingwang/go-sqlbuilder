package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestCreateIndex(t *testing.T) {
	sql := CreateIndex("PersonIndex").
		Unique().
		OnTable("t_users").
		Columns("LastName", "FirstName").
		Column("Age", true).
		GetSQL()

	checkSQLMatches(sql, "CREATE UNIQUE INDEX `PersonIndex` ON `t_users`(`LastName`, `FirstName`, `Age` DESC);", t)
}
