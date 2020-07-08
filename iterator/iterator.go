package iterator

type Iterator struct {
	data  []interface{}
	index int
}

func New() *Iterator {
	iterator := &Iterator{}
	iterator.data = []interface{}{}
	iterator.index = 0
	return iterator
}

func (it *Iterator) Append(val interface{}) {
	it.data = append(it.data, val)
}

func (it *Iterator) Begin() interface{} {
	it.index = 0
	if len(it.data) == 0 {
		return nil
	}
	it.index = 1
	return it.data[0]
}

func (it *Iterator) Next() interface{} {
	if it.index < len(it.data) {
		val := it.data[it.index]
		it.index++
		return val
	}
	it.index = 0
	return nil
}
