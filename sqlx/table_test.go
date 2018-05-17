package sqlx

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestCreateTable(t *testing.T) {
	sql := CreateTable("t_users").
		Column("id").Int(20).NotNull().PrimaryKey().AutoIncrement().
		Column("username").VarChar(255).NotNull().Unique().
		Column("password").VarChar(255).NotNull().
		Column("age").Int(2).Default0().
		Column("register_time").Date().DefaultGetDate().
		GetSQL()
	checkSQLMatches(sql, "CREATE TABLE IF NOT EXISTS `t_users`(`id` INT(20) NOT NULL AUTO_INCREMENT, "+
		"`username` VARCHAR(255) NOT NULL, `password` VARCHAR(255) NOT NULL, `age` INT(2) DEFAULT 0, "+
		"`register_time` DATE DEFAULT GETDATE(), PRIMARY KEY(`id`), UNIQUE(`username`)) "+
		"DEFAULT CHARSET=utf8 AUTO_INCREMENT=0;", t)
}
