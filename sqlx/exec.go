package sqlx

import "database/sql"

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type Executor struct {
	sql    string
	db     *sql.DB
	logger func(sql string, args []interface{})
}

func newExecute(sql string, db *sql.DB) *Executor {
	return &Executor{
		sql: sql,
		db:  db,
	}
}

func (slf *Executor) Logger(logger func(string, []interface{})) *Executor {
	slf.logger = logger
	return slf
}

func (slf *Executor) Exec(args ...interface{}) (sql.Result, error) {
	if nil != slf.logger {
		slf.logger(slf.sql, args)
	}

	stmt, pErr := slf.db.Prepare(slf.sql)
	if nil != pErr {
		return nil, pErr
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func (slf *Executor) Query(args ...interface{}) (*sql.Rows, error) {
	if nil != slf.logger {
		slf.logger(slf.sql, args)
	}

	stmt, pErr := slf.db.Prepare(slf.sql)
	if nil != pErr {
		return nil, pErr
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

func (slf *Executor) Exists(args ...interface{}) (bool, error) {
	if nil != slf.logger {
		slf.logger(slf.sql, args)
	}

	stmt, pErr := slf.db.Prepare(slf.sql)
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
