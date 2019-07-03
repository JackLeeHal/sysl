package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/anz-bank/sysl/src/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

type IntegrationParam struct {
	m            *sysl.Module
	excludes     StrSet
	passthroughs StrSet
	commonEndpts StrSet
	highlights   StrSet
	apps         StrSet
	deps         map[string]*AppDependency
}

type ApplicationCollectionElement struct {
	apps StrSet
}

type IntegrationStatementElement struct {
	appname string
	epname  string
	stmts   []*sysl.Statement
}

func MakeIntegrationParam(mod *sysl.Module, excludes, passthroughs, highlights StrSet) *IntegrationParam {
	return &IntegrationParam{
		m:            mod,
		excludes:     excludes,
		passthroughs: passthroughs,
		commonEndpts: MakeStrSet(".. * <- *", "*"),
		highlights:   highlights,
		apps:         StrSet{},
		deps:         map[string]*AppDependency{},
	}
}

func MakeApplicationCollectionElement(apps StrSet) *ApplicationCollectionElement {
	return &ApplicationCollectionElement{apps}
}

func MakeIntegrationStatementElement(appname, epname string, stmts []*sysl.Statement) *IntegrationStatementElement {
	return &IntegrationStatementElement{appname, epname, stmts}
}

func (v *IntegrationParam) collectApplicationDependencies(e *ApplicationCollectionElement) {
	for app := range e.apps {
		for epname, endpt := range v.m.GetApps()[app].GetEndpoints() {
			i := MakeIntegrationStatementElement(app, epname, endpt.GetStmt())
			v.handleStatement(i)
		}
	}
}

func (v *IntegrationParam) handleStatement(i *IntegrationStatementElement) {
	for _, stmt := range i.stmts {
		switch t := stmt.GetStmt().(type) {
		case *sysl.Statement_Call:
			targetName := getAppName(t.Call.GetTarget())
			if ok := v.addAppDependency(i.appname, i.epname, targetName, t.Call.Endpoint, stmt); !ok {
				continue
			}
			v.addEndpoint(targetName, t.Call.Endpoint)
		case *sysl.Statement_Action, *sysl.Statement_Ret:
			continue
		case *sysl.Statement_Cond:
			e := MakeIntegrationStatementElement(i.appname, i.epname, t.Cond.GetStmt())
			v.handleStatement(e)
		case *sysl.Statement_Loop:
			e := MakeIntegrationStatementElement(i.appname, i.epname, t.Loop.GetStmt())
			v.handleStatement(e)
		case *sysl.Statement_LoopN:
			e := MakeIntegrationStatementElement(i.appname, i.epname, t.LoopN.GetStmt())
			v.handleStatement(e)
		case *sysl.Statement_Foreach:
			e := MakeIntegrationStatementElement(i.appname, i.epname, t.Foreach.GetStmt())
			v.handleStatement(e)
		case *sysl.Statement_Group:
			e := MakeIntegrationStatementElement(i.appname, i.epname, t.Group.GetStmt())
			v.handleStatement(e)
		case *sysl.Statement_Alt:
			for _, choice := range t.Alt.GetChoice() {
				e := MakeIntegrationStatementElement(i.appname, i.epname, choice.GetStmt())
				v.handleStatement(e)
			}
		default:
			panic("No statement!")
		}
	}
}

func (v *IntegrationParam) addAppDependency(sourceApp, sourceEndpt, targetApp, targetEndpt string, stmt *sysl.Statement) bool {
	if len(v.apps) == 0 {
		v.apps = v.apps.Union(v.highlights)
	}
	apps := MakeStrSet(sourceApp, targetApp)
	excludeApps := apps.Intersection(v.excludes)
	if len(excludeApps) > 0 {
		return false
	}
	validApps := apps.Intersection(v.highlights)
	if len(validApps) > 0 {
		sourceApplication := v.m.GetApps()[sourceApp]
		if !hasPattern("human", sourceApplication.GetAttrs()) {
			v.apps.Insert(sourceApp)
		}
		targetApplication := v.m.GetApps()[targetApp]
		if !hasPattern("human", targetApplication.GetAttrs()) {
			v.apps.Insert(targetApp)
		}
	}
	if !(apps.IsSubSet(v.apps) || (v.apps.Contains(sourceApp) && v.passthroughs.Contains(targetApp)) || v.passthroughs.Contains(sourceApp)) {
		return false
	}
	endpts := MakeStrSet(sourceEndpt, targetEndpt)
	invalidEndpts := endpts.Intersection(v.commonEndpts)
	if len(invalidEndpts) > 0 {
		return false
	}
	k := fmt.Sprintf("%s:%s:%s:%s", sourceApp, sourceEndpt, targetApp, targetEndpt)
	if _, has := v.deps[k]; has {
		return false
	}
	dep := MakeAppDependency(MakeAppElement(sourceApp, sourceEndpt), MakeAppElement(targetApp, targetEndpt), stmt)
	v.deps[k] = dep

	return true
}

