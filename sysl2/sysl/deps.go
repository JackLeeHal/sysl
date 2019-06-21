package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/anz-bank/sysl/src/proto"
	"github.com/yale8848/stream"
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

type DependencySet struct {
	Deps map[string]*AppDependency
}

func (dep *AppDependency) String() string {
	return fmt.Sprintf("%s:%s:%s:%s", dep.Self.Name, dep.Self.Endpoint, dep.Target.Name, dep.Target.Endpoint)
}

func (ds *DependencySet) ToSlice() []*AppDependency {
	o := []*AppDependency{}
	keys := []string{}
	for k := range ds.Deps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		o = append(o, ds.Deps[k])
	}

	return o
}

func (dep *AppDependency) extractAppNames() StrSet {
	s := StrSet{}
	s.Insert(dep.Self.Name)
	s.Insert(dep.Target.Name)
	return s
}

func (dep *AppDependency) extractEndpoints() StrSet {
	s := StrSet{}
	s.Insert(dep.Self.Endpoint)
	s.Insert(dep.Target.Endpoint)
	return s
}

func NewDependencySet() *DependencySet {
	return &DependencySet{map[string]*AppDependency{}}
}

func MakeAppDependency(self, target *AppElement, stmt *sysl.Statement) *AppDependency {
	return &AppDependency{self, target, stmt}
}

func MakeAppElement(name, endpoint string) *AppElement {
	return &AppElement{name, endpoint}
}

type CallStatementSlice struct {
	slice []*sysl.Statement
}

func NewCallStatementSlice() *CallStatementSlice {
	return &CallStatementSlice{}
}

func (cs *CallStatementSlice) Add(call *sysl.Statement) {
	cs.slice = append(cs.slice, call)
}

func (cs *CallStatementSlice) GetSlice() []*sysl.Statement {
	return cs.slice
}

func (cs *CallStatementSlice) CollectCallStatements(stmts []*sysl.Statement) {
	for _, stat := range stmts {
		switch c := stat.GetStmt().(type) {
		case *sysl.Statement_Call:
			cs.Add(stat)
		case *sysl.Statement_Action, *sysl.Statement_Ret:
			continue
		case *sysl.Statement_Cond:
			cs.CollectCallStatements(c.Cond.GetStmt())
		case *sysl.Statement_Loop:
			cs.CollectCallStatements(c.Loop.GetStmt())
		case *sysl.Statement_LoopN:
			cs.CollectCallStatements(c.LoopN.GetStmt())
		case *sysl.Statement_Foreach:
			cs.CollectCallStatements(c.Foreach.GetStmt())
		case *sysl.Statement_Group:
			cs.CollectCallStatements(c.Group.GetStmt())
		case *sysl.Statement_Alt:
			for _, choice := range c.Alt.GetChoice() {
				cs.CollectCallStatements(choice.GetStmt())
			}
		default:
			panic("No statement!")
		}
	}
}

func FindApps(module *sysl.Module, excludes, integrations StrSet, ds []*AppDependency, highlight bool) StrSet {
	r := []string{}
	appReStr := toPattern(integrations.ToSlice())
	re := regexp.MustCompile(appReStr)
	st, _ := stream.Of(ds)
	st.ForEach(func(v stream.T) bool {
		if dep, ok := v.(*AppDependency); ok {
			appNames := dep.extractAppNames()
			excludeApps := appNames.Intersection(excludes)
			if len(excludeApps) > 0 {
				return true
			}
			highlightApps := appNames.Intersection(integrations)
			if !highlight && len(highlightApps) == 0 {
				return true
			}
			appStream, _ := stream.Of(appNames.ToSlice())
			appStream.ForEach(func(v stream.T) bool {
				item := v.(string)
				app := module.GetApps()[item]
				if highlight {
					if re.MatchString(item) &&
						!hasPattern("human", app.GetAttrs()) {
						r = append(r, item)
					}
					return true
				}
				if !hasPattern("human", app.GetAttrs()) {
					r = append(r, item)
				}
				return true
			})
		}
		return true
	})

	return MakeStrSet(r...)
}

