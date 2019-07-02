package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

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

type DependencySet struct {
	Deps     map[string]*AppDependency
	AppCalls map[string][]*sysl.Statement
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
	return &DependencySet{map[string]*AppDependency{}, map[string][]*sysl.Statement{}}
}

func MakeAppDependency(self, target *AppElement, stmt *sysl.Statement) *AppDependency {
	return &AppDependency{self, target, stmt}
}

func MakeAppElement(name, endpoint string) *AppElement {
	return &AppElement{name, endpoint}
}

func FindApps(module *sysl.Module, excludes, integrations StrSet, ds []*AppDependency, highlight bool) StrSet {
	output := []string{}
	appReStr := toPattern(integrations.ToSlice())
	re := regexp.MustCompile(appReStr)
	for _, dep := range ds {
		appNames := dep.extractAppNames()
		excludeApps := appNames.Intersection(excludes)
		if len(excludeApps) > 0 {
			continue
		}
		highlightApps := appNames.Intersection(integrations)
		if !highlight && len(highlightApps) == 0 {
			continue
		}
		for _, item := range appNames.ToSlice() {
			app := module.GetApps()[item]
			if highlight {
				if re.MatchString(item) &&
					!hasPattern("human", app.GetAttrs()) {
					output = append(output, item)
				}
				continue
			}
			if !hasPattern("human", app.GetAttrs()) {
				output = append(output, item)
			}
		}
	}

	return MakeStrSet(output...)
}

func walkPassthrough(excludes, passthroughs StrSet, dep *AppDependency, integrations *DependencySet, module *sysl.Module, appCalls map[string][]*sysl.Statement) {
	excludedApps := MakeStrSet(dep.Target.Name).Intersection(excludes)
	undeterminedEps := MakeStrSet(dep.Target.Endpoint).Intersection(MakeStrSet(".. * <- *", "*"))
	//Add to integration array since all dependencies are determined.
	if len(excludedApps) == 0 && len(undeterminedEps) == 0 {
		integrations.Deps[dep.String()] = dep
	}

	// find the next outbound dep
	if passthroughs.Contains(dep.Target.Name) {
		//cs := NewCallStatementSlice()
		//cs.CollectCallStatements(module.GetApps()[dep.Target.Name].GetEndpoints()[dep.Target.Endpoint].GetStmt())
		statements := appCalls[makeAppCallKey(dep.Target.Name, dep.Target.Endpoint)]
		for _, stmt := range statements {
			call := stmt.GetStmt().(*sysl.Statement_Call).Call
			nextAppName := strings.Join(call.GetTarget().GetPart(), " :: ")
			nextEpName := call.GetEndpoint()
			next := MakeAppElement(nextAppName, nextEpName)
			nextDep := MakeAppDependency(dep.Target, next, stmt)
			walkPassthrough(excludes, passthroughs, nextDep, integrations, module, appCalls)
		}
	}
}

func (ds *DependencySet) FindIntegrations(apps, excludes, passthroughs StrSet, module *sysl.Module) *DependencySet {
	integrations := NewDependencySet()
	outboundDeps := NewDependencySet()
	lenPassthroughs := len(passthroughs)
	commonEndpoints := MakeStrSet(".. * <- *", "*")
	for _, dep := range ds.Deps {
		appNames := dep.extractAppNames()
		endpoints := dep.extractEndpoints()
		isSubsection := appNames.IsSubSet(apps)
		isSelfSubsection := MakeStrSet(dep.Self.Name).IsSubSet(apps)
		isTargetSubsection := MakeStrSet(dep.Target.Name).IsSubSet(passthroughs)
		interExcludes := appNames.Intersection(excludes)
		interEndpoints := endpoints.Intersection(commonEndpoints)
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
			walkPassthrough(excludes, passthroughs, dep, integrations, module, ds.AppCalls)
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
	appEndpt := makeAppCallKey(appname, epname)
	for _, stat := range stmts {
		switch c := stat.GetStmt().(type) {
		case *sysl.Statement_Call:
			targetName := getAppName(c.Call.GetTarget())
			dep := MakeAppDependency(MakeAppElement(appname, epname), MakeAppElement(targetName, c.Call.GetEndpoint()), stat)
			ds.Deps[dep.String()] = dep
			if len(ds.AppCalls[appEndpt]) == 0 {
				ds.AppCalls[appEndpt] = []*sysl.Statement{}
			}
			ds.AppCalls[appEndpt] = append(ds.AppCalls[appEndpt], stat)
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

func toPattern(comp []string) string {
	return fmt.Sprintf(`^(?:%s)(?: *::|$)`, strings.Join(comp, "|"))
}

func makeAppCallKey(appname, epname string) string {
	return strings.Join([]string{appname, epname}, ":")
}
