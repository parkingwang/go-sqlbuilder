package gsb

import "testing"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestDropIndex(t *testing.T) {
	sbc := NewContext()
	sql := sbc.DropIndex("idx_Uid").
		OnTable("t_username").
		ToSQL()
	checkSQLMatches(sql, "ALTER TABLE `t_username` DROP INDEX `idx_Uid`;", t)
}
