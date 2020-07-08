package document_test

import (
	"reflect"
	"testing"

	"github.com/Lugo05/gds/document"
)

func SetUp() *document.Document {
	doc := document.New()
	doc.Add("UInt0", uint(1234)).Add("Int0", int32(123)).Add("Float0", float32(1234.1234)).Add("Complex0", complex64(12+34i))
	doc.Add("String0", "Ejemplo0").Add("Boolean0", true)
	doc.Add("ArrayString0", []string{"A", "B"}).Add("ArrayInt0", []rune{12, 34}).Add("ArrayFloat0", []float32{1.2, 3.4})
	doc.Add("ArrayUInt0", []byte{12, 34}).Add("ArrayComplex0", []complex64{1 + 2i, 3 + 4i}).Add("ArrayBoolean0", []bool{true, false})
	doc.Add("ArrayInterface0", []interface{}{"A", 1, 2.34, true})
	doc.Add("DocumentMap0", map[string]interface{}{"Int1": 1234.0, "String1": "Ejemplo1"})

	sub1 := document.New()
	sub1.Add("UInt1", uint16(1)).Add("Int1", int8(1)).Add("Float1", float64(1.1)).Add("Complex1", complex128(1+1i))
	sub1.Add("String1", "Ejemplo1").Add("Boolean1", true)
	sub1.Add("Document1", map[string]interface{}{"Int2": 456, "String2": "Ejemplo2"})

	doc.Add("Document0", sub1)

	return doc
}

func TestDocumentAdd(t *testing.T) {
	matrix := []map[string]string{
		{"key": "Document0", "type": "*document.Document"},
		{"key": "DocumentMap0", "type": "*document.Document"},
		{"key": "String0", "type": "string"},
		{"key": "Int0", "type": "int64"},
		{"key": "UInt0", "type": "uint64"},
		{"key": "Float0", "type": "float64"},
		{"key": "Complex0", "type": "complex128"},
		{"key": "Boolean0", "type": "bool"},
		{"key": "ArrayString0", "type": "[]string"},
		{"key": "ArrayInt0", "type": "[]int64"},
		{"key": "ArrayUInt0", "type": "[]uint64"},
		{"key": "ArrayFloat0", "type": "[]float64"},
		{"key": "ArrayComplex0", "type": "[]complex128"},
		{"key": "ArrayBoolean0", "type": "[]bool"},
	}
	document := SetUp()
	for _, testCase := range matrix {
		if !document.Contains(testCase["key"]) {
			t.Errorf("Error. Expected %v key in Document.\n", testCase["key"])
		} else if document.Type(testCase["key"]) != testCase["type"] {
			t.Errorf("Error. %v should be %v. Got %v.\n", testCase["key"], testCase["type"], document.Type(testCase["key"]))
		}
	}
}

func TestDocumentRemove(t *testing.T) {
	document := SetUp()
	document.Remove("Int0")
	document.Remove("ArrayInt0")
	if document.Contains("Int0") || document.Contains("ArrayInt0") {
		t.Error("Error. Int0 and ArrayInt0 should not be in document")
	}
}

func TestDocumentChange(t *testing.T) {
	matrix := []map[string]interface{}{
		{"typePrev": "int64", "key": "Int0", "typeNext": "float64", "new": float64(1.1)},
		{"typePrev": "float64", "key": "Float0", "typeNext": "float64", "new": float64(1.1)},
		{"typePrev": "[]uint64", "key": "ArrayUInt0", "typeNext": "complex128", "new": complex128(1 + 1i)},
		{"typePrev": "string", "key": "String0", "typeNext": "string", "new": "Ejemplo2"},
	}

	document := SetUp()

	for _, testCase := range matrix {
		key := testCase["key"].(string)
		typePrev := testCase["typePrev"].(string)
		typeNext := testCase["typeNext"].(string)
		new := testCase["new"]

		if !reflect.DeepEqual(typePrev, document.Type(key)) {
			t.Errorf("Error. %v must be %v\n", key, typePrev)
		}
		document.Change(key, new)
		if !reflect.DeepEqual(typeNext, document.Type(key)) {
			t.Errorf("Error. %v should've change to %v but is %v\n", key, typeNext, document.Type(key))
		}
		if !reflect.DeepEqual(new, document.Value(key)) {
			t.Errorf("Error. %v should be now %v but is  %v\n", key, new, document.Value(key))
		}
	}
}

