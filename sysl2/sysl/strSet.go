package main

import (
	"sort"

	"github.com/anz-bank/sysl/src/proto"
)

type StrSet map[string]struct{}

func MakeStrSet(initial ...string) StrSet {
	s := StrSet{}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

func MakeStrSetFromSpecificAttr(attr string, attrs map[string]*sysl.Attribute) StrSet {
	s := StrSet{}

	if patterns, has := attrs[attr]; has {
		if x := patterns.GetA(); x != nil {
			for _, y := range x.Elt {
				if v := y.GetS(); len(v) > 0 {
					s.Insert(y.GetS())
				}
			}
		}
	}

	return s
}

func MakeStrSetFromActionStatement(stmts []*sysl.Statement) StrSet {
	s := StrSet{}
	for _, stmt := range stmts {
		if _, ok := stmt.Stmt.(*sysl.Statement_Action); ok {
			if v := stmt.GetAction().GetAction(); len(v) > 0 {
				s.Insert(v)
			}
		}
	}

	return s
}

func (s StrSet) Contains(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s StrSet) Insert(elem string) {
	s[elem] = struct{}{}
}

func (s StrSet) Remove(elem string) {
	delete(s, elem)
}

func (s StrSet) ToSlice() []string {
	o := make([]string, 0, len(s))

	for k := range s {
		o = append(o, k)
	}

	return o
}

func (s StrSet) ToSortedSlice() []string {
	slice := s.ToSlice()
	sort.Strings(slice)

	return slice
}

func (s StrSet) Clone() StrSet {
	out := StrSet{}

	for k := range s {
		out.Insert(k)
	}

	return out
}

func (s StrSet) Union(other StrSet) StrSet {
	out := StrSet{}

	for k := range s {
		out.Insert(k)
	}

	for k := range other {
		out.Insert(k)
	}

	return out
}

func (s StrSet) Intersection(other StrSet) StrSet {
	out := StrSet{}

	for k := range s {
		if other.Contains(k) {
			out.Insert(k)
		}
	}

	return out
}

func (s StrSet) Difference(other StrSet) StrSet {
	if len(other) == 0 {
		return s
	}
	out := StrSet{}

	for k := range s {
		if !other.Contains(k) {
			out.Insert(k)
		}
	}

	return out
}

// Returns true if child set is a subset of parent set
func (s StrSet) IsSubSet(parent StrSet) bool {
	if len(s) == 0 && len(parent) == 0 {
		return true
	}
	if len(parent) == 0 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	return len(s.Difference(parent)) == 0
}
