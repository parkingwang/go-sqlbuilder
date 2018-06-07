# SQL Builder in GoLang

`Go-SQLBuilder`是一个用于创建SQL语句的工具函数库，提供一系列灵活的、与原生SQL语法一致的链式函数。

## Projects

`Go-SQLBuilder`项目归属于艾润物联公司，项目地址：

开发项目由开发者Fork并持续开发：

- 开发项目：[https://github.com/yoojia/go-sqlbuilder](https://github.com/yoojia/go-sqlbuilder)
- 开发项目：[https://gitee.com/yoojia/go-sqlbuilder](https://gitee.com/yoojia/go-sqlbuilder)

当版本稳定时，通过PR合并到发布项目：
- 发布项目：[https://github.com/parkingwang/go-sqlbuilder](https://github.com/parkingwang/go-sqlbuilder)
- 发布项目：[https://gitee.com/iRainIoT/go-sqlbuilder](https://gitee.com/iRainIoT/go-sqlbuilder)

## Goals

`Go-SQLBuilder`的目标是提供一个简洁易用的函数工具集，它可以与SQL原生语法一致地使用，来用生成SQL语句，提供组`database/sql`包的原生数据库操作函数使用。
最显著的是，Go-SQLBuilder生成的SQL语句，最大程序上使用占位符号`?`来替代数值位，并建议使用`sql.DB.Query(args)`和`sql.DB.Exec(args)`来设值避免SQL注入等问题。

## Features

- 已支持MySQL基本Select/Update/Insert/Delete/Where等语法；
- 目前只支持MySQL语法；
- 未支持多表查询；

## Install

> go get -u github.com/parkingwang/go-sqlbuilder

## Usage

### Base Usage

```go

sql1 := gsb.Select().
    From("t_users").
    OrderBy("username").ASC().
    Column("password").DESC().
    Limit(10).
    Offset(20).
    ToSQL()

fmt.Println(sql1)
```

Output:

```sql
SELECT * FROM `t_users` ORDER BY `username` ASC, `password` DESC LIMIT 10 OFFSET 20;
```

### Advance Usage

```go

sql1 := gsb.Select("id", "username", "password").
    Distinct().
    From("t_users").
    Where(gsb.Group(gsb.Equal("username").And().EqualTo("password", "123456")).
        And().
        Group(gsb.LessThen("age").Or().In("nick_name", "yoojia", "yoojiachen"))).
    ToSQL()

fmt.Println(sql1)

sql2 := gsb.Insert("t_vehicles").
    Columns("id", "number", "color").
    Values(1, "粤BF49883", "GREEN").
    ToSQL()

fmt.Println(sql2)
```

Output:

```sql
SELECT DISTINCT `id`,`username`,`password` FROM `t_users`
      WHERE (`username` = ? AND `password` = '123456') AND (`age` < ? OR `nick_name` IN ('yoojia','yoojiachen'));

INSERT INTO `t_vehicles`(`id`, `number`, `color`) VALUES (1, '粤BF49883', 'GREEN');
```

### Inner Select

```go
sql := Select("id", "username").
		FromSelect(Select().From("t_users_bak").Where(NotEqual("name"))).
		Where(Equal("password")).
		Limit(10).Offset(200).
		ToSQL()

fmt.Println(sql)
```

Output:
```sql
SELECT `id`, `username` FROM (SELECT * FROM `t_users_bak` WHERE `name` <> ?) WHERE `password` = ? LIMIT 10 OFFSET 200;
```

## License

    Copyright 2018 西安艾润物联网技术服务有限责任公司
    Copyright 2018 陈哈哈(chenyongjia@parkingwang, yoojiachen@gmail.com)

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.