package document

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	. "github.com/Lugo05/gds/collection"
	"github.com/Lugo05/gds/iterator"
	. "github.com/Lugo05/gds/types"
)

/*
Data structure similar to Map.
We avoid assertion every level.
Manipulation is easier.
*/
type Document struct {
	data map[string]interface{}
}

/*We define which types are assignable. It helps to encapsulate different variations of types in just one.*/
var allowedTypes = []Type{
	Int, UInt, Float, Complex, String, Bool,
	ArrayBool, ArrayInt, ArrayUInt, ArrayFloat,
	ArrayComplex, ArrayString, ArrayInterface, &Document{},
}

//
func (document *Document) String() string {

	stringValue := "(*document.Document) {\n"
	for k, v := range document.data {
		t := document.Type(k)
		if strings.HasSuffix(t, "Document") {
			aux := document.Child(k).String()
			aux = strings.ReplaceAll(aux, "\n", "\n\t")
			stringValue += fmt.Sprintf("\t'%v': %v\n", k, aux)
		} else {
			stringValue += fmt.Sprintf("\t'%v': (%T) %v\n", k, v, v)
		}
	}
	stringValue += "}"
	return stringValue
}

func (document *Document) Add(data ...interface{}) Collection {
	if len(data)%2 != 0 {
		panic("Method Add in document must have an even number of parameters.")
	}

	for i := 0; i < len(data); i += 2 {
		key := fmt.Sprintf("%v", data[0])
		val := data[1]
		found := false
		for _, allTyp := range allowedTypes {
			if value, valid := allTyp.IsAssignable(val); valid {
				document.data[key] = value
				found = true
				break
			}
		}
		if !found {
			document.data[key] = val
		}
	}
	return document
}

func (document *Document) Remove(key interface{}) Collection {
	keyS, valid := key.(string)
	if !valid {
		panic("Method Remove must receive a string key")
	}
	delete(document.data, keyS)
	return document
}

func (document *Document) Clear() bool {
	document.data = make(map[string]interface{})
	return true
}

func (document *Document) Contains(k interface{}) bool {
	key, valid := k.(string)
	if !valid {
		panic("Method Contains must receive a string key")
	}
	_, ok := document.data[key]
	return ok
}

func (document *Document) IsEmpty() bool {
	return document.Size() == 0
}

func (document *Document) Size() int {
	return len(document.data)
}

func (document *Document) Equals(c Collection) bool {
	docComp, correct := c.(*Document)
	if !correct {
		return false
	}
	for k, v := range document.data {
		if !docComp.Contains(k) {
			return false
		} else if docComp.Type(k) != document.Type(k) {
			return false
		} else if document.Type(k) == "*document.Document" {
			aux := document.Child(k).Equals(docComp.Child(k))
			if !aux {
				return aux
			}
		} else if !reflect.DeepEqual(v, docComp.Value(k)) {
			return false
		}
	}
	return true
}

func (document *Document) Change(key string, newValue interface{}) *Document {
	document.Remove(key)
	document.Add(key, newValue)
	return document
}

func (document *Document) Type(key string) string {
	if document.Contains(key) {
		t := fmt.Sprintf("%T", document.data[key])
		if strings.HasSuffix(t, "Document") {
			return "*document.Document"
		}
		return t
	}
	return ""
}

func (document *Document) Child(key string) *Document {
	if document.Contains(key) {
		return document.data[key].(*Document)
	}
	return nil
}

func (document *Document) String_(key string) string {
	if document.Contains(key) {
		return document.data[key].(string)
	}
	return ""
}

func (document *Document) Int(key string) int64 {
	if document.Contains(key) {
		return document.data[key].(int64)
	}
	return 0
}

func (document *Document) UInt(key string) uint64 {
	if document.Contains(key) {
		return document.data[key].(uint64)
	}
	return 0
}

func (document *Document) Float(key string) float64 {
	if document.Contains(key) {
		return document.data[key].(float64)
	}
	return 0.0
}

func (document *Document) Complex(key string) complex128 {
	if document.Contains(key) {
		return document.data[key].(complex128)
	}
	return 0.0
}

func (document *Document) Bool(key string) bool {
	if document.Contains(key) {
		return document.data[key].(bool)
	}
	return false
}

func (document *Document) Value(key string) interface{} {
	if document.Contains(key) {
		return document.data[key]
	}
	return nil
}

func (document *Document) Array(key string) []interface{} {
	if document.Contains(key) {
		return document.data[key].([]interface{})
	}
	return nil
}

func (document *Document) ArrayString(key string) []string {
	if document.Contains(key) {
		return document.data[key].([]string)
	}
	return nil
}

func (document *Document) ArrayInt(key string) []int64 {
	if document.Contains(key) {
		return document.data[key].([]int64)
	}
	return nil
}

func (document *Document) ArrayUInt(key string) []uint64 {
	if document.Contains(key) {
		return document.data[key].([]uint64)
	}
	return nil
}

func (document *Document) ArrayFloat(key string) []float64 {
	if document.Contains(key) {
		return document.data[key].([]float64)
	}
	return nil
}

func (document *Document) ArrayComplex(key string) []complex128 {
	if document.Contains(key) {
		return document.data[key].([]complex128)
	}
	return nil
}

func (document *Document) ArrayBool(key string) []bool {
	if document.Contains(key) {
		return document.data[key].([]bool)
	}
	return nil
}

func (document *Document) ToArray() []interface{} {
	array := make([]interface{}, document.Size())
	keys := make([]string, document.Size())
	i := 0
	for k, _ := range document.data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i = 0; i < len(keys); i++ {
		array[i] = []interface{}{keys[i], document.Value(keys[i])}
	}
	return array
}

func (document *Document) GetIterator() *iterator.Iterator {
	it := iterator.New()
	for k, _ := range document.data {
		it.Append(k)
	}
	return it
}

func (document *Document) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := fmt.Sprintf("%T", val); t {
	case "map[string]interface {}":
		document := New()
		aux := val.(map[string]interface{})
		for k, v := range aux {
			document.Add(k, v)
		}
		return document, true
	case "Document":
		aux := val.(Document)
		return &aux, true
	case "*Document":
		aux := val.(*Document)
		return aux, true
	default:
		return nil, false
	}
}
