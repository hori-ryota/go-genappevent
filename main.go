/*
```go
    //events {EventName}[,paramName paramType]...
```

with `go generate` command

```go
    //go:generate go-genappevent {templateName} {dstFileName} [recommendImportPackagePath]...
```

`recommendImportPackagePath` is used for import resolution in generated code.
*/
package main

import (
	"errors"
	"log"
	"os"

	"github.com/hori-ryota/go-genappevent/genappevent"
)

var defaultTemplates = []genappevent.TemplateRenderer{
	genappevent.GoApplicationTmpl,
	genappevent.GoExternalClientTmpl,
	genappevent.GoAdapterPubsubberBaseTmpl,
}

func main() {
	if err := Main(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Main(args []string) error {
	templateName := args[1]
	dstFileName := args[2]
	recommendedImportPackages := args[3:]

	tmplMap := make(map[string]genappevent.TemplateRenderer, len(defaultTemplates))
	for i := range defaultTemplates {
		tmplMap[defaultTemplates[i].Tmpl.Name()] = defaultTemplates[i]
	}

	tmpl, ok := tmplMap[templateName]
	if !ok {
		return errors.New("unknown template name")
	}

	f, err := os.Create(dstFileName)
	if err != nil {
		return err
	}

	return genappevent.Run(".", tmpl.RendererFunc(f, recommendedImportPackages...))
}
