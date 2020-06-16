package document

import (
	"errors"
	"fmt"
)

func arrayString(variable []interface{}) ([]string, error) {
	array := make([]string, len(variable))
	for i, value := range variable {
		typeValue := getTypeVariable(value)
		if typeValue != "string" {
			return nil, errors.New(fmt.Sprintf("Element %v is not a string. It is %T", value, value))
		}
		array[i] = value.(string)
	}
	return array, nil
}

func getTypeVariable(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func isInt(val interface{}) (int64, bool) {
	switch t := getTypeVariable(val); t {
	case "float32":
		aux := val.(float32)
		if float32(int64(aux)) == aux {
			return int64(aux), true
		} else {
			return 0, false
		}
	case "float64":
		aux := val.(float64)
		if float64(int64(aux)) == aux {
			return int64(aux), true
		} else {
			return 0, false
		}
	case "int":
		return int64(val.(int)), true
	case "int8":
		return int64(val.(int8)), true
	case "int16":
		return int64(val.(int16)), true
	case "int32":
		return int64(val.(int32)), true
	case "rune":
		return int64(val.(rune)), true
	case "int64":
		return val.(int64), true
	default:
		return 0, false
	}
}

func isUInt(val interface{}) (uint64, bool) {
	switch t := getTypeVariable(val); t {
	case "uint":
		return uint64(val.(uint)), true
	case "uint8":
		return uint64(val.(uint8)), true
	case "byte":
		return uint64(val.(byte)), true
	case "uint16":
		return uint64(val.(uint16)), true
	case "uint32":
		return uint64(val.(uint32)), true
	case "uint64":
		return val.(uint64), true
	default:
		return 0, false
	}
}

func isFloat(val interface{}) (float64, bool) {
	switch t := getTypeVariable(val); t {
	case "float32":
		return float64(val.(float32)), true
	case "float64":
		return val.(float64), true
	default:
		return 0, false
	}
}

func isComplex(val interface{}) (complex128, bool) {
	switch t := getTypeVariable(val); t {
	case "complex64":
		return complex128(val.(complex64)), true
	case "complex128":
		return val.(complex128), true
	default:
		return 0, false
	}
}

func isString(val interface{}) (string, bool) {
	switch t := getTypeVariable(val); t {
	case "string":
		return val.(string), true
	default:
		return "", false
	}
}

func isBool(val interface{}) (bool, bool) {
	switch t := getTypeVariable(val); t {
	case "bool":
		return val.(bool), true
	default:
		return false, false
	}
}

func isArrayString(val interface{}) ([]string, bool) {
	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]string, len(array))
		for i, v := range array {
			if _, valid := isString(v); !valid {
				return nil, false
			}
			aux[i] = v.(string)
		}
		return aux, true
	case "[]string":
		return val.([]string), true
	default:
		return nil, false
	}
}

func isArrayInt(val interface{}) ([]int64, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]int64, len(array))
		for i, v := range array {
			if value, valid := isInt(v); valid {
				aux[i] = value
			} else {
				return nil, false
			}
		}
		return aux, true
	case "[]int":
		array := make([]int64, len(val.([]int)))
		for i, v := range val.([]int) {
			array[i] = int64(v)
		}
		return array, true
	case "[]int8":
		array := make([]int64, len(val.([]int8)))
		for i, v := range val.([]int8) {
			array[i] = int64(v)
		}
		return array, true
	case "[]int16":
		array := make([]int64, len(val.([]int16)))
		for i, v := range val.([]int16) {
			array[i] = int64(v)
		}
		return array, true
	case "[]int32":
		array := make([]int64, len(val.([]int32)))
		for i, v := range val.([]int32) {
			array[i] = int64(v)
		}
		return array, true
	case "[]rune":
		array := make([]int64, len(val.([]rune)))
		for i, v := range val.([]rune) {
			array[i] = int64(v)
		}
		return array, true
	case "[]int64":
		return val.([]int64), true
	default:
		return nil, false
	}
}

func isArrayUInt(val interface{}) ([]uint64, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]uint64, len(array))
		for i, v := range array {
			if value, valid := isUInt(v); valid {
				aux[i] = value
			} else {
				return nil, false
			}
		}
		return aux, true
	case "[]uint":
		array := make([]uint64, len(val.([]uint)))
		for i, v := range val.([]uint) {
			array[i] = uint64(v)
		}
		return array, true
	case "[]uint8":
		array := make([]uint64, len(val.([]uint8)))
		for i, v := range val.([]uint8) {
			array[i] = uint64(v)
		}
		return array, true
	case "[]uint16":
		array := make([]uint64, len(val.([]uint16)))
		for i, v := range val.([]uint16) {
			array[i] = uint64(v)
		}
		return array, true
	case "[]uint32":
		array := make([]uint64, len(val.([]uint32)))
		for i, v := range val.([]uint32) {
			array[i] = uint64(v)
		}
		return array, true
	case "[]byte":
		array := make([]uint64, len(val.([]byte)))
		for i, v := range val.([]byte) {
			array[i] = uint64(v)
		}
		return array, true
	case "[]uint64":
		return val.([]uint64), true
	default:
		return nil, false
	}
}

func isArrayFloat(val interface{}) ([]float64, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]float64, len(array))
		for i, v := range array {
			if value, valid := isFloat(v); valid {
				aux[i] = value
			} else {
				return nil, false
			}
		}
		return aux, true
	case "[]float32":
		array := make([]float64, len(val.([]float32)))
		for i, v := range val.([]float32) {
			array[i] = float64(v)
		}
		return array, true
	case "[]float64":
		return val.([]float64), true
	default:
		return nil, false
	}
}

func isArrayComplex(val interface{}) ([]complex128, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]complex128, len(array))
		for i, v := range array {
			if value, valid := isComplex(v); valid {
				aux[i] = value
			} else {
				return nil, false
			}
		}
		return aux, true
	case "[]complex64":
		array := make([]complex128, len(val.([]complex64)))
		for i, v := range val.([]complex64) {
			array[i] = complex128(v)
		}
		return array, true
	case "[]float64":
		return val.([]complex128), true
	default:
		return nil, false
	}
}

func isArrayBool(val interface{}) ([]bool, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]bool, len(array))
		for i, v := range array {
			if value, valid := isBool(v); valid {
				aux[i] = value
			} else {
				return nil, false
			}
		}
		return aux, true
	case "[]bool":
		return val.([]bool), true
	default:
		return nil, false
	}
}

func isArrayInterface(val interface{}) ([]interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		return val.([]interface{}), true
	default:
		return nil, false
	}
}

func isDocument(val interface{}) (*Document, bool) {
	switch t := getTypeVariable(val); t {
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
