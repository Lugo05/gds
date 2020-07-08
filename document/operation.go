package document

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func New() *Document {
	document := Document{}
	document.data = make(map[string]interface{})
	return &document
}

func Clone(d *Document) *Document {
	document := New()
	for k, v := range d.data {
		if d.Type(k) == "*document.Document" {
			document.Add(k, Clone(d.Child(k)))
		} else {
			document.Add(k, v)
		}
	}
	return document
}

func MapToDocument(data map[string]interface{}) (*Document, error) {
	document := New()
	for key, value := range data {
		document.Add(key, value)
	}
	return document, nil
}

func getDataFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return []byte(byteValue), nil
}

func JsonToDocument(path string) (*Document, error) {
	var data map[string]interface{}
	byteValue, err := getDataFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		return nil, err
	}
	document, err := MapToDocument(data)
	if err != nil {
		return nil, err
	}
	return document, nil
}

func JsonToArrayDocument(path string) ([]*Document, error) {
	var dataArray []map[string]interface{}
	byteValue, err := getDataFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(byteValue), &dataArray)
	if err != nil {
		return nil, err
	}

	arrayDocument, err := ArrayMapToArrayDocument(dataArray)
	if err != nil {
		return nil, err
	}
	return arrayDocument, nil
}

func ArrayMapToArrayDocument(data []map[string]interface{}) ([]*Document, error) {
	arrayDocument := make([]*Document, len(data))

	for i, value := range data {
		arrayDocument[i] = New()
		document, err := MapToDocument(value)
		if err != nil {
			return nil, err
		}
		arrayDocument[i] = document
	}

	return arrayDocument, nil
}
