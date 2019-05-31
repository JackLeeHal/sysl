package main

import (
	"flag"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/anz-bank/sysl/src/proto"
	"github.com/ghemawat/stream"
	log "github.com/sirupsen/logrus"
)

type AppElement struct {
	Name     string
	Endpoint string
}

type AppDependency struct {
	Self   *AppElement
	Target *AppElement
}

type DependencySet struct {
	slice []*AppDependency
}

type Calls struct {
	callSlice []*sysl.Call
}

func (dep *AppDependency) Equals(target *AppDependency) bool {
	return (dep.Self.Name == target.Self.Name) && (dep.Self.Endpoint == target.Self.Endpoint) &&
		(dep.Target.Name == target.Target.Name) && (dep.Target.Endpoint == target.Target.Endpoint)
}

func (set *DependencySet) Add(dep *AppDependency) {
	if !set.Contains(dep) {
		set.slice = append(set.slice, dep)
	}
}

func (set *DependencySet) Contains(dep *AppDependency) bool {
	for _, v := range set.slice {
		if v.Equals(dep) {
			return true
		}
	}

	return false
}

func NewCalls() *Calls {
	return &Calls{[]*sysl.Call{}}
}

func NewDependencySet() *DependencySet {
	return &DependencySet{[]*AppDependency{}}
}

func MakeAppDependency(self, target *AppElement) *AppDependency {
	return &AppDependency{self, target}
}

func MakeAppElement(name, endpoint string) *AppElement {
	return &AppElement{name, endpoint}
}

func (calls *Calls) makeCalls(statements []*sysl.Statement) {
	for _, stat := range statements {
		switch c := stat.GetStmt().(type) {
		case *sysl.Statement_Call:
			calls.callSlice = append(calls.callSlice, c.Call)
		case *sysl.Statement_Action:
			continue
		case *sysl.Statement_Cond:
			calls.makeCalls(c.Cond.GetStmt())
		case *sysl.Statement_Loop:
			calls.makeCalls(c.Loop.GetStmt())
		case *sysl.Statement_LoopN:
			calls.makeCalls(c.LoopN.GetStmt())
		case *sysl.Statement_Foreach:
			calls.makeCalls(c.Foreach.GetStmt())
		case *sysl.Statement_Group:
			calls.makeCalls(c.Group.GetStmt())
		case *sysl.Statement_Alt:
			for _, choice := range c.Alt.GetChoice() {
				calls.makeCalls(choice.GetStmt())
			}
		case *sysl.Statement_Ret:
			continue
		default:
			panic("No statement!")
		}
	}
}

func formatAppName(appNames []string) string {
	return strings.Join(appNames, " :: ")
}

func checkDependencies(module *sysl.Module) []*AppDependency {
	var errStr string
	depSet := NewDependencySet()
	errors := []string{}
	apps := module.GetApps()
	for appname, app := range apps {
		for epname, endpoint := range app.GetEndpoints() {
			calls := NewCalls()

			calls.makeCalls(endpoint.GetStmt())
			for _, call := range calls.callSlice {
				targetName := formatAppName(call.GetTarget().GetPart())
				targetApp := apps[targetName]
				if targetApp == nil {
					errStr = fmt.Sprintf("%s <- %s: calls non-existent app %s", appname, epname, targetName)
					errors = append(errors, errStr)
				} else {
					isValid := !hasPattern("abstract", extractApplicationAttr("patterns", targetApp))
					if !isValid {
						panic(fmt.Sprintf("call target '%s' must not be ~abstract", targetName))
					}
					callEndpoint := call.GetEndpoint()
					if targetApp.GetEndpoints()[callEndpoint] == nil {
						errStr = fmt.Sprintf(
							"%s <- %s: calls non-existent endpoint %s -> %s", appname, epname, targetName, callEndpoint)
						errors = append(errors, errStr)
					} else {
						selfApp := MakeAppElement(appname, epname)
						targetApp := MakeAppElement(targetName, callEndpoint)
						dep := MakeAppDependency(selfApp, targetApp)
						depSet.Add(dep)
					}
				}
			}
		}
	}

	if len(errors) > 0 {
		panic(fmt.Sprintf("broken deps:\n  %s", strings.Join(errors, "\n")))
	}
	return depSet.slice
}

func (dep *AppDependency) extractAppNames() []string {
	apps := []string{}
	apps = append(apps, dep.Self.Name)
	apps = append(apps, dep.Target.Name)

	return apps
}

func (dep *AppDependency) extractEndpoints() []string {
	eps := []string{}
	eps = append(eps, dep.Self.Endpoint)
	eps = append(eps, dep.Target.Endpoint)

	return eps
}

func arrayToMap(a []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}

	return m
}

func intersection(a, b []string) []string {
	result := []string{}
	m := arrayToMap(a)
	for _, k := range b {
		if _, ok := m[k]; ok {
			result = append(result, k)
		}
	}

	return result
}

