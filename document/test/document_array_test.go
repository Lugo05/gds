package document_test

import (
	"fmt"
	"testing"
)

func TestCollectionArray(t *testing.T) {
	document := SetUp()

	expected := []int{4, 3, 0, 0}
	got := make([]int, 4)

	document.Add("Array", []interface{}{1, 2.0, "Cadena", true})
	document.Child("Collection").Add("Array", []interface{}{"A", "B", 3})
	document.Child("Collection").Child("Collection").Add("ArrayFloat", []float64{3.2, 4.5, 6.0, 1.2})

	got[0] = len(document.Array("Array"))
	got[1] = len(document.Child("Collection").Array("Array"))
	got[2] = len(document.Child("Collection").Child("Collection").Array("ArrayFloat"))
	got[3] = len(document.Child("Collection").Child("Collection").Array("Array"))

	for i, exp := range expected {
		if exp != got[i] {
			t.Error(fmt.Sprintf("Error. Expected %v. Got %v.", exp, got[i]))
		}
	}
}

func TestCollectionArrayString(t *testing.T) {
	document := SetUp()

	expected := []int{4, 3, 0, 0}
	got := make([]int, 4)

	document.Add("ArrayString", []string{"a", "b", "c", "d"})
	document.Child("Collection").Add("ArrayString", []string{"a", "b", "c"})
	document.Child("Collection").Child("Collection").Add("ArrayInt", []int{1, 2, 3, 4, 5, 6, 7, 8})

	got[0] = len(document.ArrayString("ArrayString"))
	got[1] = len(document.Child("Collection").ArrayString("ArrayString"))
	got[2] = len(document.Child("Collection").Child("Collection").ArrayString("ArrayInt"))
	got[3] = len(document.Child("Collection").Child("Collection").ArrayString("ArrayString"))

	for i, exp := range expected {
		if exp != got[i] {
			t.Error(fmt.Sprintf("Error. Expected %v. Got %v.", exp, got[i]))
		}
	}
}

func TestCollectionArrayInt(t *testing.T) {
	document := SetUp()

	expected := []int{4, 3, 0, 0}
	got := make([]int, 4)

	document.Add("ArrayInt", []int64{1, 2, 3, 4})
	document.Child("Collection").Add("ArrayInt", []int64{1, 2, 3})
	document.Child("Collection").Child("Collection").Add("ArrayFloat", []float64{3.2, 4.5, 6.0, 1.2})

	got[0] = len(document.ArrayInt("ArrayInt"))
	got[1] = len(document.Child("Collection").ArrayInt("ArrayInt"))
	got[2] = len(document.Child("Collection").Child("Collection").ArrayInt("ArrayFloat"))
	got[3] = len(document.Child("Collection").Child("Collection").ArrayInt("ArrayInt"))

	for i, exp := range expected {
		if exp != got[i] {
			t.Error(fmt.Sprintf("Error. Expected %v. Got %v.", exp, got[i]))
		}
	}
}

func TestCollectionArrayFloat(t *testing.T) {
	document := SetUp()

	expected := []int{4, 3, 0, 0}
	got := make([]int, 4)

	document.Add("ArrayFloat", []float64{1.5, 2.5, 3.5, 4.5})
	document.Child("Collection").Add("ArrayFloat", []float64{1.1, 2.2, 3.3})
	document.Child("Collection").Child("Collection").Add("ArrayInt", []bool{true, false, false, false, true})

	got[0] = len(document.ArrayFloat("ArrayFloat"))
	got[1] = len(document.Child("Collection").ArrayFloat("ArrayFloat"))
	got[2] = len(document.Child("Collection").Child("Collection").ArrayFloat("ArrayInt"))
	got[3] = len(document.Child("Collection").Child("Collection").ArrayFloat("ArrayFloat"))

	for i, exp := range expected {
		if exp != got[i] {
			t.Error(fmt.Sprintf("Error. Expected %v. Got %v.", exp, got[i]))
		}
	}
}

func TestCollectionArrayBool(t *testing.T) {
	document := SetUp()

	expected := []int{4, 3, 0, 0}
	got := make([]int, 4)

	document.Add("ArrayBool", []bool{true, true, true, false})
	document.Child("Collection").Add("ArrayBool", []bool{false, false, true})
	document.Child("Collection").Child("Collection").Add("ArrayInt", []int{1, 2, 3, 4, 5, 6})

	got[0] = len(document.ArrayBool("ArrayBool"))
	got[1] = len(document.Child("Collection").ArrayBool("ArrayBool"))
	got[2] = len(document.Child("Collection").Child("Collection").ArrayBool("ArrayInt"))
	got[3] = len(document.Child("Collection").Child("Collection").ArrayBool("ArrayBool"))

	for i, exp := range expected {
		if exp != got[i] {
			t.Error(fmt.Sprintf("Error. Expected %v. Got %v.", exp, got[i]))
		}
	}
}

/*func TestCollectionArray(t *testing.T) {
    document := document()

    expected := []bool{false,false,true}
    got := make([]bool,3)
    got[0] = document.Child("Collection").Child("Collection").Bool("Booleano")
    got[1] = document.Child("Collection").Child("Collection").Bool("Cadena")
    got[2] = document.Child("Collection").Bool("Booleano");

    for i,exp := range expected {
        if exp != got[i] {
            t.Error(fmt.Sprintf("Error. Expected %v. Got %v.",exp,got[i]))
        }
    }
}*/
