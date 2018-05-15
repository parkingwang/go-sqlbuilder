package main

import (
	"fmt"
	"github.com/go-sqlbuilder/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func main() {
	fmt.Println(sql.Select("id", "username", "password").
		Distinct().
		From("t_users").
		Where().
		Equal("password").
		And().Equal("username").
		Or().GreaterEqualThen("age").
		SQL())
}
