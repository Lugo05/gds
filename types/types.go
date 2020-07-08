package types

import (
	"fmt"
	"strconv"
	"strings"
)

type Type interface {
	IsAssignable(interface{}) (interface{}, bool)
}

type int_ struct{}
type uInt_ struct{}
type float_ struct{}
type complex_ struct{}
type string_ struct{}
type bool_ struct{}
type arrayString_ struct{}
type arrayInt_ struct{}
type arrayUInt_ struct{}
type arrayFloat_ struct{}
type arrayComplex_ struct{}
type arrayBool_ struct{}
type arrayInterface_ struct{}

var (
	Int            = int_{}
	UInt           = uInt_{}
	Float          = float_{}
	Complex        = complex_{}
	Bool           = bool_{}
	String         = string_{}
	ArrayInt       = arrayInt_{}
	ArrayUInt      = arrayUInt_{}
	ArrayFloat     = arrayFloat_{}
	ArrayComplex   = arrayComplex_{}
	ArrayBool      = arrayBool_{}
	ArrayString    = arrayString_{}
	ArrayInterface = arrayInterface_{}
)

func getTypeVariable(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func (i int_) IsAssignable(val interface{}) (interface{}, bool) {
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

func (u uInt_) IsAssignable(val interface{}) (interface{}, bool) {
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

func (f float_) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "float32":
		//return float64(val.(float32)), true
		valStr := fmt.Sprintf("%v", val)
		v, _ := strconv.ParseFloat(valStr, len(valStr)-strings.Index(valStr, "."))
		return v, true
	case "float64":
		return val.(float64), true
	default:
		return 0, false
	}
}

func (c complex_) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "complex64":
		return complex128(val.(complex64)), true
	case "complex128":
		return val.(complex128), true
	default:
		return 0, false
	}
}

func (s string_) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "string":
		return val.(string), true
	default:
		return "", false
	}
}

func (b bool_) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "bool":
		return val.(bool), true
	default:
		return false, false
	}
}

func (ai arrayInt_) IsAssignable(val interface{}) (interface{}, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]int64, len(array))
		for i, v := range array {
			if value, valid := Int.IsAssignable(v); valid {
				aux[i] = value.(int64)
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

func (au arrayUInt_) IsAssignable(val interface{}) (interface{}, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]uint64, len(array))
		for i, v := range array {
			if value, valid := UInt.IsAssignable(v); valid {
				aux[i] = value.(uint64)
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

func (af arrayFloat_) IsAssignable(val interface{}) (interface{}, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]float64, len(array))
		for i, v := range array {
			if value, valid := Float.IsAssignable(v); valid {
				aux[i] = value.(float64)
			} else {
				return nil, false
			}
		}
		return aux, true
	case "[]float32":
		array := make([]float64, len(val.([]float32)))
		for i, v := range val.([]float32) {
			value, _ := Float.IsAssignable(v)
			array[i] = value.(float64)
		}
		return array, true
	case "[]float64":
		return val.([]float64), true
	default:
		return nil, false
	}
}

func (ac arrayComplex_) IsAssignable(val interface{}) (interface{}, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]complex128, len(array))
		for i, v := range array {
			if value, valid := Complex.IsAssignable(v); valid {
				aux[i] = value.(complex128)
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
	case "[]complex128":
		return val.([]complex128), true
	default:
		return nil, false
	}
}

func (as arrayString_) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]string, len(array))
		for i, v := range array {
			if _, valid := String.IsAssignable(v); !valid {
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

func (ab arrayBool_) IsAssignable(val interface{}) (interface{}, bool) {

	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		array := val.([]interface{})
		aux := make([]bool, len(array))
		for i, v := range array {
			if value, valid := Bool.IsAssignable(v); valid {
				aux[i] = value.(bool)
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

func (a arrayInterface_) IsAssignable(val interface{}) (interface{}, bool) {
	switch t := getTypeVariable(val); t {
	case "[]interface {}":
		return val.([]interface{}), true
	default:
		return nil, false
	}
}
