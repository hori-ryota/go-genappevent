package genappevent

import (
	"path"
	"strings"
	"text/template"
	"unicode"
)

var GoApplicationTmpl = TemplateRenderer{
	Tmpl: template.Must(template.New("GoApplicationTmpl").Funcs(map[string]interface{}{
		"FmtImports":   FmtImports,
		"ToUpperCamel": ToUpperCamel,
	}).Parse(`
// Code generated ; DO NOT EDIT

package {{ .PackageName }}

{{ FmtImports .ImportPackages }}

{{ range .Events }}
	type {{ .Name }}Event struct {
		occurredOn time.Time
		{{- range .Params }}
		{{ .Name }} {{ .Type }}
		{{- end }}
	}

	func New{{ .Name }}Event(
			occurredOn time.Time,
		{{- range .Params }}
		{{ .Name }} {{ .Type }},
		{{- end }}
	) {{ .Name }}Event {
		return {{ .Name }}Event{
			occurredOn: occurredOn,
			{{- range .Params }}
			{{ .Name }}: {{ .Name }},
			{{- end }}
		}
	}

	func (e {{ .Name }}Event) EventName() string {
		return "{{ .Name }}"
	}

	func (e {{ .Name }}Event) OccurredOn() time.Time {
		return e.occurredOn
	}

	{{ $EventName := .Name }}
	{{- range .Params }}
	func (e {{ $EventName }}Event) {{ ToUpperCamel .Name }}() {{ .Type }} {
		return e.{{ .Name }}
	}
	{{- end }}

	type {{ .Name }}EventPublisher interface {
		Publish{{ .Name }}Event(context.Context, {{ .Name }}Event) <-chan domain.Error
	}

	type {{ .Name }}EventSubscriber interface {
		Subscribe{{ .Name }}Event(context.Context) (<-chan {{ .Name }}Event, domain.Error)
	}
{{- end }}
`)),
	ResolveImportPackages: func(paramType string, recommendedImportPackages ...string) []string {
		results := make([]string, 0, 4)
		results = append(results, "context", "time")

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
