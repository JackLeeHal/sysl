package main

import (
	"fmt"
	"github.com/anz-bank/sysl/src/proto"
)

type AppElement struct {
	Name     string
	Endpoint string
}

type AppDependency struct {
	Self      *AppElement
	Target    *AppElement
	Statement *sysl.Statement
}

func (dep *AppDependency) String() string {
	return fmt.Sprintf("%s:%s:%s:%s", dep.Self.Name, dep.Self.Endpoint, dep.Target.Name, dep.Target.Endpoint)
}

func MakeAppDependency(self, target *AppElement, stmt *sysl.Statement) *AppDependency {
	return &AppDependency{self, target, stmt}
}

func MakeAppElement(name, endpoint string) *AppElement {
	return &AppElement{name, endpoint}
}
