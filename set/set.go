package set

import (
	"fmt"

	. "github.com/Lugo05/gds/collection"
)

type Set struct {
	data     map[interface{}]int
	hasNext  bool
	iterator []interface{}
}

func (set *Set) Add(elems ...interface{}) Collection {
	for _, elem := range elems {
		if !set.Contains(elem) {
			set.data[elem] = 0
		}
	}
	return set
}

func (set *Set) Remove(elem interface{}) Collection {
	if set.Contains(elem) {
		delete(set.data, elem)
	}
	return set
}

func (set *Set) Clear() bool {
	set.data = make(map[interface{}]int)
	return true
}

func (set *Set) Contains(elem interface{}) bool {
	_, ok := set.data[elem]
	return ok
}

func (set *Set) IsEmpty() bool {
	return set.Size() == 0
}

func (set *Set) Size() int {
	return len(set.data)
}

func (set *Set) ToArray() []interface{} {
	card := set.Size()
	array := make([]interface{}, card)
	i := 0
	for value, _ := range set.data {
		array[i] = value
		i++
	}
	return array
}

func (set *Set) Next() interface{} {
	if len(set.iterator) == 0 {
		if set.hasNext {
			set.iterator = set.ToArray()
		} else {
			set.hasNext = true
			return nil
		}
	}
	size := len(set.iterator)
	if size == 0 {
		return nil
	}
	val := set.iterator[0]
	set.iterator = set.iterator[1:size]
	if size-1 == 0 {
		set.hasNext = false
	}
	return val
}

func (set *Set) Equals(collection Collection) bool {
	setCmp, correct := collection.(*Set)
	if !correct {
		return false
	}
	for v, _ := range set.data {
		if !setCmp.Contains(v) {
			return false
		}
	}
	return true
}

func (set *Set) String() string {
	return fmt.Sprintf("(util.Set) %v", set.ToArray())
}
