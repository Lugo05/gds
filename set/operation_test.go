package set_test

import (
	"testing"

	"github.com/Lugo05/gds/set"
)

func TestIntersectionWithElements(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6, 7, 8)
	set2 := set.New()
	set2.Add(1, 2, 12, 9, 10, 11, 8)
	i := set1.Intersection(set2)
	if i.Size() != 3 {
		t.Errorf("Error. Size must be 3. Set %v", i)
	}
	if !i.Contains(1) || !i.Contains(2) || !i.Contains(8) {
		t.Errorf("Error. Set must have 1,2and 8 but contains %v", i)
	}
}

func TestIntersectionWithoutElements(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6, 7, 8)
	set2 := set.New()
	set2.Add(9, 10, 11, 12, 13, 14)
	i := set1.Intersection(set2)
	if i.Size() != 0 {
		t.Errorf("Error. Size must be 0. Set %v", i)
	}
}

func TestUnionWithTraslape(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6, 7, 8)
	set2 := set.New()
	set2.Add(1, 2, 9, 10, 11, 8)
	u := set1.Union(set2)
	if u.Size() != 11 {
		t.Errorf("Error. Size must be 11. Set %v", u)
	}
}

func TestUnionWithoutTraslape(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6, 7, 8)
	set2 := set.New()
	set2.Add(9, 10, 11)
	u := set1.Union(set2)
	if u.Size() != 11 {
		t.Errorf("Error. Size must be 11. Set %v", u)
	}
}

func TestDiferenceWithElements(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5, 6, 7, 8)
	set2 := set.New()
	set2.Add(1, 2, 3, 4, 5)
	d := set1.Difference(set2)
	if d.Size() != 3 {
		t.Errorf("Error. Size must be 3. Set %v", d)
	}
	if !d.Contains(6) || !d.Contains(7) || !d.Contains(8) {
		t.Errorf("Error. Set must have 1,2and 8 but contains %v", d)
	}
}

func TestDiferenceWithoutElements(t *testing.T) {
	set1 := set.New()
	set1.Add(1, 2, 3, 4, 5)
	set2 := set.New()
	set2.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	d := set1.Difference(set2)
	if d.Size() != 0 {
		t.Errorf("Error. Size must be 0. Set %v", d)
	}
}
