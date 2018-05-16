# Go-SQLBuilder

`Go-SQLBuilder`是一个提供系列灵活的、与SQL语法一致的SQL语句构建工具函数库。

## 目标

`GoSQLBuilder`的目标是提供一个简洁灵活的工具集，可以与SQL语法顺序一致地使用，并生成SQL语句，供database/sql原生数据库操作函数使用。
最显著的是，生成的SQL语句，最大限度使用占位符号，并使用`sql.DB.Query(args)`和`sql.DB.Exec(args)`来设值避免SQL注入等问题。

## 支持特性

- 已支持MySQL基本Select/Update/Insert/Delete/Where等语法；
- 目前只支持MySQL语法；
- 未支持多表查询；

## Install

> go get -u https://github.com/parkingwang/go-sqlbuilder

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-sqlbuilder/sqlx"
)

func main() {
	sql1 := sqlx.Select("id", "username", "password").
		Distinct().
		From("t_users").
		Where(sqlx.Group(sqlx.Equal("username").And().EqualTo("password", "123456")).
			And().
			Group(sqlx.LessThen("age").Or().In("nick_name", "yoojia", "yoojiachen"))).
		GetSQL()

	fmt.Println(sql1)

	sql2 := sqlx.Select().
		From("t_users").
		OrderBy("username").ASC().
		Column("password").DESC().
		Limit(10).
		Offset(20).
		GetSQL()

	fmt.Println(sql2)

	sql3 := sqlx.Insert("t_vehicles").
		Columns("id", "number", "color").
		Values(1, "粤BF49883", "GREEN").
		GetSQL()

	fmt.Println(sql3)
}
```

输出结果如下：

```sql
SELECT DISTINCT `id`,`username`,`password` FROM `t_users`
    WHERE (`username` = ? AND `password` = '123456') AND (`age` < ? OR `nick_name` IN ('yoojia','yoojiachen'));

SELECT * FROM `t_users` ORDER BY `username` ASC, `password` DESC LIMIT 10 OFFSET 20;

INSERT INTO `t_vehicles`(`id`,`number`,`color`) VALUES (1,'粤BF49883','GREEN');
```

## License

    Copyright [2018] Xi'an iRain IoT Technology Service CO., Ltd.
    Copyright [2018] Yoojia Chen chenyongjia@parkingwang, yoojiachen@gmail.com

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.