package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestDropIndex(t *testing.T) {
	sql := DropIndex("idx_Uid").
		OnTable("t_username").
		GetSQL()
	checkSQLMatches(sql, "ALTER TABLE `t_username` DROP INDEX `idx_Uid`;", t)
}