func TestDocumentEquals(t *testing.T) {
	doc1 := SetUp()
	doc2 := SetUp()
	if doc1 == doc2 {
		t.Error("Documents to compare should be different in memory")
	}
	if !doc1.Equals(doc2) {
		t.Error("Documents should be equals at this point")
	}
	doc2.Remove("String0")
	if doc1.Equals(doc2) {
		t.Error("Documents should be different at this point")
	}
}

func TestDocumentToArray(t *testing.T) {
	document := SetUp()
	expected := []interface{}{
		[]interface{}{"Int2", int64(456)},
		[]interface{}{"String2", "Ejemplo2"},
	}
	got := document.Child("Document0").Child("Document1").ToArray()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ToArray Expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentGetIterator(t *testing.T) {
	document := SetUp()
	expected := 15
	got := 0
	it := document.GetIterator()
	for v := it.Begin(); v != nil; v = it.Next() {
		got++
	}
	if expected != got {
		t.Errorf("Error. Expected %v. Got %v.", expected, got)
	}
}

func TestDocumentInt(t *testing.T) {
	var expected int64 = 123
	document := SetUp()
	got := document.Int("Int0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. Int expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentUInt(t *testing.T) {
	var expected uint64 = 1234
	document := SetUp()
	got := document.UInt("UInt0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. UInt expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentFloat(t *testing.T) {
	var expected float64 = 1234.1234
	document := SetUp()
	got := document.Float("Float0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. Float expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentComplex(t *testing.T) {
	var expected complex128 = 12 + 34i
	document := SetUp()
	got := document.Complex("Complex0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. Complex expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentString(t *testing.T) {
	var expected string = "Ejemplo0"
	document := SetUp()
	got := document.String_("String0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. String expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentBool(t *testing.T) {
	var expected bool = true
	document := SetUp()
	got := document.Bool("Boolean0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. Bool expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentValue(t *testing.T) {
	var expected interface{} = "Ejemplo0"
	document := SetUp()
	got := document.Value("String0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. Bool expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArrayInt(t *testing.T) {
	expected := []int64{12, 34}
	document := SetUp()
	got := document.ArrayInt("ArrayInt0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayInt expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArrayUInt(t *testing.T) {
	expected := []uint64{12, 34}
	document := SetUp()
	got := document.ArrayUInt("ArrayUInt0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayUInt expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArrayFloat(t *testing.T) {
	expected := []float64{1.2, 3.4}
	document := SetUp()
	got := document.ArrayFloat("ArrayFloat0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayFloat expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArrayComplex(t *testing.T) {
	expected := []complex128{1 + 2i, 3 + 4i}
	document := SetUp()
	got := document.ArrayComplex("ArrayComplex0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayComplex expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArrayString(t *testing.T) {
	expected := []string{"A", "B"}
	document := SetUp()
	got := document.ArrayString("ArrayString0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayString expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArrayBool(t *testing.T) {
	expected := []bool{true, false}
	document := SetUp()
	got := document.ArrayBool("ArrayBoolean0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayBool expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentArray(t *testing.T) {
	expected := []interface{}{"A", 1, 2.34, true}
	document := SetUp()
	got := document.Array("ArrayInterface0")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. ArrayInterface expected %v %T. Got %v %T", expected, expected, got, got)
	}
}

func TestDocumentChild(t *testing.T) {
	document := SetUp()
	subDoc := document.Child("Document0").Child("Document1")
	subDoc.Add("Float2", 1.32)
	if !document.Child("Document0").Child("Document1").Contains("Float2") {
		//fmt.Println(document.Child("Document0").Child("Document1"))
		t.Error("Error. Document0/Document1/Float2 should exist but it doesn't.\n")
	}
}
