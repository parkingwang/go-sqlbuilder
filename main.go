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
		Where().
		Equal("password").
		And().EqualTo("username", "yoojia").
		Or().GreaterEqualThen("age").
		SQL()

	fmt.Println(sql1)

	sql2 := sqlx.Select().
		From("t_users").
		OrderBy("username").ASC().
		Column("password").DESC().
		Limit(10).
		Offset(20).
		SQL()

	fmt.Println(sql2)

	sql3 := sqlx.Insert("t_vehicles").
		Columns("id", "number", "color").
		Values(1, "粤BF49883", "YELLOW").
		SQL()

	fmt.Println(sql3)
}
