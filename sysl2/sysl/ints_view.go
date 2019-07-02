package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/anz-bank/sysl/src/proto"
)

type IntsParam struct {
	apps         []string
	highlights   map[string]struct{}
	integrations []*AppDependency
	app          *sysl.Application
	endpt        *sysl.Endpoint
}

type Args struct {
	title     string
	project   string
	clustered bool
	epa       bool
}

type IntsDiagramVisitor struct {
	mod           *sysl.Module
	stringBuilder *strings.Builder
	highlights    map[string]struct{}
	symbols       map[string]*_var
	topSymbols    map[string]*_topVar
	project       string
}

type _topVar struct {
	topLabel string
	topAlias string
}

type CallsDrawn struct {
	Self   string
	Target string
}

type viewParams struct {
	restrictBy         string
	endptAttrs         map[string]*sysl.Attribute
	highLightColor     string
	arrowColor         string
	indirectArrowColor string
	diagramTitle       string
}

func MakeIntsParam(apps []string, highlights map[string]struct{}, dependencies []*AppDependency, app *sysl.Application, endpt *sysl.Endpoint) *IntsParam {
	return &IntsParam{apps, highlights, dependencies, app, endpt}
}

func MakeArgs(title, project string, clustered, epa bool) *Args {
	return &Args{title, project, clustered, epa}
}

func MakeIntsDiagramVisitor(mod *sysl.Module, stringBuilder *strings.Builder, highlights map[string]struct{}, project string) *IntsDiagramVisitor {
	return &IntsDiagramVisitor{
		mod:           mod,
		stringBuilder: stringBuilder,
		highlights:    highlights,
		symbols:       map[string]*_var{},
		topSymbols:    map[string]*_topVar{},
		project:       project,
	}
}

func (v *IntsDiagramVisitor) VarManagerForComponent(appName string, nameMap map[string]string) string {
	if key, ok := nameMap[appName]; ok {
		appName = key
	}
	if s, ok := v.symbols[appName]; ok {
		return s.alias
	}

	i := len(v.symbols)
	alias := fmt.Sprintf("_%d", i)
	attrs := getAttrs(v.mod, appName)
	attrs["appname"] = appName
	fp := MakeFormatParser(v.mod.Apps[v.project].GetAttrs()["appfmt"].GetS())
	label := fp.Parse(attrs)

	s := &_var{
		label: label,
		alias: alias,
	}
	v.symbols[appName] = s
	r := fmt.Sprintf("[%s] as %s", label, alias)
	if _, ok := v.highlights[appName]; ok {
		r += " <<highlight>>"
	}
	fmt.Fprintln(v.stringBuilder, r)
	return s.alias
}

func (v *IntsDiagramVisitor) VarManagerForTopState(appName string) string {
	var alias, label string
	attrs := map[string]string{}
	if ts, ok := v.topSymbols[appName]; ok {
		return ts.topAlias
	}
	i := len(v.topSymbols)
	alias = fmt.Sprintf("_%d", i)

	attrs = getAttrs(v.mod, appName)
	attrs["appname"] = appName
	fp := MakeFormatParser(v.mod.Apps[v.project].GetAttrs()["appfmt"].GetS())
	label = fp.Parse(attrs)

	ts := &_topVar{
		topLabel: label,
		topAlias: alias,
	}
	v.topSymbols[appName] = ts
	r := ""
	if _, ok := v.highlights[appName]; ok {
		r = fmt.Sprintf("state \"%s\" as X%s <<highlight>> {", label, alias)
	} else {
		r = fmt.Sprintf("state \"%s\" as X%s {", label, alias)
	}
	fmt.Fprintln(v.stringBuilder, r)

	return ts.topAlias
}

