package gsb

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

const SQLPlaceHolder = "?"

type SQLStatement interface {
	Compile() string
}

type SQLGenerator interface {
	ToSQL() string
}

type Executable interface {
	Execute() *Executor
}