func walkPassthrough(excludes, passthroughs StrSet, dep *AppDependency, integrations *DependencySet, module *sysl.Module) {
	target := dep.Target
	targetName := target.Name
	targetEndpoint := target.Endpoint
	excludedApps := MakeStrSet(targetName).Intersection(excludes)
	undeterminedEps := MakeStrSet(targetEndpoint).Intersection(MakeStrSet(".. * <- *", "*"))
	//Add to integration array since all dependencies are determined.
	if len(excludedApps) == 0 && len(undeterminedEps) == 0 {
		integrations.Deps[dep.String()] = dep
	}

	// find the next outbound dep
	if passthroughs.Contains(targetName) {
		cs := NewCallStatementSlice()
		cs.CollectCallStatements(module.GetApps()[targetName].GetEndpoints()[targetEndpoint].GetStmt())
		for _, stmt := range cs.GetSlice() {
			call := stmt.GetStmt().(*sysl.Statement_Call).Call
			nextAppName := strings.Join(call.GetTarget().GetPart(), " :: ")
			nextEpName := call.GetEndpoint()
			next := MakeAppElement(nextAppName, nextEpName)
			nextDep := MakeAppDependency(target, next, stmt)
			walkPassthrough(excludes, passthroughs, nextDep, integrations, module)
		}
	}
}

func (ds *DependencySet) FindIntegrations(apps, excludes, passthroughs StrSet, module *sysl.Module) *DependencySet {
	integrations := NewDependencySet()
	outboundDeps := NewDependencySet()
	lenPassthroughs := len(passthroughs)
	for _, dep := range ds.Deps {
		appNames := dep.extractAppNames()
		endpoints := dep.extractEndpoints()
		isSubsection := isSub(appNames, apps)
		isSelfSubsection := isSub(MakeStrSet(dep.Self.Name), apps)
		isTargetSubsection := isSub(MakeStrSet(dep.Target.Name), passthroughs)
		interExcludes := appNames.Intersection(excludes)
		interEndpoints := endpoints.Intersection(MakeStrSet(".. * <- *", "*"))
		lenInterExcludes := len(interExcludes)
		lenInterEndpoints := len(interEndpoints)
		if isSubsection && lenInterExcludes == 0 && lenInterEndpoints == 0 {
			integrations.Deps[dep.String()] = dep
		}
		// collect outbound dependencies
		if lenPassthroughs > 0 &&
			((isSubsection || (isSelfSubsection && isTargetSubsection)) && lenInterExcludes == 0 && lenInterEndpoints == 0) {
			outboundDeps.Deps[dep.String()] = dep
		}
	}
	if lenPassthroughs > 0 {
		for _, dep := range outboundDeps.Deps {
			walkPassthrough(excludes, passthroughs, dep, integrations, module)
		}
	}

	return integrations
}

func (ds *DependencySet) CollectAppDependencies(module *sysl.Module) {
	for appname, app := range module.GetApps() {
		for epname, endpoint := range app.GetEndpoints() {
			ds.collectStatementDependencies(endpoint.GetStmt(), appname, epname)
		}
	}
}

func (ds *DependencySet) collectStatementDependencies(stmts []*sysl.Statement, appname, epname string) {
	for _, stat := range stmts {
		switch c := stat.GetStmt().(type) {
		case *sysl.Statement_Call:
			targetName := getAppName(c.Call.GetTarget())
			dep := MakeAppDependency(MakeAppElement(appname, epname), MakeAppElement(targetName, c.Call.GetEndpoint()), stat)
			ds.Deps[dep.String()] = dep
		case *sysl.Statement_Action, *sysl.Statement_Ret:
			continue
		case *sysl.Statement_Cond:
			ds.collectStatementDependencies(c.Cond.GetStmt(), appname, epname)
		case *sysl.Statement_Loop:
			ds.collectStatementDependencies(c.Loop.GetStmt(), appname, epname)
		case *sysl.Statement_LoopN:
			ds.collectStatementDependencies(c.LoopN.GetStmt(), appname, epname)
		case *sysl.Statement_Foreach:
			ds.collectStatementDependencies(c.Foreach.GetStmt(), appname, epname)
		case *sysl.Statement_Group:
			ds.collectStatementDependencies(c.Group.GetStmt(), appname, epname)
		case *sysl.Statement_Alt:
			for _, choice := range c.Alt.GetChoice() {
				ds.collectStatementDependencies(choice.GetStmt(), appname, epname)
			}
		default:
			panic("No statement!")
		}
	}
}

func isSub(child, parent StrSet) bool {
	if len(child) == 0 && len(parent) == 0 {
		return true
	}
	if len(parent) == 0 {
		return false
	}
	if len(child) == 0 {
		return true
	}
	return len(child.Difference(parent)) == 0
}

func toPattern(comp []string) string {
	return fmt.Sprintf(`^(?:%s)(?: *::|$)`, strings.Join(comp, "|"))
}
