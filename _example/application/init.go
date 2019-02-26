package application

//go:generate go-genaccessor
//go:generate go-genconstructor

//go:generate go-genappevent GoApplicationTmpl event_gen.go github.com/hori-ryota/go-genappevent/_example/domain
//go:generate go-genappevent GoExternalClientTmpl ../external/appevent/event_gen.go github.com/hori-ryota/go-genappevent/_example/application github.com/hori-ryota/go-genappevent/_example/domain
//go:generate go-genappevent GoAdapterPubsubberBaseTmpl ../adapter/appevent/event_gen.go github.com/hori-ryota/go-genappevent/_example/application github.com/hori-ryota/go-genappevent/_example/domain github.com/hori-ryota/go-genappevent/_example/external/appevent
