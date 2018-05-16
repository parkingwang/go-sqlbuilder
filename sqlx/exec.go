package sqlx

import "database/sql"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type Executor struct {
	sql string
	db  *sql.DB
}

func newExecute(sql string, db *sql.DB) *Executor {
	return &Executor{
		sql: sql,
		db:  db,
	}
}

func (eb *Executor) Exec(args ...interface{}) (sql.Result, error) {
	stmt, pErr := eb.db.Prepare(eb.sql)
	if nil != pErr {
		return nil, pErr
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func (eb *Executor) Query(args ...interface{}) (*sql.Rows, error) {
	stmt, pErr := eb.db.Prepare(eb.sql)
	if nil != pErr {
		return nil, pErr
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

func (eb *Executor) Exists(args ...interface{}) (bool, error) {
	stmt, pErr := eb.db.Prepare(eb.sql)
	if nil != pErr {
		return false, pErr
	}
	defer stmt.Close()

	rs, qErr := stmt.Query(args...)
	if nil != qErr {
		return false, qErr
	}
	defer rs.Close()

	return rs.Next(), nil
}
