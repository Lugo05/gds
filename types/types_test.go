package types_test

import (
	_ "errors"
	"testing"

	"github.com/Lugo05/gds/types"
)

/*

matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}

*/

func TestIsIntTrue(t *testing.T) {
	matrix := []interface{}{int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1)}
	for _, testcase := range matrix {
		if _, valid := types.Int.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as Int.\n", testcase, testcase)
		}
	}
}

func TestIsIntFalse(t *testing.T) {
	matrix := []interface{}{
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.Int.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as Int.\n", testcase, testcase)
		}
	}
}

func TestIsUIntTrue(t *testing.T) {
	matrix := []interface{}{uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1)}
	for _, testcase := range matrix {
		if _, valid := types.UInt.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as UInt.\n", testcase, testcase)
		}
	}
}

func TestIsUIntFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.UInt.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as UInt.\n", testcase, testcase)
		}
	}
}

func TestIsBoolTrue(t *testing.T) {
	matrix := []interface{}{true, false}
	for _, testcase := range matrix {
		if _, valid := types.Bool.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as Bool.\n", testcase, testcase)
		}
	}
}

func TestIsBoolFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi",
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.Bool.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as Bool.\n", testcase, testcase)
		}
	}
}

func TestIsComplexTrue(t *testing.T) {
	matrix := []interface{}{complex64(1 + 1i), complex128(1 + 1i)}
	for _, testcase := range matrix {
		if _, valid := types.Complex.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as Complex.\n", testcase, testcase)
		}
	}
}

func TestIsComplexFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.Complex.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as Complex.\n", testcase, testcase)
		}
	}
}

func TestIsFloatTrue(t *testing.T) {
	matrix := []interface{}{float32(1.1), float64(1.1)}
	for _, testcase := range matrix {
		if _, valid := types.Float.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as Float.\n", testcase, testcase)
		}
	}
}

func TestIFloatFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		"hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.Float.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as Float.\n", testcase, testcase)
		}
	}
}

func TestIsStringTrue(t *testing.T) {
	matrix := []interface{}{"hi", "bye"}
	for _, testcase := range matrix {
		if _, valid := types.String.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as String.\n", testcase, testcase)
		}
	}
}

func TestIsStringFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.String.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as String.\n", testcase, testcase)
		}
	}
}

func TestIsArrayIntTrue(t *testing.T) {
	matrix := []interface{}{
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2},
		[]interface{}{1, 2, 3, 4}, []rune{1, 2},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayInt.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as ArrayInt.\n", testcase, testcase)
		}
	}
}

func TestIsArrayIntFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayInt.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as ArrayInt.\n", testcase, testcase)
		}
	}
}

func TestIsArrayUIntTrue(t *testing.T) {
	matrix := []interface{}{
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2},
		[]interface{}{uint(1), uint(2), uint(3)}, []byte{1, 2},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayUInt.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as ArrayUInt.\n", testcase, testcase)
		}
	}
}

func TestIsArrayUIntFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayUInt.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as ArrayUInt.\n", testcase, testcase)
		}
	}
}

func TestIsArrayFloatTrue(t *testing.T) {
	matrix := []interface{}{
		[]float32{1.1, 2.1}, []float64{1.1, 2.1},
		[]interface{}{1.1, 2.1, 3.1, 4.1},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayFloat.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as ArrayFloat.\n", testcase, testcase)
		}
	}
}

func TestIsArrayFloatFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]string{"hi", "bye"}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayFloat.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as ArrayFloat.\n", testcase, testcase)
		}
	}
}

func TestIsArrayComplexTrue(t *testing.T) {
	matrix := []interface{}{
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1 + 1i, 1 + 1i},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayComplex.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as ArrayComplex.\n", testcase, testcase)
		}
	}
}

func TestIsArrayComplexFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"}, []bool{true, true},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayComplex.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as ArrayComplex.\n", testcase, testcase)
		}
	}
}

func TestIsArrayBoolTrue(t *testing.T) {
	matrix := []interface{}{
		[]bool{true, true}, []bool{false, true},
		[]interface{}{true, false},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayBool.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as ArrayBool.\n", testcase, testcase)
		}
	}
}

func TestIsArrayBoolFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []string{"hi", "bye"},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayBool.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as ArrayBool.\n", testcase, testcase)
		}
	}
}

func TestIsArrayStringTrue(t *testing.T) {
	matrix := []interface{}{
		[]string{"hi", "How", "are", "you", "?"},
		[]interface{}{"Good", "bye"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayString.IsAssignable(testcase); !valid {
			t.Errorf("Error. %v with type %T must be assignable as ArrayString.\n", testcase, testcase)
		}
	}
}

func TestIsArrayStringFalse(t *testing.T) {
	matrix := []interface{}{
		int8(1), int16(1), int(1), int32(1), int64(1), float32(1), float64(1), rune(1),
		uint8(1), uint16(1), uint(1), uint32(1), uint64(1), byte(1),
		float32(1.1), float64(1.1), "hi", true,
		complex64(1 + 1i), complex128(1 + 1i),
		[]int8{1, 2}, []int16{1, 2}, []int{1, 2}, []int32{1, 2}, []int64{1, 2}, []rune{1, 2},
		[]uint8{1, 2}, []uint16{1, 2}, []uint{1, 2}, []uint32{1, 2}, []uint64{1, 2}, []byte{1, 2},
		[]float32{1.1, 1.1}, []float64{1.1, 1.1}, []bool{true, true},
		[]complex64{1 + 1i, 1 + 1i}, []complex128{1 + 1i, 1 + 1i},
		[]interface{}{1, 2, 3, 4, 5, 6, 7, "8"},
	}
	for _, testcase := range matrix {
		if _, valid := types.ArrayString.IsAssignable(testcase); valid {
			t.Errorf("Error. %v with type %T shouldn't be assignable as ArrayString.\n", testcase, testcase)
		}
	}
}