func (v *IntsDiagramVisitor) VarManagerForState(name string) string {
	var appName, alias, label string
	attrs := map[string]string{}

	appName = strings.Split(name, " : ")[0]
	epName := strings.Split(name, " : ")[1]

	if s, ok := v.symbols[name]; ok {
		return s.alias
	}
	i := len(v.symbols)
	alias = fmt.Sprintf("_%d", i)

	if v.mod.Apps[appName].Endpoints[epName] != nil {
		for k, v := range v.mod.Apps[appName].Endpoints[epName].Attrs {
			attrs["@"+k] = v.GetS()
		}
	}
	attrs["appname"] = epName
	fp := MakeFormatParser(v.mod.Apps[v.project].GetAttrs()["appfmt"].GetS())
	label = fp.Parse(attrs)

	s := &_var{
		label: label,
		alias: alias,
	}
	v.symbols[name] = s
	r := fmt.Sprintf("  state \"%s\" as %s", label, alias)
	if _, ok := v.highlights[name]; ok {
		r += " <<highlight>>"
	}
	fmt.Fprintln(v.stringBuilder, r)
	return s.alias
}

func (v *IntsDiagramVisitor) buildClusterForStateView(deps []*AppDependency, restrictBy string) {
	clusters := map[string][]string{}
	for _, dep := range deps {
		appA := dep.Self.Name
		appB := dep.Target.Name
		epA := dep.Self.Endpoint
		epB := dep.Target.Endpoint
		_, okA := v.mod.Apps[appA].Endpoints[epA].Attrs[restrictBy]
		_, okB := v.mod.Apps[appB].Endpoints[epB].Attrs[restrictBy]
		if _, ok := v.mod.Apps[appA].Attrs[restrictBy]; !ok && restrictBy != "" {
			if _, ok := v.mod.Apps[appB].Attrs[restrictBy]; !ok {
				continue
			}
		}
		if !okA && restrictBy != "" && !okB {
			continue
		}
		clusters[appA] = append(clusters[appA], epA)
		if appA != appB && !v.mod.Apps[appA].Endpoints[epA].IsPubsub {
			clusters[appA] = append(clusters[appA], epB+" client")
		}
		clusters[appB] = append(clusters[appB], epB)
	}

	var keys []string
	for k := range clusters {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v.VarManagerForTopState(k)
		strSet := MakeStrSet(clusters[k]...)
		for _, m := range strSet.ToSortedSlice() {
			v.VarManagerForState(k + " : " + m)
		}
		fmt.Fprintln(v.stringBuilder, "}")
	}
}

func (v *IntsDiagramVisitor) buildClusterForComponentView(apps []string) map[string]string {
	nameMap := map[string]string{}
	clusters := map[string][]string{}
	for _, v := range apps {
		cluster := strings.Split(v, " :: ")
		if len(cluster) > 1 {
			clusters[cluster[0]] = append(clusters[cluster[0]], v)
		}
	}

	for k, v := range clusters {
		if len(v) <= 1 {
			delete(clusters, k)
		}
		for _, s := range v {
			nameMap[s] = strings.Split(s, " :: ")[1]
		}
	}

	for k, apps := range clusters {
		fmt.Fprintf(v.stringBuilder, "package \"%s\" {", k)
		fmt.Fprintln(v.stringBuilder)
		for _, n := range apps {
			v.VarManagerForComponent(n, nameMap)
		}
		fmt.Fprintln(v.stringBuilder, "}")
	}

	return nameMap
}

