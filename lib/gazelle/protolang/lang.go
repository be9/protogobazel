package protolang

import (
	"fmt"
	"slices"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

const (
	protoLibraryRule = "proto_library"
)

type customProtoLang struct {
	language.BaseLang

	libProtoExtras map[string]string
}

func (*customProtoLang) Name() string { return "custom_proto" }

// NewLanguage instantiates the "custom_proto" language.
func NewLanguage() language.Language {
	return &customProtoLang{
		libProtoExtras: make(map[string]string),
	}
}

var kinds = map[string]rule.KindInfo{
	//tsProtoLibraryRule: {
	//	MatchAttrs:    []string{"srcs"},
	//	NonEmptyAttrs: map[string]bool{"srcs": true},
	//	MergeableAttrs: map[string]bool{
	//		"srcs":         true,
	//		"has_services": true,
	//	},
	//	ResolveAttrs: map[string]bool{"deps": true},
	//},
}

func (*customProtoLang) Kinds() map[string]rule.KindInfo { return kinds }

func (*customProtoLang) ApparentLoads(func(string) string) []rule.LoadInfo {
	return []rule.LoadInfo{
		//{
		//	Name: "//build/bzl:proto.bzl",
		//	Symbols: []string{
		//		"ts_proto_library",
		//	},
		//},
	}
}

func (lang *customProtoLang) GenerateRules(args language.GenerateArgs) (result language.GenerateResult) {
	const (
		libProtoPrefix = "lib/proto/"
		protoSrcPrefix = "proto/src/"
	)

	// Populate lang.libProtoExtras; any subdir under lib/proto/ that has any extra sources (not *.pb.go),
	// will have a go_library rule which is recorded into the map.
	if strings.HasPrefix(args.Rel, libProtoPrefix) {
		goLibraryIdx := slices.IndexFunc(args.OtherGen, func(r *rule.Rule) bool {
			return r.Kind() == "go_library"
		})

		if goLibraryIdx >= 0 {
			r := args.OtherGen[goLibraryIdx]
			importPath := r.AttrString("importpath")
			lang.libProtoExtras[importPath] = fmt.Sprintf("//%s:%s", args.Rel, r.Name())
		}
	}

	// Since lib/proto precedes proto/src, lang.libProtoExtras will be populated
	// before the first subdir under proto/src/ will be analyzed.
	if strings.HasPrefix(args.Rel, protoSrcPrefix) && args.File != nil {
		for _, r := range args.File.Rules {
			if r.Kind() == "go_library" {

				embed := r.AttrStrings("embed")
				if libProtoLib, ok := lang.libProtoExtras[r.AttrString("importpath")]; ok {
					if slices.IndexFunc(embed, func(s string) bool { return s == libProtoLib }) < 0 {
						r.SetAttr("embed", append(embed, libProtoLib))
					}
				} else {
					origLen := len(embed)
					embed = slices.DeleteFunc(embed, func(s string) bool {
						return strings.HasPrefix(s, "//"+libProtoPrefix)
					})
					if len(embed) != origLen {
						r.SetAttr("embed", embed)
					}
				}
			}
		}
	}
	return
}

//func getProtoLibraries(args language.GenerateArgs) ([]*rule.Rule, []*rule.Rule) {
//	protos := make([]*rule.Rule, 0, len(args.OtherGen))
//	emptyProtos := make([]*rule.Rule, 0)
//
//	for _, r := range args.OtherGen {
//		if r.Kind() == protoLibraryRule {
//			protos = append(protos, r)
//		}
//	}
//
//	for _, r := range args.OtherEmpty {
//		if r.Kind() == protoLibraryRule {
//			emptyProtos = append(emptyProtos, r)
//		}
//	}
//
//	return protos, emptyProtos
//}

//func addTsProtoRule(args language.GenerateArgs, protoLibrary *rule.Rule, ruleName string, result *language.GenerateResult) {
//	protoRuleLabel := label.New("", args.Rel, protoLibrary.Name())
//	protoRuleLabelStr := protoRuleLabel.Rel("", args.Rel)
//
//	r := rule.NewRule(tsProtoLibraryRule, ruleName)
//	r.SetAttr("proto", protoRuleLabelStr.String())
//
//	// this should work in `gazelle:proto file` mode where proto source is 1:1 mapped to proto_library name
//	protoName := protoLibrary.Name()
//	if !strings.HasSuffix(protoName, "_proto") {
//		log.Printf("cannot obtain source for proto_library: %v", protoName)
//		return
//	}
//
//	protoPath := filepath.Join(args.Dir, protoName[:len(protoName)-6]+".proto")
//	protoInfo, err := analyzeProtoFile(protoPath)
//	if err != nil {
//		log.Printf("cannot parse proto: %v", err)
//	}
//
//	if !protoInfo.hasMessages && !protoInfo.hasEnums && !protoInfo.hasServices {
//		return
//	}
//
//	if protoInfo.hasServices {
//		r.SetAttr("has_services", true)
//	}
//
//	result.Gen = append(result.Gen, r)
//	result.Imports = append(result.Imports, nil)
//}
