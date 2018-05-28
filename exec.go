package gsb

import (
	"database/sql"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type SQLPrepare interface {
	Prepare(query string) (*sql.Stmt, error)
}

type Executor struct {
	query         string
	DBPrepare     SQLPrepare
	logger        func(sql string, args []interface{}) // 日志输出接口
	slowThreshold time.Duration                        // 慢执行阀值
}

func newExecute(query string, prepare SQLPrepare) *Executor {
	return &Executor{
		query:         query,
		DBPrepare:     prepare,
		slowThreshold: 0,
	}
}

func (slf *Executor) SetSlowExecThreshold(t time.Duration) *Executor {
	slf.slowThreshold = t
	return slf
}

func (slf *Executor) Logger(logger func(string, []interface{})) *Executor {
	slf.logger = logger
	return slf
}

func (slf *Executor) Exec(args ...interface{}) (sql.Result, error) {
	start := time.Now()
	defer slf.checkSlowExecution(start)

	if nil != slf.logger {
		slf.logger(slf.query, args)
	}

	stmt, pErr := slf.DBPrepare.Prepare(slf.query)
	if nil != pErr {
		return nil, pErr
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func (slf *Executor) Query(args ...interface{}) (*sql.Rows, error) {
	start := time.Now()
	defer slf.checkSlowExecution(start)

	if nil != slf.logger {
		slf.logger(slf.query, args)
	}

	stmt, pErr := slf.DBPrepare.Prepare(slf.query)
	if nil != pErr {
		return nil, pErr
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

func (slf *Executor) Exists(args ...interface{}) (bool, error) {
	start := time.Now()
	defer slf.checkSlowExecution(start)

	if nil != slf.logger {
		slf.logger(slf.query, args)
	}

	stmt, pErr := slf.DBPrepare.Prepare(slf.query)
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

func (slf *Executor) checkSlowExecution(start time.Time) {
	if slf.slowThreshold <= 0 {
		return
	}
	takes := time.Now().Sub(start)
	if takes > slf.slowThreshold {
		slf.logger(slf.query, []interface{}{"takes=" + takes.String()})
	}
}
