package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromSlice(t *testing.T) {
	set := FromSlice([]string{"a", "b", "c"})
	require.Len(t, set, 3)
}

func TestMakeSet(t *testing.T) {
	set := make(Set[int])
	set.Add(1, 2, 3)
	require.Len(t, set, 3)
}

func TestSetOperations(t *testing.T) {
	set := FromSlice([]int{1, 2, 3})
	require.Len(t, set, 3)
	require.True(t, set.Has(1))
	require.True(t, set.Has(2))
	require.True(t, set.Has(3))
	require.False(t, set.Has(4))

	set.Add(1)
	require.Len(t, set, 3)
	set.Add(-1)
	require.Len(t, set, 4)
	set.Delete(1, -1)
	require.Len(t, set, 2)
	require.False(t, set.Has(1))
	require.False(t, set.Has(-1))

	set.Delete(100)
	require.Len(t, set, 2)

	set.Delete(2, 3, 50000)
	require.Len(t, set, 0)

	set2 := FromSlice([]int{1, 1, 1, 1, 1})
	require.Len(t, set2, 1)
}

func TestSetOperationsStrings(t *testing.T) {
	set := FromSlice([]string{"a", "b", "c"})
	require.Len(t, set, 3)
	require.True(t, set.Has("a"))
	require.True(t, set.Has("b"))
	require.True(t, set.Has("c"))
	require.False(t, set.Has("f"))

	set.Add("foobar")
	require.Len(t, set, 4)
	set.Add("")
	require.Len(t, set, 5)
}

func TestEquals(t *testing.T) {
	set := FromSlice([]int{3, 2, 1})
	set2 := FromSlice([]int{1, 2, 3})
	require.True(t, set.Equals(set2))

	emptySet1 := FromSlice([]int{})
	emptySet2 := FromSlice([]int{})
	require.True(t, emptySet1.Equals(emptySet2))

	set3 := FromSlice([]int{1, 2, 3})
	set4 := FromSlice([]int{1, 2, 3, 4})
	require.False(t, set3.Equals(set4))
	require.False(t, set4.Equals(set3))
}

func TestUnion(t *testing.T) {
	set1 := FromSlice([]int{1, 2, 3})
	set2 := FromSlice([]int{3, 4, 5})

	union := set1.Union(set2)
	require.Len(t, union, 5)

	test := FromSlice([]int{1, 2, 3, 4, 5})
	require.True(t, union.Equals(test))
}

func TestIntersection(t *testing.T) {
	set1 := FromSlice([]int{1, 2, 3})
	set2 := FromSlice([]int{2, 3, 4})

	intersection := set1.Intersection(set2)
	require.Len(t, intersection, 2)

	test := FromSlice([]int{2, 3})
	require.True(t, intersection.Equals(test))
}

func TestDifference(t *testing.T) {
	var a, b, diff, test Set[int]

	a = FromSlice([]int{1, 2, 3})
	b = FromSlice([]int{2, 3, 4})
	diff = a.Difference(b)
	require.Len(t, diff, 1)
	test = FromSlice([]int{1})
	require.True(t, diff.Equals(test))

	// test inverse
	diff = b.Difference(a)
	require.Len(t, diff, 1)
	test = FromSlice([]int{4})
	require.True(t, diff.Equals(test))

	// test what might be a false assumption by a naive user
	// remember, Difference is unidirectional! See SymmetricalDifference
	test = FromSlice([]int{1, 4})
	require.False(t, diff.Equals(test))
}

func TestSymmetricalDifference(t *testing.T) {
	var a, b, diff, test Set[int]

	a = FromSlice([]int{1, 2, 3})
	b = FromSlice([]int{2, 3, 4})
	diff = a.SymmetricalDifference(b)
	require.Len(t, diff, 2)
	test = FromSlice([]int{1, 4})
	require.True(t, diff.Equals(test))

	a = FromSlice([]int{1, 2})
	b = FromSlice([]int{2, 1})
	diff = a.SymmetricalDifference(b)
	require.Len(t, diff, 0)

	a = FromSlice([]int{1, 2})
	b = FromSlice([]int{3, 4})
	diff = a.SymmetricalDifference(b)
	require.Len(t, diff, 4)
	test = FromSlice([]int{1, 2, 3, 4})
	require.True(t, diff.Equals(test))
}
