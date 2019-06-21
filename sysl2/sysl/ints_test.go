package main

import (
	"bytes"
	"flag"
	"testing"

	"github.com/anz-bank/sysl/src/proto"
	"github.com/stretchr/testify/assert"
)

func TestGenerateIntegrations(t *testing.T) {
	m, _ := Parse("demo/simple/sysl-ints.sysl", "../../")
	stmt := &sysl.Statement{}
	args := MakeArgs("", "Project", false, false)
	apps := []string{"System1", "IntegratedSystem", "System2"}
	highlights := MakeStrSet("IntegratedSystem", "System1", "System2")
	s1 := MakeAppElement("IntegratedSystem", "integrated_endpoint_1")
	t1 := MakeAppElement("System1", "endpoint")
	dep1 := MakeAppDependency(s1, t1, stmt)
	s2 := MakeAppElement("IntegratedSystem", "integrated_endpoint_2")
	t2 := MakeAppElement("System2", "endpoint")
	dep2 := MakeAppDependency(s2, t2, stmt)
	deps := []*AppDependency{dep1, dep2}
	endpt := &sysl.Endpoint{
		Name: "_",
		Stmt: []*sysl.Statement{
			{
				Stmt: &sysl.Statement_Action{
					Action: &sysl.Action{
						Action: "IntegratedSystem",
					},
				},
			},
			{
				Stmt: &sysl.Statement_Action{
					Action: &sysl.Action{
						Action: "System1",
					},
				},
			},
			{
				Stmt: &sysl.Statement_Action{
					Action: &sysl.Action{
						Action: "System2",
					},
				},
			},
		},
	}
	intsParam := MakeIntsParam(apps, highlights, deps, m.GetApps()["Project"], endpt)
	r := GenerateView(args, intsParam, m)

	expected := `''''''''''''''''''''''''''''''''''''''''''
''                                      ''
''  AUTOGENERATED CODE -- DO NOT EDIT!  ''
''                                      ''
''''''''''''''''''''''''''''''''''''''''''

@startuml
hide stereotype
scale max 16384 height
skinparam component {
  BackgroundColor FloralWhite
  BorderColor Black
  ArrowColor Crimson
}
[IntegratedSystem] as _0 <<highlight>>
[System1] as _1 <<highlight>>
_0 --> _1
[System2] as _2 <<highlight>>
_0 --> _2
@enduml
`

	assert.NotNil(t, m)
	assert.NotNil(t, r)
	assert.Equal(t, expected, r)
}

type intsArg struct {
	root_model string
	title      string
	output     string
	project    string
	filter     string
	modules    string
	exclude    []string
	clustered  bool
	epa        bool
}

func TestGenerateIntegrationsWithTestFile(t *testing.T) {
	// Given
	args := &intsArg{
		root_model: "./tests/",
		modules:    "integration_test.sysl",
		output:     "%(epname).png",
		project:    "Project",
	}
	expectContent := `''''''''''''''''''''''''''''''''''''''''''
''                                      ''
''  AUTOGENERATED CODE -- DO NOT EDIT!  ''
''                                      ''
''''''''''''''''''''''''''''''''''''''''''

@startuml
hide stereotype
scale max 16384 height
skinparam component {
  BackgroundColor FloralWhite
  BorderColor Black
  ArrowColor Crimson
}
[IntegratedSystem] as _0 <<highlight>>
[System1] as _1 <<highlight>>
_0 --> _1
[System2] as _2 <<highlight>>
_0 --> _2
@enduml
`
	expected := map[string]string{
		"_.png": expectContent,
	}

	// When
	result := GenerateIntegrations(args.root_model, args.title, args.output,
		args.project, args.filter, args.modules, args.exclude, args.clustered, args.epa)

	// Then
	assert.Equal(t, expected, result, "unexpected content!")
}

func TestGenerateIntegrationsWithDependencySetTestFile(t *testing.T) {
	// Given
	args := &intsArg{
		root_model: "./tests/",
		modules:    "integration_dependency_set_test.sysl",
		output:     "%(epname).png",
		project:    "Project",
	}
	expectContent := `''''''''''''''''''''''''''''''''''''''''''
''                                      ''
''  AUTOGENERATED CODE -- DO NOT EDIT!  ''
''                                      ''
''''''''''''''''''''''''''''''''''''''''''

@startuml
hide stereotype
scale max 16384 height
skinparam component {
  BackgroundColor FloralWhite
  BorderColor Black
  ArrowColor Crimson
}
[IntegratedSystem] as _0 <<highlight>>
[System1] as _1 <<highlight>>
_0 --> _1
[System2] as _2 <<highlight>>
_0 --> _2
@enduml
`
	expected := map[string]string{
		"_.png": expectContent,
	}

	// When
	result := GenerateIntegrations(args.root_model, args.title, args.output,
		args.project, args.filter, args.modules, args.exclude, args.clustered, args.epa)

	// Then
	assert.Equal(t, expected, result, "unexpected content!")
}

func TestGenerateIntegrationsWithTestFileAndFilters(t *testing.T) {
	// Given
	args := &intsArg{
		root_model: "./tests/",
		modules:    "integration_test.sysl",
		output:     "%(epname).png",
		project:    "Project",
		filter:     "test",
	}
	expected := map[string]string{}

	// When
	result := GenerateIntegrations(args.root_model, args.title, args.output,
		args.project, args.filter, args.modules, args.exclude, args.clustered, args.epa)

	// Then
	assert.Equal(t, expected, result, "unexpected content!")
}

func TestDoGenerateIntegrations(t *testing.T) {
	type args struct {
		flags *flag.FlagSet
		args  []string
	}
	argsData := []string{"ints"}
	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantStderr string
	}{
		{
			"Case-Do generate integrations",
			args{
				flag.NewFlagSet(argsData[0], flag.PanicOnError),
				argsData,
			},
			"",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			DoGenerateIntegrations(stdout, stderr, tt.args.flags, tt.args.args)
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("DoGenerateIntegrations() = %v, want %v", gotStdout, tt.wantStdout)
			}
			if gotStderr := stderr.String(); gotStderr != tt.wantStderr {
				t.Errorf("DoGenerateIntegrations() = %v, want %v", gotStderr, tt.wantStderr)
			}
		})
	}
}
