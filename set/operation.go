package set

func New() *Set {
	set := Set{}
	set.data = make(map[interface{}]int)
	return &set
}

func Clone(s *Set) *Set {
	set := New()
	for k, v := range s.data {
		set.data[k] = v
	}
	return set
}

func (s *Set) Intersection(abstract AbstractSet) AbstractSet {
	t, correct := abstract.(*Set)
	if !correct {
		panic("Method Intersection must receive a set as parameter")
	}
	i := New()
	for k, _ := range s.data {
		if t.Contains(k) {
			i.Add(k)
		}
	}
	return i
}

func (s *Set) Union(abstract AbstractSet) AbstractSet {
	t, correct := abstract.(*Set)
	if !correct {
		panic("Method Union must receive a set as parameter")
	}
	u := New()
	for k, _ := range s.data {
		u.Add(k)
	}
	for k, _ := range t.data {
		u.Add(k)
	}
	return u
}

func (s *Set) Difference(abstract AbstractSet) AbstractSet {
	t, correct := abstract.(*Set)
	if !correct {
		panic("Method Difference must receive a set as parameter")
	}
	d := New()
	for k, _ := range s.data {
		if !t.Contains(k) {
			d.Add(k)
		}
	}
	return d
}
