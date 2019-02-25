package genappevent

import (
	"bytes"
	"go/format"
	"io"
	"text/template"
)

type TemplateRenderer struct {
	Tmpl                 *template.Template
	AppendImportPackages func(paramType string, recommendedImportPackages ...string) []string
}

func (r TemplateRenderer) RendererFunc(w io.Writer, recommendedImportPackages ...string) func(param TmplParam) error {
	return func(param TmplParam) error {
		for _, event := range param.Events {
			for _, p := range event.Params {
				param.ImportPackages = append(param.ImportPackages, r.AppendImportPackages(p.Type, recommendedImportPackages...)...)
			}
		}

		buf := new(bytes.Buffer)
		err := r.Tmpl.Execute(buf, param)
		if err != nil {
			return err
		}

		out, err := format.Source(buf.Bytes())
		if err != nil {
			return err
		}
		_, err = w.Write(out)
		return err
	}
}
