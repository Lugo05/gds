package document

import (
	"fmt"
	"strings"

	. "github.com/Lugo05/gds/collection"
)

type Document struct {
	data map[string]interface{}
}

func (document *Document) String() string {

	stringValue := "(*document.Document) {\n"
	for k, v := range document.data {
		switch document.Type(k) {
		case "*document.Document":
			aux := document.Child(k).String()
			aux = strings.ReplaceAll(aux, "\n", "\n\t")
			stringValue += fmt.Sprintf("\t'%v': %v\n", k, aux)
		default:
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

		if value, valid := isInt(val); valid {
			document.data[key] = value
		} else if value, valid := isUInt(val); valid {
			document.data[key] = value
		} else if value, valid := isFloat(val); valid {
			document.data[key] = value
		} else if value, valid := isComplex(val); valid {
			document.data[key] = value
		} else if value, valid := isString(val); valid {
			document.data[key] = value
		} else if value, valid := isBool(val); valid {
			document.data[key] = value
		} else if value, valid := isArrayString(val); valid {
			document.data[key] = value
		} else if value, valid := isArrayInt(val); valid {
			document.data[key] = value
		} else if value, valid := isArrayFloat(val); valid {
			document.data[key] = value
		} else if value, valid := isArrayComplex(val); valid {
			document.data[key] = value
		} else if value, valid := isArrayInterface(val); valid {
			document.data[key] = value
		} else if value, valid := isDocument(val); valid {
			document.data[key] = value
		} else {
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

func (document *Document) ToArray() []interface{} {
	return nil
}

func (document *Document) Next() interface{} {
	return nil
}

func (document *Document) Equals(c Collection) bool {
	return true
}

func (document *Document) Change(key string, newValue interface{}) *Document {
	document.Remove(key)
	document.Add(key, newValue)
	return document
}

func (document *Document) Type(key string) string {
	if document.Contains(key) {
		return fmt.Sprintf("%T", document.data[key])
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
