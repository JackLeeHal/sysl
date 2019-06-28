package main

import (
	"testing"

	"github.com/anz-bank/sysl/src/proto"
	"github.com/stretchr/testify/assert"
)

func TestGetAppName(t *testing.T) {
	// given
	a := &sysl.AppName{
		Part: []string{"test", "name"},
	}

	// when
	actual := getAppName(a)

	// then
	assert.Equal(t, "test :: name", actual, "unexpected result")
}

func TestGetApp(t *testing.T) {
	// Given
	m := &sysl.Module{
		Apps: map[string]*sysl.Application{
			"test :: name": {
				Attrs: map[string]*sysl.Attribute{},
			},
		},
	}
	a := &sysl.AppName{
		Part: []string{"test", "name"},
	}
	expected := &sysl.Application{
		Attrs: map[string]*sysl.Attribute{},
	}

	// When
	actual := getApp(a, m)

	// Then
	assert.Equal(t, expected, actual)
}

func TestHasAbstractPattern(t *testing.T) {
	// Given
	attrs := map[string]*sysl.Attribute{
		"patterns": {
			Attribute: &sysl.Attribute_A{
				A: &sysl.Attribute_Array{
					Elt: []*sysl.Attribute{
						{
							Attribute: &sysl.Attribute_S{
								S: "abstract",
							},
						},
						{
							Attribute: &sysl.Attribute_S{
								S: "human",
							},
						},
					},
				},
			},
		},
	}

	// When
	actual := hasPattern("abstract", attrs)

	// Then
	assert.Equal(t, true, actual)
}

func TestHasNotAbstractPattern2(t *testing.T) {
	// Given
	attrs := map[string]*sysl.Attribute{
		"patterns": {
			Attribute: &sysl.Attribute_A{
				A: &sysl.Attribute_Array{
					Elt: []*sysl.Attribute{
						{
							Attribute: &sysl.Attribute_S{
								S: "ui",
							},
						},
						{
							Attribute: &sysl.Attribute_S{
								S: "human",
							},
						},
					},
				},
			},
		},
	}

	// When
	actual := hasPattern("abstract", attrs)

	// Then
	assert.Equal(t, false, actual)
}

func TestIsNotSameAppWithPartLength(t *testing.T) {
	// Given
	a := &sysl.AppName{
		Part: []string{"test", "name"},
	}
	b := &sysl.AppName{
		Part: []string{"name1"},
	}

	// When
	actual := isSameApp(a, b)

	// Then
	assert.Equal(t, false, actual)
}

func TestIsNotSameAppWithPartContent(t *testing.T) {
	// Given
	a := &sysl.AppName{
		Part: []string{"test", "name"},
	}
	b := &sysl.AppName{
		Part: []string{"test", "name1"},
	}

	// When
	actual := isSameApp(a, b)

	// Then
	assert.Equal(t, false, actual)
}

func TestIsSameApp(t *testing.T) {
	// Given
	a := &sysl.AppName{
		Part: []string{"test", "name"},
	}
	b := &sysl.AppName{
		Part: []string{"test", "name"},
	}

	// When
	actual := isSameApp(a, b)

	// Then
	assert.Equal(t, true, actual)
}

func TestIsSameCall(t *testing.T) {
	// Given
	a := &sysl.Call{
		Target: &sysl.AppName{
			Part: []string{"test", "name"},
		},
		Endpoint: "endpt",
	}
	b := &sysl.Call{
		Target: &sysl.AppName{
			Part: []string{"test", "name"},
		},
		Endpoint: "endpt",
	}

	// When
	actual := isSameCall(a, b)

	// Then
	assert.Equal(t, true, actual)
}
