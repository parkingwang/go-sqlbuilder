# Go-SQLBuilder

以SQL语法顺序，使用流式API来创建SQL语句。

## Features

- 目前只支持MySQL语法；
- 未支持多表查询；

## Install

> go get -u https://github.com/parkingwang/go-sqlbuilder

## Usage

```go

import "github.com/go-sqlbuilder/sqlx"

func main() {
	sql1 := sqlx.Select("id", "username", "password").
        Distinct().
        From("t_users").
        Where(sqlx.Group(sqlx.Equal("username").And().EqualTo("password", "123456")).
            And().
            Group(sqlx.LessThen("age").Or().In("nick_name", "yoojia", "yoojiachen"))).
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
        Values(1, "粤BF49883", "GREEN").
        MakeSQL()

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