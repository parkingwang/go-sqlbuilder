package sqlbuilder

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//
//

const SQLMaxColumns = 255
const SQLDefaultColumns = 64

type List struct {
	items  []interface{}
	cursor int
}

func newItems() *List {
	return &List{
		items:  make([]interface{}, SQLDefaultColumns, SQLMaxColumns),
		cursor: 0,
	}
}

func (slf *List) Count() int {
	return slf.cursor
}

func (slf *List) Range(it func(index int, val interface{}) bool) {
	for i, v := range slf.AvailableItems() {
		if !it(i, v) {
			break
		}
	}
}

func (slf *List) Add(val interface{}) {
	for slf.cursor >= len(slf.items) {
		slf.items = append(slf.items, make([]interface{}, 8)...)
	}
	slf.items[slf.cursor] = val
	slf.cursor++
}

func (slf *List) SetAt(val interface{}, index int) {
	slf.items[index] = val
}

func (slf *List) AvailableItems() []interface{} {
	return slf.items[:slf.cursor]
}

func (slf *List) AvailableStrItems() []string {
	return MapAny(slf.AvailableItems(), func(val interface{}) string {
		return val.(string)
	})
}
