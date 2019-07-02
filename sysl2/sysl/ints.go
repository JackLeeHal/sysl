package main

import (
	"flag"
	"io"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

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
	ds := NewDependencySet()
	ds.CollectAppDependencies(mod)

	// The "project" app that specifies the required view of the integration
	app := mod.GetApps()[project]
	of := MakeFormatParser(output)
	arr := ds.ToSlice()
	// Interate over each endpoint within the selected project
	for epname, endpt := range app.GetEndpoints() {
		// build the set of excluded items
		excludes := MakeStrSetFromSpecificAttr("exclude", endpt.GetAttrs())
		passthroughs := MakeStrSetFromSpecificAttr("passthrough", endpt.GetAttrs())
		// endpt.stmt's "action" will conatain the "apps" whose integration is to be drawn
		// each one of these will be placed into the "integration" list
		integrations := MakeStrSetFromActionStatement(endpt.GetStmt())

		highlights := FindApps(mod, excludeStrSet, integrations, arr, true)
		apps := FindApps(mod, excludeStrSet, highlights, arr, false)
		apps = apps.Difference(excludes)
		apps = apps.Difference(passthroughs)
		output_dir := of.FmtOutput(project, epname, endpt.GetLongName(), endpt.GetAttrs())

		if filter != "" {
			re := regexp.MustCompile(filter)
			if !re.MatchString(output) {
				continue
			}
		}

		// invoke generate_view string
		dependencySet := ds.FindIntegrations(apps, excludes, passthroughs, mod)
		intsParam := MakeIntsParam(apps.ToSlice(), highlights, dependencySet.ToSlice(), app, endpt)
		args := MakeArgs(title, project, clustered, epa)
		r[output_dir] = GenerateView(args, intsParam, mod)
	}

	return r
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
