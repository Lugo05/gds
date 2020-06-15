package set

import (
	. "github.com/Lugo05/gds/collection"
)

type AbstractSet interface {
	Collection
	Intersection(AbstractSet) AbstractSet
	Union(AbstractSet) AbstractSet
	Difference(AbstractSet) AbstractSet
}
