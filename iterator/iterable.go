package iterator

//Abstraction of all data structures which can be iterabled
type Iterable interface {
	GetIterator() *Iterator
	ToArray() []interface{}
}
