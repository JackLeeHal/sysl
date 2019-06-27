package utils

import (
	"strings"

	"github.com/anz-bank/sysl/src/proto"
)

func GetAppName(appname *sysl.AppName) string {
	return strings.Join(appname.Part, " :: ")
}

func GetApp(appName *sysl.AppName, mod *sysl.Module) *sysl.Application {
	return mod.Apps[GetAppName(appName)]
}

func HasAbstractPattern(attrs map[string]*sysl.Attribute) bool {
	patterns, has := attrs["patterns"]
	if has {
		if x := patterns.GetA(); x != nil {
			for _, y := range x.Elt {
				if y.GetS() == "abstract" {
					return true
				}
			}
		}
	}
	return false
}

func IsSameApp(a *sysl.AppName, b *sysl.AppName) bool {
	if len(a.Part) != len(b.Part) {
		return false
	}
	for i := range a.Part {
		if a.Part[i] != b.Part[i] {
			return false
		}
	}
	return true
}

func IsSameCall(a *sysl.Call, b *sysl.Call) bool {
	return IsSameApp(a.Target, b.Target) && a.Endpoint == b.Endpoint
}