func (v *IntsDiagramVisitor) generateStateView(args *Args, viewParams viewParams, params *IntsParam) string {

	fmt.Fprintln(v.stringBuilder, "@startuml")
	if viewParams.diagramTitle != "" {
		fmt.Fprintln(v.stringBuilder, "title "+viewParams.diagramTitle)
	}
	fmt.Fprintln(v.stringBuilder, "left to right direction")
	fmt.Fprintln(v.stringBuilder, "scale max 16384 height")
	fmt.Fprintln(v.stringBuilder, "hide empty description")
	fmt.Fprintln(v.stringBuilder, "skinparam state {")
	fmt.Fprintln(v.stringBuilder, "  BackgroundColor FloralWhite")
	fmt.Fprintln(v.stringBuilder, "  BorderColor Black")
	fmt.Fprintln(v.stringBuilder, "  ArrowColor Crimson")
	if viewParams.highLightColor != "" {
		fmt.Fprintln(v.stringBuilder, "  BackgroundColor<<highlight>> "+viewParams.highLightColor)
	}
	if viewParams.arrowColor != "" {
		fmt.Fprintln(v.stringBuilder, "  ArrowColor "+viewParams.arrowColor)
	}

	if viewParams.indirectArrowColor != "" && viewParams.indirectArrowColor != "none" {
		fmt.Fprintln(v.stringBuilder, "  ArrowColor<<indirect>> "+viewParams.indirectArrowColor)
		fmt.Fprintln(v.stringBuilder, "  ArrowColor<<internal>> "+viewParams.indirectArrowColor)
	}
	fmt.Fprintln(v.stringBuilder, "}")

	v.buildClusterForStateView(params.integrations, viewParams.restrictBy)
	var processed []string
	for _, dep := range params.integrations {
		appA := dep.Self.Name
		appB := dep.Target.Name
		epA := dep.Self.Endpoint
		epB := dep.Target.Endpoint
		_, restrictByAppA := v.mod.Apps[appA].Attrs[viewParams.restrictBy]
		_, restrictByAppB := v.mod.Apps[appB].Attrs[viewParams.restrictBy]
		_, restrictByEpA := v.mod.Apps[appA].Endpoints[epA].Attrs[viewParams.restrictBy]
		_, restrictByEpB := v.mod.Apps[appB].Endpoints[epB].Attrs[viewParams.restrictBy]
		if viewParams.restrictBy != "" && !restrictByAppA && !restrictByAppB {
			continue
		}
		if viewParams.restrictBy != "" && !restrictByEpA && !restrictByEpB {
			continue
		}
		matchApp := appB
		matchEp := epB
		label := ""
		needsInt := appA != matchApp

		var pubSubSrcPtrns []string
		if v.mod.Apps[appA].Endpoints[epA].Attrs["patterns"] != nil {
			for _, v := range v.mod.Apps[appA].Endpoints[epA].Attrs["patterns"].GetA().Elt {
				pubSubSrcPtrns = append(pubSubSrcPtrns, v.GetS())
			}
		}

		targetPatterns := MakeStrSet()
		if v.mod.Apps[matchApp].Endpoints[matchEp].Attrs["patterns"] != nil {
			for _, v := range v.mod.Apps[matchApp].Endpoints[matchEp].Attrs["patterns"].GetA().Elt {
				targetPatterns.Insert(v.GetS())
			}
		} else {
			if v.mod.Apps[matchApp].Attrs["patterns"] != nil {
				for _, v := range v.mod.Apps[matchApp].Attrs["patterns"].GetA().Elt {
					targetPatterns.Insert(v.GetS())
				}
			}
		}
		attrs := map[string]string{}
		for k, v := range dep.Statement.GetAttrs() {
			attrs["@"+k] = v.GetS()
		}
		var srcPtrns []string
		if dep.Statement.GetAttrs()["patterns"] != nil {
			for _, v := range dep.Statement.GetAttrs()["patterns"].GetA().Elt {
				srcPtrns = append(srcPtrns, v.GetS())
			}
		} else {
			srcPtrns = pubSubSrcPtrns
		}
		var ptrns string
		if srcPtrns != nil || targetPatterns != nil {
			ptrns = strings.Join(srcPtrns, ", ") + " → " + strings.Join(targetPatterns.ToSlice(), ", ")
		} else {
			ptrns = ""
		}
		attrs["patterns"] = ptrns
		if needsInt {
			attrs["needs_int"] = strconv.FormatBool(needsInt)
		}
		fp := MakeFormatParser(params.app.Attrs["epfmt"].GetS())
		label = fp.Parse(attrs)

		flow := strings.Join([]string{appA, epB, appB, epB}, ".")
		isPubSub := v.mod.Apps[appA].Endpoints[epA].GetIsPubsub()
		epBClient := epB + " client"

		if appA != appB {
			if isPubSub {
				if label != "" {
					label = " : " + label
				}
				fmt.Fprintf(v.stringBuilder, "%s -%s> %s%s", v.VarManagerForState(appA+" : "+epA), "[#blue]", v.VarManagerForState(appB+" : "+epB), label)
				fmt.Fprintln(v.stringBuilder)
			} else {
				color := ""
				if viewParams.indirectArrowColor == "" {
					color = "[#silver]-"
				} else {
					color = "[#" + viewParams.indirectArrowColor + "]-"
				}
				fmt.Fprintf(v.stringBuilder, "%s -%s> %s", v.VarManagerForState(appA+" : "+epA), color, v.VarManagerForState(appA+" : "+epBClient))
				fmt.Fprintln(v.stringBuilder)
				if !stringInSlice(flow, processed) {
					fmt.Fprintf(v.stringBuilder, "%s -%s> %s : %s", v.VarManagerForState(appA+" : "+epBClient), "[#black]", v.VarManagerForState(appB+" : "+epB), label)
					fmt.Fprintln(v.stringBuilder)
					processed = append(processed, flow)
				}
			}
		} else {
			color := ""
			if viewParams.indirectArrowColor == "" {
				color = "[#silver]-"
			} else {
				color = "[#" + viewParams.indirectArrowColor + "]-"
			}
			fmt.Fprintf(v.stringBuilder, "%s -%s> %s%s", v.VarManagerForState(appA+" : "+epA), color, v.VarManagerForState(appB+" : "+epB), label)
			fmt.Fprintln(v.stringBuilder)
		}
	}
	fmt.Fprintln(v.stringBuilder, "@enduml")
	return v.stringBuilder.String()

}

