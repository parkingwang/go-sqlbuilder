package main

import (
	"fmt"
	"github.com/go-sqlbuilder/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func main() {
	sql1 := sql.Select("id", "username", "password").
		Distinct().
		From("t_users").
		Where().
		Equal("password").
		And().EqualTo("username", "yoojia").
		Or().GreaterEqualThen("age").
		SQL()

	fmt.Println(sql1)

	sql2 := sql.Select().
		From("t_users").
		OrderBy("username").ASC().
		Column("password").DESC().
		Limit(10).
		Offset(20)

	fmt.Println(sql2)
}
