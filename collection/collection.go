package collection

type Collection interface {
	Add(...interface{}) Collection
	Remove(interface{}) Collection

	Clear() bool
	Contains(interface{}) bool
	IsEmpty() bool
	Size() int
	ToArray() []interface{}
	Next() interface{}
	Equals(Collection) bool
}

//https://play.golang.org/p/OHbouik5Oe9
