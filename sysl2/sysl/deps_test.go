package main

import (
	"testing"

	"github.com/anz-bank/sysl/src/proto"
	"github.com/stretchr/testify/assert"
)

func TestAppDependency_String(t *testing.T) {
	// Given
	stmt := &sysl.Statement{}
	dep := MakeAppDependency(MakeAppElement("AppA", "EndptA"), MakeAppElement("AppB", "EndptB"), stmt)
	expected := "AppA:EndptA:AppB:EndptB"

	// When
	actual := dep.String()

	// Then
	assert.Equal(t, actual, expected)
}
