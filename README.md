# Go-SQLBuilder

以SQL语法顺序，使用流式API来创建SQL语句。

## Features

- 目前只支持MySQL语法；

## Install

> go get -u https://github.com/parkingwang/go-sqlbuilder

## Usage

```go
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

sql3 := sql.InsertInto("t_vehicles").
    Columns("id", "number", "color").
    Values(1, "粤BF49883", "YELLOW").
    SQL()

fmt.Println(sql3)
```