// execute like a - b
func subtraction(a, b []string) []string {
	result := []string{}
	if b == nil || len(b) == 0 {
		return a
	}
	if a == nil || len(a) == 0 {
		return result
	}
	m := arrayToMap(b)
	for _, k := range a {
		if _, has := m[k]; !has {
			result = append(result, k)
		}
	}

	return result
}

func isSub(child, parent []string) bool {
	if parent == nil || len(parent) == 0 {
		return false
	}
	if child == nil || len(child) == 0 {
		return true
	}
	return len(subtraction(parent, child)) == 0
}

func hasPattern(a string, arr []string) bool {
	m := arrayToMap(arr)

	return m[a]
}

func toPattern(comp []string) string {
	return fmt.Sprintf(`^(?:%s)(?: *::|$)`, strings.Join(comp, "|"))
}

func findMatchingApps(module *sysl.Module, excludes []string, integrations []string, deps []*AppDependency) []string {
	result := []string{}

	appReStr := toPattern(integrations)
	re := regexp.MustCompile(appReStr)
	for _, dep := range deps {
		appNames := dep.extractAppNames()
		inter := intersection(appNames, excludes)
		if len(inter) > 0 {
			continue
		}
		filtered, err := stream.Contents(stream.Items(appNames...), stream.If(func(item string) bool {
			app := module.GetApps()[item]
			return re.MatchString(item) && !hasPattern("human", extractApplicationAttr("patterns", app))
		}))
		if err != nil {
			log.Error(err)
		}
		result = append(result, filtered...)
	}

	log.Warnf("------------------matching app:")
	for _, v := range result {
		log.Warnf(v)
	}
	return result
}

func findApps(module *sysl.Module, excludes, matchingApps []string, deps []*AppDependency) []string {
	result := []string{}
	for _, dep := range deps {
		appNames := dep.extractAppNames()
		interExcludes := intersection(appNames, excludes)
		if len(interExcludes) > 0 {
			continue
		}
		interMatchings := intersection(appNames, matchingApps)
		if len(interMatchings) == 0 {
			continue
		}
		filtered, err := stream.Contents(stream.Items(appNames...), stream.If(func(item string) bool {
			app := module.GetApps()[item]
			return !hasPattern("human", extractApplicationAttr("patterns", app))
		}))
		if err != nil {
			log.Error(err)
		}
		result = append(result, filtered...)
	}

	log.Warnf("------------------ apps:")
	for _, v := range result {
		log.Warnf(v)
	}
	return result
}

func walkPassthrough(excludes, passthroughs []string, dep *AppDependency, integrations []*AppDependency, module *sysl.Module) {
	target := dep.Target
	targetName := target.Name
	targetEndpoint := target.Endpoint
	interApp := intersection([]string{targetName}, excludes)
	interEndpoint := intersection([]string{targetEndpoint}, excludes)
	// add to integration array
	if len(interApp) == 0 && len(interEndpoint) == 0 {
		integrations = append(integrations, dep)
	}

	// find the next outbound dep
	if hasPattern(targetName, passthroughs) {
		endpointStmts := module.GetApps()[targetName].GetEndpoints()[targetEndpoint].GetStmt()
		calls := NewCalls()

		calls.makeCalls(endpointStmts)
		for _, call := range calls.callSlice {
			nextAppName := strings.Join(call.GetTarget().GetPart(), " :: ")
			nextEpName := call.GetEndpoint()
			next := MakeAppElement(nextAppName, nextEpName)
			nextDep := MakeAppDependency(target, next)
			walkPassthrough(excludes, passthroughs, nextDep, integrations, module)
		}
	}
}

func findIntegrations(apps, excludes, passthroughs []string, deps []*AppDependency, module *sysl.Module) []*AppDependency {
	integrations := []*AppDependency{}
	outboundDeps := []*AppDependency{}
	lenPassthroughs := len(passthroughs)
	for _, dep := range deps {
		appNames := dep.extractAppNames()
		endpoints := dep.extractEndpoints()
		isSubsection := isSub(appNames, apps)
		isSelfSubsection := isSub([]string{appNames[0]}, apps)
		isTargetSubsection := isSub([]string{appNames[1]}, apps)
		interExcludes := intersection(appNames, excludes)
		interEndpoints := intersection(endpoints, []string{".. * <- *", "*"})
		lenInterExcludes := len(interExcludes)
		lenInterEndpoints := len(interEndpoints)
		if isSubsection && lenInterExcludes == 0 && lenInterEndpoints == 0 {
			integrations = append(integrations, dep)
		}
		// collect outbound dependencies
		if lenPassthroughs > 0 &&
			(isSubsection || isSelfSubsection && isTargetSubsection && lenInterExcludes == 0 && lenInterEndpoints == 0) {
			outboundDeps = append(outboundDeps, dep)
		}
	}
	if lenPassthroughs > 0 {
		for _, dep := range outboundDeps {
			walkPassthrough(excludes, passthroughs, dep, integrations, module)
		}
	}

	return integrations
}

