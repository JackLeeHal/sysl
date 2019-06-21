package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleExp(t *testing.T) {
	//Given
	input := `%(epname)`
	attr := map[string]string{
		"epname": "searchP",
	}

	//When
	result := ParseFmt(attr, input)

	//Then
	assert.Equal(t, "searchP", result)
}

func TestExpWithQn(t *testing.T) {
	//Given
	input := `%(@c2?//%(@c4?--%(cc?dd|edd)--|bc)//\n|cc)`
	attr := map[string]string{}

	//When
	result := ParseFmt(attr, input)

	//Then
	assert.Equal(t, `cc`, result)
}

func TestExpWithCmp(t *testing.T) {
	//Given
	input := `%(@c2=='aaa'?//%(@c4?--%(cc?dd|edd)--|bc)//\n|cc)`
	attr := map[string]string{
		"@c2": "aaa",
	}

	//When
	result := ParseFmt(attr, input)

	//Then
	assert.Equal(t, `//bc//\n`, result)
}

func TestFmtWithPrefix(t *testing.T) {
	//Given
	input := `1ba%%%%(DT?%(@c2?//%(@c4?--%(cc?dd|edd)--|bc)//\n|cc)\|bb\)**%(appname)**`
	attr := map[string]string{}

	//When
	result := ParseFmt(attr, input)

	//Then
	assert.Equal(t, `1ba%%(DT?cc\|bb\)****`, result)
}

func TestFmtWithRecursion(t *testing.T) {
	//Given
	input := `%(@ggg?//«%(@ggg)»//**%(patterns? %(patterns~/\btba|tbd\b/?<color red>%(patterns)</color>|<color green>%(patterns)</color>)| <color red>pattern?</color>)**\n|%(int?<color red>(missing INT%)</color>\n))%(epname)%(args?\n(%(args)%))`
	attr := map[string]string{
		"@s1":       "TT",
		"patterns":  "rt → hp, ap",
		"int": "int",
		"epname":    "searchP",
	}

	//When
	result := ParseFmt(attr, input)

	//Then
	assert.Equal(t, `<color red>(missing INT)</color>\nsearchP`, result)
}

func TestExpWithQnWithoutNoStmt(t *testing.T) {
	//Given
	input := `%(args?\n(%(args)%))`
	attr := map[string]string{
		"args": "<color red>(missing INT)</color>\nLogin\n(<color blue>where.Token</color> <<color red>R, T</color>>)",
	}

	//When
	result := ParseFmt(attr, input)

	// Then
	assert.Equal(t, `\n(<color red>(missing INT)</color>\nLogin\n(<color blue>where.Token</color> <<color red>R, T</color>>))`, result)
}
