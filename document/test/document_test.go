package document_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Lugo05/gds/set"
)

func TestCollectionAdd(t *testing.T) {
	document := SetUp()

	s := set.New()
	s.Add(1, 2)
	document.Child("Documento0").Add("Conjunto1", s)
	aux := document.Child("Documento0").Value("Conjunto1").(*set.Set)
	aux.Add(3, 4, 5)

	expected := 5
	got := document.Child("Documento0").Value("Conjunto1").(*set.Set).Size()
	if expected != got {
		t.Error(fmt.Sprintf("Error. Expected %v. Got %v", expected, got))
	}
}

func TestCollectionChange(t *testing.T) {
	matrix := []map[string]interface{}{
		{"data": []interface{}{"Prueba1", "String", "CambioString"}, "expected": "CambioString"},
		{"data": []interface{}{"Prueba2", 20, 4}, "expected": int64(4)},
		{"data": []interface{}{"Prueba3", []int{1, 2}, true}, "expected": true},
		{"data": []interface{}{"Prueba4", []int{1, 2}, []int{1}}, "expected": []int64{1}},
		{"data": []interface{}{"Prueba5", "Cadena", 3.1415902}, "expected": 3.1415902},
	}

	document := SetUp()
	fmt.Println(document)
	for _, testCase := range matrix {
		key := testCase["data"].([]interface{})[0].(string)
		old := testCase["data"].([]interface{})[1]
		new := testCase["data"].([]interface{})[2]
		document.Add(key, old)
		document.Change(key, new)
		got := document.Value(key)
		if !reflect.DeepEqual(got, testCase["expected"]) {
			t.Error(fmt.Sprintf("Error. Expected %v. Got %v. ", testCase["expected"], got))
		}
	}
}
