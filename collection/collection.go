package collection

//Abstraction of all Data Structures methods
type Collection interface {
	Add(...interface{}) Collection
	Remove(interface{}) Collection

	Clear() bool
	Contains(interface{}) bool
	IsEmpty() bool
	Size() int
	Equals(Collection) bool
}

//https://play.golang.org/p/OHbouik5Oe9
