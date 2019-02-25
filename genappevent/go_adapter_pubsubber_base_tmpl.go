package genappevent

import (
	"fmt"
	"path"
	"strings"
	"text/template"
	"unicode"
)

var GoAdaptorPubsubberBaseTmpl = TemplateRenderer{
	Tmpl: template.Must(template.New("GoAdaptorPubsubberBaseTmpl").Funcs(map[string]interface{}{
		"FmtImports":   FmtImports,
		"ToUpperCamel": ToUpperCamel,
		"ToLowerCamel": ToLowerCamel,
		"ToApplicationType": func(paranName, paramType string) string {
			typ := paramType
			typ = strings.TrimPrefix(typ, "[]")
			typ = strings.TrimPrefix(typ, "*")
			switch typ {
			case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16",
				"uint32", "uint64", "uintptr",
				"byte", "rune",
				"float32", "float64",
				"complex64",
				"string",
				"time.Time":
				return fmt.Sprintf("event.%s()", ToUpperCamel(paranName))
			default:
				return fmt.Sprintf("%s(event.%s())", paramType, ToUpperCamel(paranName))
			}
		},
		"ToAppeventType": func(paranName, paramType string) string {
			typ := paramType
			typ = strings.TrimPrefix(typ, "[]")
			typ = strings.TrimPrefix(typ, "*")
			switch typ {
			case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16",
				"uint32", "uint64", "uintptr",
				"byte", "rune",
				"float32", "float64",
				"complex64",
				"string",
				"time.Time":
				return fmt.Sprintf("event.%s()", ToUpperCamel(paranName))
			default:
				return fmt.Sprintf("string(event.%s())", ToUpperCamel(paranName))
			}
		},
	}).Parse(`
// Code generated ; DO NOT EDIT

package appevent

{{ FmtImports .ImportPackages }}

type EventPublisher struct {
	{{- range .Events }}
	{{ .Name }}EventPublisher
	{{- end }}
}

func NewEventPublisher(
	{{- range .Events }}
	{{ ToLowerCamel .Name }}EventPublisher {{ .Name }}EventPublisher,
	{{- end }}
) EventPublisher {
	return EventPublisher{
		{{- range .Events }}
		{{ .Name }}EventPublisher: {{ ToLowerCamel .Name }}EventPublisher,
		{{- end }}
	}
}

{{ range .Events }}

	type {{ .Name }}EventPublisher struct {
		publisher appevent.{{ .Name }}EventPublisher
		bufSize int
	}

	func New{{ .Name }}EventPublisher(
		publisher appevent.{{ .Name }}EventPublisher,
		bufSize int,
	) {{ .Name }}EventPublisher {
		return {{ .Name }}EventPublisher{
			publisher: publisher,
			bufSize: bufSize,
		}
	}

	func (p {{ .Name }}EventPublisher) Publish{{ .Name }}Event(ctx context.Context, event application.{{ .Name }}Event) <-chan domain.Error {
		c := make(chan domain.Error, p.bufSize)
		go func() {
			defer close(c)
			errSrc := p.publisher.Publish{{ .Name }}Event(
				ctx,
				appevent.New{{ .Name }}Event(
					event.OccurredOn(),
					{{- range .Params }}
					{{ ToAppeventType .Name .Type }},
					{{- end }}
				),
			)
			for err := range errSrc {
				c <- domain.ErrorUnknown(err)
			}
		}()
		return c
	}

	type {{ .Name }}EventSubscriber struct {
		subscriber appevent.{{ .Name }}EventSubscriber
		bufSize int
	}

	func New{{ .Name }}EventSubscriber(
		subscriber appevent.{{ .Name }}EventSubscriber,
		bufSize int,
	) {{ .Name }}EventSubscriber {
		return {{ .Name }}EventSubscriber{
			subscriber: subscriber,
			bufSize: bufSize,
		}
	}

	func (p {{ .Name }}EventSubscriber) Subscribe{{ .Name }}Event(ctx context.Context) (<-chan application.{{ .Name }}Event, domain.Error) {
		src, err := p.subscriber.Subscribe{{ .Name }}Event(ctx)
		if err != nil {
			return nil, domain.ErrorUnknown(err)
		}
		c := make(chan application.{{ .Name }}Event, p.bufSize)
		go func() {
			defer close(c)
			for event := range src {
				c <- application.New{{ .Name }}Event(
					event.OccurredOn(),
					{{- range .Params }}
					{{ ToApplicationType .Name .Type }},
					{{- end }}
				)
			}
		}()
		return c, nil
	}
{{- end }}
`)),
	AppendImportPackages: func(paramType string, recommendedImportPackages ...string) []string {
		results := make([]string, 0, 4)
		results = append(results, "context")

		for _, pkgPath := range recommendedImportPackages {
			if "appevent" == path.Base(pkgPath) || "application" == path.Base(pkgPath) {
				results = append(results, pkgPath)
			}
		}

		if !strings.Contains(paramType, ".") {
			return results
		}
		for _, s := range strings.FieldsFunc(paramType, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '.'
		}) {
			ss := strings.SplitN(s, ".", 2)
			if len(ss) != 2 {
				continue
			}
			for _, pkgPath := range recommendedImportPackages {
				if ss[0] == path.Base(pkgPath) {
					results = append(results, pkgPath)
					break
				}
			}
		}
		return results
	},
}
