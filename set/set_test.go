package set_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/Lugo05/gds/set"
)

func TestSetAdd(t *testing.T) {

	set := set.New()
	set.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)
	set.Add(1, 2, 3, 4, 5, 6, 10, 11, 12)
	set.Add(13, 14, 15, 16, "1", "2", "3")
	expected := 19
	got := set.Size()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Error. Expected %v. Gotten %v", expected, got)
	}
}

func TestSetRemove(t *testing.T) {
	set := set.New()
	set.Add("A", "B", "C", "D", "E", "F", "G", "H", "I")
	set.Remove("B").Remove("C").Remove("D")
	if set.Contains("B") || set.Contains("C") || set.Contains("D") {
		t.Errorf("Error. B,C and D must not be in the set")
	}
}

func TestSetToArray(t *testing.T) {
	set := set.New()
	set.Add(1, 2, 2, 2, 3, 4, 5, 1, 3, 5, 6, 8) //1,2,3,4,5,6,8
	set.Remove(6).Remove(8)
	expected := []interface{}{1, 2, 3, 4, 5}
	got := set.ToArray()
	sort.Slice(got, func(i, j int) bool {
		return got[i].(int) < got[j].(int)
	})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Error. Expected %v. Gotten %v", expected, got)
	}
}

func TestSetNext(t *testing.T) {
	set := set.New()
	set.Add(1, 2, 3, 4, 5, 6)
	expected := 12
	got := 0
	for next := set.Next(); next != nil; next = set.Next() {
		got++
	}
	for next := set.Next(); next != nil; next = set.Next() {
		got++
	}
	if got != expected {
		t.Errorf("Error. Expected %v. Got %v", expected, got)
	}
}

func TestSetEqualTrue(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6)
	set2 := set.New()
	set2.Add(2, 1, 4, 3, 6, 5)
	expected := true
	got := set1.Equals(set2)
	if got != expected {
		t.Errorf("Error. Expected %v. Got %v", expected, got)
	}
}

func TestSetEqualFalse(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6)
	set2 := set.New()
	set2.Add("1", "2", "3", "4", "5", "6")
	expected := false
	got := set1.Equals(set2)
	if got != expected {
		t.Errorf("Error. Expected %v. Got %v", expected, got)
	}
}