func (v *IntegrationParam) addEndpoint(targetApp, targetEndpt string) {
	if targetEndpt != ".. * <- *" {
		for epname, endpt := range v.m.GetApps()[targetApp].GetEndpoints() {
			i := MakeIntegrationStatementElement(targetApp, epname, endpt.GetStmt())
			v.handleStatement(i)
		}
	}
}

func GenerateIntegrations(
	root_model, title, output, project, filter, modules string,
	exclude []string, clustered, epa bool,
) map[string]string {
	r := make(map[string]string)
	mod := loadApp(root_model, modules)

	if len(exclude) == 0 && project != "" {
		exclude = append(exclude, project)
	}
	excludeStrSet := MakeStrSet(exclude...)

	// The "project" app that specifies the required view of the integration
	app := mod.GetApps()[project]
	of := MakeFormatParser(output)
	// Interate over each endpoint within the selected project
	for epname, endpt := range app.GetEndpoints() {
		excludes := MakeStrSetFromSpecificAttr("exclude", endpt.GetAttrs())
		passthroughs := MakeStrSetFromSpecificAttr("passthrough", endpt.GetAttrs())
		integrations := MakeStrSetFromActionStatement(endpt.GetStmt())
		output_dir := of.FmtOutput(project, epname, endpt.GetLongName(), endpt.GetAttrs())

		if filter != "" {
			re := regexp.MustCompile(filter)
			if !re.MatchString(output) {
				continue
			}
		}

		highlightApps := findHighlights(mod, integrations)
		v := MakeIntegrationParam(mod, excludeStrSet.Union(excludes), passthroughs, highlightApps)
		appCollectionElement := MakeApplicationCollectionElement(integrations)
		v.collectApplicationDependencies(appCollectionElement)
		intsParam := MakeIntsParam(v.apps.ToSlice(), v.highlights, v.deps, app, endpt)
		args := MakeArgs(title, project, clustered, epa)
		r[output_dir] = GenerateView(args, intsParam, mod)
	}

	return r
}

func findHighlights(mod *sysl.Module, apps StrSet) StrSet {
	highlights := StrSet{}
	for appName := range apps {
		app := mod.GetApps()[appName]
		if !hasPattern("human", app.GetAttrs()) {
			highlights.Insert(appName)
		}
	}

	return highlights
}

func DoGenerateIntegrations(stdout, stderr io.Writer, flags *flag.FlagSet, args []string) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	ints := kingpin.New("ints", "Generate integrations")
	root := ints.Flag("root", "sysl root directory for input model file (default: .)").Default(".").String()
	title := ints.Flag("title", "diagram title").Short('t').String()
	plantuml := ints.Flag("plantuml", strings.Join([]string{"base url of plantuml server",
		"(default: $SYSL_PLANTUML or http://localhost:8080/plantuml",
		"see http://plantuml.com/server.html#install for more info)"}, "\n")).Short('p').String()
	output := ints.Flag("output", "output file(default: %(epname).png)").Default("%(epname).png").Short('o').String()
	project := ints.Flag("project", "project pseudo-app to render").Short('j').String()
	filter := ints.Flag("filter", "Only generate diagrams whose output paths match a pattern").String()
	modules := ints.Arg("modules", strings.Join([]string{"input files without .sysl extension and with leading /",
		"eg: /project_dir/my_models",
		"combine with --root if needed"}, "\n")).String()
	exclude := ints.Flag("exclude", "apps to exclude").Short('e').Strings()
	clustered := ints.Flag("clustered", "group integration components into clusters").Short('c').Default("false").Bool()
	epa := ints.Flag("epa", "produce and EPA integration view").Default("false").Bool()
	loglevel := ints.Flag("log", "log level[debug,info,warn,off]").Default("warn").String()

	_, err := ints.Parse(args[1:])
	if err != nil {
		log.Errorf("arguments parse error: %v", err)
	}
	if level, has := defaultLevel[*loglevel]; has {
		log.SetLevel(level)
	}
	if *plantuml == "" {
		*plantuml = os.Getenv("SYSL_PLANTUML")
		if *plantuml == "" {
			*plantuml = "http://localhost:8080/plantuml"
		}
	}
	log.Debugf("root: %s\n", *root)
	log.Debugf("project: %v\n", project)
	log.Debugf("clustered: %t\n", *clustered)
	log.Debugf("exclude: %s\n", *exclude)
	log.Debugf("epa: %t\n", *epa)
	log.Debugf("title: %s\n", *title)
	log.Debugf("plantuml: %s\n", *plantuml)
	log.Debugf("filter: %s\n", *filter)
	log.Debugf("modules: %s\n", *modules)
	log.Debugf("output: %s\n", *output)
	log.Debugf("loglevel: %s\n", *loglevel)

	r := GenerateIntegrations(*root, *title, *output, *project, *filter, *modules, *exclude, *clustered, *epa)
	for k, v := range r {
		OutputPlantuml(k, *plantuml, v)
	}
}
