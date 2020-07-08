package iterator_test

import (
	"fmt"
	"testing"

	"github.com/Lugo05/gds/iterator"
)

func SetUp() *iterator.Iterator {
	it := iterator.New()
	for i := 0; i < 10; i++ {
		it.Append(i)
	}
	return it
}

func TestIteratorBegin(t *testing.T) {
	it := SetUp()
	expected := 0
	got := it.Begin()
	if got != expected {
		t.Errorf("Error. Got %v %T. Expected %v %T\n", got, got, expected, expected)
	}
}

func TestIteratorNext(t *testing.T) {
	it := SetUp()
	expected := 6
	it.Begin()
	for i := 0; i < 5; i++ {
		fmt.Println(it.Next())
	}
	got := it.Next()
	if got != expected {
		t.Errorf("Error. Got %v %T. Expected %v %T\n", got, got, expected, expected)
	}
}

func TestIteartor(t *testing.T) {
	it := iterator.New()
	expected := make([]int, 10)

	for i := 0; i < 10; i++ {
		it.Append(i)
		expected[i] = i
	}
	i := 0
	for v := it.Begin(); v != nil && i < 10; v = it.Next() {
		if expected[i] != v {
			t.Errorf("Error. Expected %v %T. Got %v %T\n", expected[i], expected[i], v, v)
		}
		i++
	}
	if i != 10 {
		t.Errorf("Error. Expected iterate over 10 elements. Num. Iterations: %v", i)
	}
}
