package main

import (
	"fmt"
	"github.com/go-sql/sql"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func main() {
	fmt.Println( sql.Select("id", "username", "password").
		From("t_users").
		Where().
		Equal("password").
		And().Equal("username").
		Or().GreaterEqualThen("age").
		SQL())
}
