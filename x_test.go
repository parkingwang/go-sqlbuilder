package gsb

import (
	"fmt"
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func checkSQLMatches(actual string, except string, t *testing.T) {
	fmt.Println("Actual: " + actual)
	if actual != except {
		t.Error("Output sql not match, was:  " + except)
	}
}

func newWhereTest(ctx *SQLContext, conditions SQLStatement) *WhereBuilder {
	return newWhereBuilder(ctx, "", conditions)
}
