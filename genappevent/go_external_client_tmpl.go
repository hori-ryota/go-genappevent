package genappevent

import (
	"strings"
	"text/template"
)

var GoExternalClientTmpl = TemplateRenderer{
	Tmpl: template.Must(template.New("GoExternalClientTmpl").Funcs(map[string]interface{}{
		"FmtImports":   FmtImports,
		"ToUpperCamel": ToUpperCamel,
		"TypeConverter": func(paramType string) string {
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
				return paramType
			default:
				return strings.Replace(paramType, typ, "string", -1)
			}
		},
	}).Parse(`
// Code generated ; DO NOT EDIT

package appevent

{{ FmtImports .ImportPackages }}

{{ range .Events }}
	type {{ .Name }}Event struct {
		occurredOn time.Time
		{{- range .Params }}
		{{ .Name }} {{ TypeConverter .Type }}
		{{- end }}
	}

	func New{{ .Name }}Event(
			occurredOn time.Time,
		{{- range .Params }}
		{{ .Name }} {{ TypeConverter .Type }},
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
	func (e {{ $EventName }}Event) {{ ToUpperCamel .Name }}() {{ TypeConverter .Type }} {
		return e.{{ .Name }}
	}
	{{- end }}

	type {{ .Name }}EventPublisher interface {
		Publish{{ .Name }}Event(context.Context, {{ .Name }}Event) <-chan error
	}

	type {{ .Name }}EventSubscriber interface {
		Subscribe{{ .Name }}Event(context.Context) (<-chan {{ .Name }}Event, error)
	}

	type {{ .Name }}EventPublishers struct {
		publishers []{{ .Name }}EventPublisher
		bufSize int
	}

	func New{{ .Name }}EventPublishers(
		publishers []{{ .Name }}EventPublisher,
		bufSize int,
	) {{ .Name }}EventPublishers {
		return {{ .Name }}EventPublishers{
			publishers: publishers,
			bufSize: bufSize,
		}
	}

	func (ps {{ .Name }}EventPublishers) Publish{{ .Name }}Event(ctx context.Context, event {{ .Name }}Event) <-chan error {
		wg := new(sync.WaitGroup)
		c := make(chan error, ps.bufSize)
		go func() {
			defer close(c)
			wg.Wait()
		}()
		for i := range ps.publishers {
			wg.Add(1)
			go func() {
				wg.Done()
				for err := range ps.publishers[i].Publish{{ .Name }}Event(ctx, event) {
					c <- err
				}
			}()
		}
		return c
	}

{{- end }}
`)),
	AppendImportPackages: func(paramType string, recommendedImportPackages ...string) []string {
		return []string{"context", "sync", "time"}
	},
}
