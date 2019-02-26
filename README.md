# go-genappevent

Application Event generator for Go.

## Feature

- Generate application event struct
- Generate publisher in application layer
- Generate subscriber in application layer
- Generate publisher in external layer
- Generate subscriber in external layer
- Generate multi publisher in external layer
- Generate publisher in adapter layer (Using external layer publisher and implemented application layer publisher)
- Generate subscriber in adapter layer (Using external layer subscriber and implemented application layer subscriber)
- Custom template

## Usage

```go
    //events {EventName}[,paramName paramType]...
```

with `go generate` command

```go
    //go:generate go-genappevent {templateName} {dstFileName} [recommendImportPackagePath]...
```

`recommendImportPackagePath` is used for import resolution in generated code.

generated Event and Publisher usage sample is following

```go
type Opener struct {
	repository     ShopRepository
	eventPublisher ShopOpenedEventPublisher
	logger         *zap.Logger
}

func (y Opener) Open(
	ctx context.Context,
	operator OperatorDescriptor,
	shopID domain.ShopID,
) domain.Error {

	// TODO: authorization

	shop, version, err := y.repository.FindByID(ctx, shopID)
	if err != nil {
		return err
	}

	if err := shop.Open(); err != nil {
		return err
	}

	if err := y.repository.Update(ctx, shop, version); err != nil {
		return err
	}

	go func() {
		//event ShopOpened,shopID domain.ShopID,operatorID string,operatorType string
		event := NewShopOpenedEvent(
			time.Now(),
			shopID,
			operator.ID(),
			operator.Typ(),
		)
		for err := range y.eventPublisher.PublishShopOpenedEvent(context.Background(), event) {
			y.logger.Error(
				"failed to publish event",
				zap.String("eventName", event.EventName()),
				zap.Any("event", event),
				zaperr.ToField(err),
			)
		}
	}()

	return nil
}
```

### Example

def
- [./_example/appliation/opener.go](./_example/application/opener.go)
- [./_example/appliation/init.go](./_example/application/init.go)

generated
- [./_example/appliation/event_gen.go](./_example/application/event_gen.go)
- [./_example/external/appevent/event_gen.go](./_example/external/appevent/event_gen.go)
- [./_example/adapter/appevent/event_gen.go](./_example/adapter/appevent/event_gen.go)

## Installation

```sh
$ go get github.com/hori-ryota/go-genappevent
```