func (v *IntsDiagramVisitor) drawComponentView(viewParams viewParams, params *IntsParam, nameMap map[string]string) {
	callsDrawn := map[CallsDrawn]struct{}{}
	if viewParams.endptAttrs["view"].GetS() == "system" {
		v.drawSystemView(viewParams, params, nameMap)
	} else {
		for _, dep := range params.integrations {
			appA := dep.Self.Name
			appB := dep.Target.Name
			apps := &CallsDrawn{
				Self:   appA,
				Target: appB,
			}
			var direct []string
			if _, ok := params.highlights[appA]; ok {
				direct = append(direct, appA)
			}
			if _, ok := params.highlights[appB]; ok {
				direct = append(direct, appB)
			}
			_, ok := callsDrawn[*apps]
			if appA != appB && !ok {
				if direct != nil || viewParams.indirectArrowColor != "none" {
					indirect := ""
					if direct == nil {
						indirect = " <<indirect>>"
					}
					fmt.Fprintf(v.stringBuilder, "%s --> %s%s", v.VarManagerForComponent(appA, nameMap), v.VarManagerForComponent(appB, nameMap), indirect)
					fmt.Fprintln(v.stringBuilder)
					callsDrawn[*apps] = struct{}{}
				}
			}
		}
		for _, app := range params.apps {
			for _, mixin := range v.mod.Apps[app].GetMixin2() {
				mixinName := strings.Join(mixin.Name.Part, " :: ")
				fmt.Fprintf(v.stringBuilder, "%s <|.. %s", v.VarManagerForComponent(mixinName, nameMap), v.VarManagerForComponent(app, nameMap))
				fmt.Fprintln(v.stringBuilder)
			}
		}
	}
}

func (v *IntsDiagramVisitor) drawSystemView(viewParams viewParams, params *IntsParam, nameMap map[string]string) {
	callsDrawn := map[CallsDrawn]struct{}{}
	for _, dep := range params.integrations {
		appA := dep.Self.Name
		appB := dep.Target.Name
		apps := &CallsDrawn{
			Self:   appA,
			Target: appB,
		}
		var direct []string
		if _, ok := params.highlights[appA]; ok {
			direct = append(direct, appA)
		}
		if _, ok := params.highlights[appB]; ok {
			direct = append(direct, appB)
		}
		appA = strings.Split(appA, " :: ")[0]
		appB = strings.Split(appB, " :: ")[0]
		_, ok := callsDrawn[*apps]
		if appA != appB && !ok {
			if direct != nil || viewParams.indirectArrowColor != "none" {
				indirect := ""
				if direct == nil {
					indirect = " <<indirect>>"
				}
				fmt.Fprintf(v.stringBuilder, "%s --> %s%s", v.VarManagerForComponent(appA, nameMap), v.VarManagerForComponent(appB, nameMap), indirect)
				fmt.Fprintln(v.stringBuilder)
				callsDrawn[*apps] = struct{}{}
			}
		}
	}
}

