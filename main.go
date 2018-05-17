package main

import (
	"fmt"
	"github.com/go-sqlbuilder/gsb"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func main() {
	sql1 := gsb.Select("id", "username", "password").
		Distinct().
		From("t_users").
		Where(gsb.Group(gsb.Equal("username").And().EqualTo("password", "123456")).
			And().
			Group(gsb.LessThen("age").Or().In("nick_name", "yoojia", "yoojiachen"))).
		GetSQL()

	fmt.Println(sql1)

	sql2 := gsb.Select().
		From("t_users").
		OrderBy("username").ASC().
		Column("password").DESC().
		Limit(10).
		Offset(20).
		GetSQL()

	fmt.Println(sql2)

	sql3 := gsb.Insert("t_vehicles").
		Columns("id", "number", "color").
		Values(1, "粤BF49883", "GREEN").
		GetSQL()

	fmt.Println(sql3)
}
