package main

import (
	"fmt"
	"github.com/go-sqlbuilder/sqlx"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func main() {
	sql1 := sqlx.Select("id", "username", "password").
		Distinct().
		From("t_users").
		Where(sqlx.Group(sqlx.Equal("username").And().Equal("password")).
			And().
			Group(sqlx.LessThen("age").Or().In("pickname", "yoojia", "yoojiachen"))).
		MakeSQL()

	fmt.Println(sql1)

	sql2 := sqlx.Select().
		From("t_users").
		OrderBy("username").ASC().
		Column("password").DESC().
		Limit(10).
		Offset(20).
		MakeSQL()

	fmt.Println(sql2)

	sql3 := sqlx.Insert("t_vehicles").
		Columns("id", "number", "color").
		Values(1, "粤BF49883", "YELLOW").
		MakeSQL()

	fmt.Println(sql3)
}
