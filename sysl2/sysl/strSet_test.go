package main

import (
	"testing"

	"github.com/anz-bank/sysl/src/proto"
	"github.com/stretchr/testify/assert"
)

func TestMakeStrSet(t *testing.T) {
	a := MakeStrSet("a", "b", "c", "e")
	assert.Equal(t, 4, len(a), "Unexpected result")
}

func TestMakeStrSetWithDuplicateInitialValues(t *testing.T) {
	a := MakeStrSet("a", "b", "c", "e", "a", "a", "c")
	assert.Equal(t, 4, len(a), "Unexpected result")
}

func TestMakeStrSetWithEmptyStringInitialValues(t *testing.T) {
	a := MakeStrSet("a", "b", "c", "e", "a", "a", "c", "", "")
	assert.Equal(t, 5, len(a), "Unexpected result")
}

func TestMakeStrSetWithoutInitialValues(t *testing.T) {
	a := MakeStrSet()
	assert.Equal(t, 0, len(a), "Unexpected result")
}

func TestMakeStrSetFromSpecificAttrWithEmptyAttrs(t *testing.T) {
	attrs := map[string]*sysl.Attribute{}

	a := MakeStrSetFromSpecificAttr("patterns", attrs)
	assert.Equal(t, 0, len(a), "Unexpected result")
}

func TestMakeStrSetFromSpecificAttrWithoutPatternAttr(t *testing.T) {
	attrs := map[string]*sysl.Attribute{
		"test": {Attribute: &sysl.Attribute_S{S: "test"}},
	}

	a := MakeStrSetFromSpecificAttr("patterns", attrs)
	assert.Equal(t, 0, len(a), "Unexpected result")
}

func TestMakeStrSetFromPatternsAttr(t *testing.T) {
	attrs := map[string]*sysl.Attribute{
		"patterns": {
			Attribute: &sysl.Attribute_A{
				A: &sysl.Attribute_Array{
					Elt: []*sysl.Attribute{
						{Attribute: &sysl.Attribute_S{S: "test"}},
					},
				},
			},
		},
	}

	a := MakeStrSetFromSpecificAttr("patterns", attrs)
	assert.Equal(t, 1, len(a), "Unexpected result")
}

func TestMakeStrSetFromActionStatement(t *testing.T) {
	stmts := []*sysl.Statement{
		{
			Stmt: &sysl.Statement_Action{
				Action: &sysl.Action{
					Action: "AppA",
				},
			},
		},
	}

	a := MakeStrSetFromActionStatement(stmts)
	assert.Equal(t, 1, len(a), "Unexpected result")
}

func TestContains(t *testing.T) {
	a := MakeStrSet("a", "b", "c", "e")
	assert.Equal(t, 4, len(a), "Unexpected result")
	assert.True(t, a.Contains("b"), "Unexpected result")
	assert.False(t, a.Contains("d"), "Unexpected result")
}

func TestInsert(t *testing.T) {
	a := MakeStrSet("a", "b", "c", "e")
	assert.Equal(t, 4, len(a), "Unexpected result")
	assert.True(t, a.Contains("b"), "Unexpected result")
	assert.False(t, a.Contains("d"), "Unexpected result")

	a.Insert("d")
	assert.Equal(t, 5, len(a), "Unexpected result")
	assert.True(t, a.Contains("b"), "Unexpected result")
	assert.True(t, a.Contains("d"), "Unexpected result")
}

func TestRemove(t *testing.T) {
	a := MakeStrSet("a", "b", "c", "e")
	assert.Equal(t, 4, len(a), "Unexpected result")
	assert.True(t, a.Contains("b"), "Unexpected result")
	assert.False(t, a.Contains("d"), "Unexpected result")

	a.Remove("d")
	assert.Equal(t, 4, len(a), "Unexpected result")
	assert.True(t, a.Contains("b"), "Unexpected result")
	assert.False(t, a.Contains("d"), "Unexpected result")

	a.Remove("b")
	assert.Equal(t, 3, len(a), "Unexpected result")
	assert.False(t, a.Contains("b"), "Unexpected result")
	assert.False(t, a.Contains("d"), "Unexpected result")
}

func TestToSlice(t *testing.T) {
	// Given
	a := MakeStrSet("c", "b", "a", "e")

	// When
	slice := a.ToSlice()
	sorted := a.ToSortedSlice()

	// Then
	sameValue := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		set := map[string]struct{}{}
		for _, v := range a {
			set[v] = struct{}{}
		}
		for _, v := range b {
			if _, ok := set[v]; !ok {
				return false
			}
		}
		return true
	}

	assert.True(t, sameValue([]string{"a", "b", "c", "e"}, slice), "Unexpected result")
	assert.Equal(t, []string{"a", "b", "c", "e"}, sorted, "Unexpected result")
}

func TestClone(t *testing.T) {
	a := MakeStrSet("c", "b", "a", "e")
	b := a.Clone()
	assert.Equal(t, a, b, "Unexpected result")

	b.Remove("c")
	assert.NotEqual(t, a, b, "Unexpected result")
}

func TestIntersection(t *testing.T) {
	a := MakeStrSet("c", "b", "a", "e")
	b := MakeStrSet("d", "b", "a", "e")

	c := a.Intersection(b)
	assert.Equal(t, 3, len(c), "Unexpected result")
	assert.Equal(t, []string{"a", "b", "e"}, c.ToSortedSlice(), "Unexpected result")
}

func TestDifference(t *testing.T) {
	a := MakeStrSet("c", "b", "a", "e")
	b := MakeStrSet("d", "b", "a", "e")

	c := a.Difference(b)
	assert.Equal(t, 1, len(c), "Unexpected result")
	assert.Equal(t, []string{"c"}, c.ToSortedSlice(), "Unexpected result")
}
