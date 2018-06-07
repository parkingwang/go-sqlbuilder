package sqlbuilder

import (
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestCreateTable(t *testing.T) {
	sbc := NewContext()
	sql := sbc.CreateTable("t_users").
		Column("id").Int(20).NotNull().PrimaryKey().AutoIncrement().
		Column("username").VarChar(255).NotNull().Unique().
		Column("password").VarChar(255).NotNull().
		Column("age").Int(2).Default0().
		Column("register_time").Date().DefaultNow().
		ToSQL()
	checkSQLMatches(sql, "CREATE TABLE IF NOT EXISTS `t_users`("+
		"`id` INT(20) NOT NULL AUTO_INCREMENT, "+
		"`username` VARCHAR(255) NOT NULL, "+
		"`password` VARCHAR(255) NOT NULL, "+
		"`age` INT(2) DEFAULT 0, "+
		"`register_time` DATE DEFAULT NOW(), "+
		"PRIMARY KEY(`id`), "+
		"UNIQUE (`username`)"+
		") DEFAULT CHARSET=utf8 AUTO_INCREMENT=0;", t)
}

func TestTableBuilder_ForeignKey(t *testing.T) {
	sbc := NewContext()
	sql := sbc.CreateTable("t_user").
		Column("id").Int(12).NotNull().PrimaryKey().AutoIncrement().
		Column("pid").Int(12).ForeignKey("t_profile", "prof_id").
		ToSQL()
	checkSQLMatches(sql, "CREATE TABLE IF NOT EXISTS `t_user`("+
		"`id` INT(12) NOT NULL AUTO_INCREMENT, "+
		"`pid` INT(12), PRIMARY KEY(`id`), "+
		"FOREIGN KEY (`pid`) REFERENCES `t_profile`(`prof_id`)"+
		") DEFAULT CHARSET=utf8 AUTO_INCREMENT=0;", t)
}

func TestTableBuilder_ForeignKeyNamed(t *testing.T) {
	sbc := NewContext()
	sql := sbc.CreateTable("t_user").
		Column("id").Int(12).NotNull().PrimaryKey().AutoIncrement().
		Column("pid").Int(12).ForeignKeyNamed("FK_PID", "t_profile", "prof_id").
		ToSQL()
	checkSQLMatches(sql, "CREATE TABLE IF NOT EXISTS `t_user`("+
		"`id` INT(12) NOT NULL AUTO_INCREMENT, "+
		"`pid` INT(12), PRIMARY KEY(`id`), "+
		"CONSTRAINT `FK_PID` FOREIGN KEY (`pid`) REFERENCES `t_profile`(`prof_id`)"+
		") DEFAULT CHARSET=utf8 AUTO_INCREMENT=0;", t)
}

func TestTableBuilder_UniqueNamed(t *testing.T) {
	sbc := NewContext()
	sql := sbc.CreateTable("t_user").
		Column("id").Int(12).NotNull().PrimaryKey().AutoIncrement().
		Column("pid").Int(12).UniqueNamed("uc_Id_P").
		Column("pid_bak").Int(12).UniqueNamed("uc_Id_P").
		ToSQL()
	checkSQLMatches(sql, "CREATE TABLE IF NOT EXISTS `t_user`("+
		"`id` INT(12) NOT NULL AUTO_INCREMENT, "+
		"`pid` INT(12), "+
		"`pid_bak` INT(12), "+
		"PRIMARY KEY(`id`), "+
		"CONSTRAINT `uc_Id_P` UNIQUE (`pid`, `pid_bak`)"+
		") DEFAULT CHARSET=utf8 AUTO_INCREMENT=0;", t)
}