func (v *IntsDiagramVisitor) generateComponentView(args *Args, viewParams viewParams, params *IntsParam) string {

	fmt.Fprintln(v.stringBuilder, "@startuml")
	if viewParams.diagramTitle != "" {
		fmt.Fprintln(v.stringBuilder, "title "+viewParams.diagramTitle)
	}
	fmt.Fprintln(v.stringBuilder, "hide stereotype")
	fmt.Fprintln(v.stringBuilder, "scale max 16384 height")
	fmt.Fprintln(v.stringBuilder, "skinparam component {")
	fmt.Fprintln(v.stringBuilder, "  BackgroundColor FloralWhite")
	fmt.Fprintln(v.stringBuilder, "  BorderColor Black")
	fmt.Fprintln(v.stringBuilder, "  ArrowColor Crimson")
	if viewParams.highLightColor != "" {
		fmt.Fprintln(v.stringBuilder, "  BackgroundColor<<highlight>> "+viewParams.highLightColor)
	}
	if viewParams.arrowColor != "" {
		fmt.Fprintln(v.stringBuilder, "  ArrowColor "+viewParams.arrowColor)
	}

	if viewParams.indirectArrowColor != "" && viewParams.indirectArrowColor != "none" {
		fmt.Fprintln(v.stringBuilder, "  ArrowColor<<indirect>> "+viewParams.indirectArrowColor)
	}
	fmt.Fprintln(v.stringBuilder, "}")

	nameMap := map[string]string{}
	if args.clustered || viewParams.endptAttrs["view"].GetS() == "clustered" {
		nameMap = v.buildClusterForComponentView(params.apps)
	}
	v.drawComponentView(viewParams, params, nameMap)
	fmt.Fprintln(v.stringBuilder, "@enduml")
	return v.stringBuilder.String()
}

func GenerateView(args *Args, params *IntsParam, mod *sysl.Module) string {
	var stringBuilder strings.Builder
	var titleParser *FormatParser
	v := MakeIntsDiagramVisitor(mod, &stringBuilder, params.highlights, args.project)
	restrictBy := ""
	if params.endpt.Attrs["restrict_by"] != nil {
		restrictBy = params.endpt.Attrs["restrict_by"].GetS()
	}

	appAttrs := params.app.Attrs
	endptAttrs := params.endpt.Attrs
	highLightColor := appAttrs["highlight_color"].GetS()
	arrowColor := appAttrs["arrow_color"].GetS()
	indirectArrowColor := appAttrs["indirect_arrow_color"].GetS()

	attrs := map[string]string{
		"epname":     params.endpt.Name,
		"eplongname": params.endpt.LongName,
	}
	if appAttrs["title"].GetS() != "" {
		titleParser = MakeFormatParser(appAttrs["title"].GetS())
	} else {
		titleParser = MakeFormatParser(args.title)
	}
	diagramTitle := titleParser.Parse(attrs)

	viewParams := &viewParams{
		restrictBy:         restrictBy,
		endptAttrs:         endptAttrs,
		highLightColor:     highLightColor,
		arrowColor:         arrowColor,
		indirectArrowColor: indirectArrowColor,
		diagramTitle:       diagramTitle,
	}

	fmt.Fprintln(&stringBuilder, "''''''''''''''''''''''''''''''''''''''''''")
	fmt.Fprintln(&stringBuilder, "''                                      ''")
	fmt.Fprintln(&stringBuilder, "''  AUTOGENERATED CODE -- DO NOT EDIT!  ''")
	fmt.Fprintln(&stringBuilder, "''                                      ''")
	fmt.Fprintln(&stringBuilder, "''''''''''''''''''''''''''''''''''''''''''")
	fmt.Fprintln(&stringBuilder)

	if args.epa || endptAttrs["view"].GetS() == "epa" {
		return v.generateStateView(args, *viewParams, params)
	} else {
		return v.generateComponentView(args, *viewParams, params)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getAttrs(m *sysl.Module, appName string) map[string]string {
	val := map[string]string{}
	if app, ok := m.Apps[appName]; ok {
		attrs := app.Attrs
		for k, v := range attrs {
			val["@"+k] = v.GetS()
		}
		return val
	}
	return val
}