func extractApplicationAttr(attrStr string, app *sysl.Application) []string {
	return transformAttributes(app.GetAttrs()[attrStr])
}

func extractEndpointAttr(attrStr string, endpt *sysl.Endpoint) []string {
	return transformAttributes(endpt.GetAttrs()[attrStr])
}

func transformAttributes(attribute *sysl.Attribute) []string {
	result := []string{}
	attrs := []*sysl.Attribute{}
	if attribute != nil {
		attrs = attribute.GetA().GetElt()
	}
	for _, v := range attrs {
		result = append(result, v.GetS())
	}

	return result
}

func extractAction(stmts []*sysl.Statement) []string {
	result := []string{}
	for _, v := range stmts {
		action := v.GetAction()
		if action == nil {
			panic("No action statement!")
		}
		result = append(result, action.GetAction())
	}

	return result
}

func GenerateIntegrations(
	root_model, title, plantuml, output, project, filter, modules string,
	exclude []string,
	clustered, epa bool,
) {
	log.Warnf("-----------------root:")

	log.Warnf(root_model)
	log.Warnf("------------------modules:")
	log.Warnf(modules)
	mod := loadApp(root_model, modules)
	if len(exclude) == 0 && project != "" {
		exclude = append(exclude, project)
	}
	deps := checkDependencies(mod)

	var out_fmt func(output string) string
	// The "project" app that specifies the required view of the integrations
	app := mod.GetApps()[project]
	// Interate over each endpoint within the selected project
	for _, endpt := range app.GetEndpoints() {
		// build the set of excluded items
		excludes := extractEndpointAttr("exclude", endpt)
		passthroughs := extractEndpointAttr("passthrough", endpt)
		// endpt.stmt's "action" will conatain the "apps" whose integration is to be drawn
		// each one of these will be placed into the "integrations" list
		integrations := extractAction(endpt.GetStmt())

		matchingApps := findMatchingApps(mod, exclude, integrations, deps)
		apps := findApps(mod, exclude, matchingApps, deps)
		apps = subtraction(apps, excludes)
		apps = subtraction(apps, passthroughs)
		output = out_fmt(output)

		if filter != "" {
			re := regexp.MustCompile(filter)
			if !re.MatchString(output) {
				continue
			}
		}

		// invoke generate_view string
	}
}

func DoGenerateIntegrations(stdout, stderr io.Writer, flags *flag.FlagSet, args []string) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	var exclude arrayFlags
	root_model := flags.String("root-model", ".", "sysl root directory for input model file (default: .)")
	title := flags.String("title", "", "diagram title")
	plantuml := flags.String("plantuml", "", strings.Join([]string{"base url of plantuml server",
		"(default: $SYSL_PLANTUML or http://localhost:8080/plantuml",
		"see http://plantuml.com/server.html#install for more info)"}, "\n"))
	output := flags.String("output", "%(epname).png", "output file(default: %(epname).png)")
	project := flags.String("project", "", "project pseudo-app to render")
	filter := flags.String("filter", "", "Only generate diagrams whose output paths match a pattern")
	modules := flags.String("modules", ".", strings.Join([]string{"input files without .sysl extension and with leading /",
		"eg: /project_dir/my_models",
		"combine with --root if needed"}, "\n"))
	flags.Var(&exclude, "exclude", "apps to exclude")
	clustered := flags.Bool("clustered", false, "group integration components into clusters")
	epa := flags.Bool("epa", false, "produce and EPA integration view")

	// Following variables currently are unused. Keep them to align with the python version.
	expire_cache := flags.Bool("expire-cache", false, "Expire cache entries to force checking against real destination(default: false)")
	dry_run := flags.Bool("dry-run", false, "Don't perform confluence uploads, but show what would have happened(default: false)")
	verbose := flags.Bool("verbose", false, "Report each output(default: false)")

	err := flags.Parse(args[1:])
	if err != nil {
		log.Errorf("arguments parse error: %v", err)
	}
	log.Debugf("root_model: %s\n", *root_model)
	log.Debugf("project: %v\n", project)
	log.Debugf("clustered: %t\n", *clustered)
	log.Debugf("epa: %t\n", *epa)
	log.Debugf("title: %s\n", *title)
	log.Debugf("plantuml: %s\n", *plantuml)
	log.Debugf("verbose: %t\n", *verbose)
	log.Debugf("expire_cache: %t\n", *expire_cache)
	log.Debugf("dry_run: %t\n", *dry_run)
	log.Debugf("filter: %s\n", *filter)
	log.Debugf("modules: %s\n", *modules)
	log.Debugf("output: %s\n", *output)

	GenerateIntegrations(*root_model, *title, *plantuml, *output, *project, *filter, *modules, exclude, *clustered, *epa)
}
