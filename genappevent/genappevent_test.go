package genappevent_test

import (
	"log"
	"os"

	"github.com/hori-ryota/go-genappevent/genappevent"
)

func ExampleRun_genForGoApplicationTmpl() {
	targetDir := "../_example/application"
	if err := genappevent.Run(
		targetDir,
		genappevent.GoApplicationTmpl.RendererFunc(
			os.Stdout,
			"github.com/hori-ryota/go-genappevent/_example/domain",
		),
	); err != nil {
		log.Fatal(err)
	}
	// Output:
	// // Code generated ; DO NOT EDIT
	//
	// package application
	//
	// import (
	// 	"context"
	// 	"time"
	//
	// 	"github.com/hori-ryota/go-genappevent/_example/domain"
	// )
	//
	// type ShopOpenedEvent struct {
	// 	occurredOn   time.Time
	// 	shopID       domain.ShopID
	// 	operatorID   string
	// 	operatorType string
	// }
	//
	// func NewShopOpenedEvent(
	// 	occurredOn time.Time,
	// 	shopID domain.ShopID,
	// 	operatorID string,
	// 	operatorType string,
	// ) ShopOpenedEvent {
	// 	return ShopOpenedEvent{
	// 		occurredOn:   occurredOn,
	// 		shopID:       shopID,
	// 		operatorID:   operatorID,
	// 		operatorType: operatorType,
	// 	}
	// }
	//
	// func (e ShopOpenedEvent) EventName() string {
	// 	return "ShopOpened"
	// }
	//
	// func (e ShopOpenedEvent) OccurredOn() time.Time {
	// 	return e.occurredOn
	// }
	//
	// func (e ShopOpenedEvent) ShopID() domain.ShopID {
	// 	return e.shopID
	// }
	// func (e ShopOpenedEvent) OperatorID() string {
	// 	return e.operatorID
	// }
	// func (e ShopOpenedEvent) OperatorType() string {
	// 	return e.operatorType
	// }
	//
	// type ShopOpenedEventPublisher interface {
	// 	PublishShopOpenedEvent(context.Context, ShopOpenedEvent) <-chan domain.Error
	// }
	//
	// type ShopOpenedEventSubscriber interface {
	// 	SubscribeShopOpenedEvent(context.Context) (<-chan ShopOpenedEvent, domain.Error)
	// }
}

func ExampleRun_genForGoExternalClientTmpl() {
	targetDir := "../_example/application"
	if err := genappevent.Run(
		targetDir,
		genappevent.GoExternalClientTmpl.RendererFunc(
			os.Stdout,
			"github.com/hori-ryota/go-genappevent/_example/application",
			"github.com/hori-ryota/go-genappevent/_example/domain",
		),
	); err != nil {
		log.Fatal(err)
	}
	// Output:
	// // Code generated ; DO NOT EDIT
	//
	// package appevent
	//
	// import (
	// 	"context"
	// 	"sync"
	// 	"time"
	// )
	//
	// type ShopOpenedEvent struct {
	// 	occurredOn   time.Time
	// 	shopID       string
	// 	operatorID   string
	// 	operatorType string
	// }
	//
	// func NewShopOpenedEvent(
	// 	occurredOn time.Time,
	// 	shopID string,
	// 	operatorID string,
	// 	operatorType string,
	// ) ShopOpenedEvent {
	// 	return ShopOpenedEvent{
	// 		occurredOn:   occurredOn,
	// 		shopID:       shopID,
	// 		operatorID:   operatorID,
	// 		operatorType: operatorType,
	// 	}
	// }
	//
	// func (e ShopOpenedEvent) EventName() string {
	// 	return "ShopOpened"
	// }
	//
	// func (e ShopOpenedEvent) OccurredOn() time.Time {
	// 	return e.occurredOn
	// }
	//
	// func (e ShopOpenedEvent) ShopID() string {
	// 	return e.shopID
	// }
	// func (e ShopOpenedEvent) OperatorID() string {
	// 	return e.operatorID
	// }
	// func (e ShopOpenedEvent) OperatorType() string {
	// 	return e.operatorType
	// }
	//
	// type ShopOpenedEventPublisher interface {
	// 	PublishShopOpenedEvent(context.Context, ShopOpenedEvent) <-chan error
	// }
	//
	// type ShopOpenedEventSubscriber interface {
	// 	SubscribeShopOpenedEvent(context.Context) (<-chan ShopOpenedEvent, error)
	// }
	//
	// type ShopOpenedEventPublishers struct {
	// 	publishers []ShopOpenedEventPublisher
	// 	bufSize    int
	// }
	//
	// func NewShopOpenedEventPublishers(
	// 	publishers []ShopOpenedEventPublisher,
	// 	bufSize int,
	// ) ShopOpenedEventPublishers {
	// 	return ShopOpenedEventPublishers{
	// 		publishers: publishers,
	// 		bufSize:    bufSize,
	// 	}
	// }
	//
	// func (ps ShopOpenedEventPublishers) PublishShopOpenedEvent(ctx context.Context, event ShopOpenedEvent) <-chan error {
	// 	wg := new(sync.WaitGroup)
	// 	c := make(chan error, ps.bufSize)
	// 	go func() {
	// 		defer close(c)
	// 		wg.Wait()
	// 	}()
	// 	for i := range ps.publishers {
	// 		wg.Add(1)
	// 		go func() {
	// 			wg.Done()
	// 			for err := range ps.publishers[i].PublishShopOpenedEvent(ctx, event) {
	// 				c <- err
	// 			}
	// 		}()
	// 	}
	// 	return c
	// }
}

func ExampleRun_genForGoAdapterPubsubberBaseTmpl() {
	targetDir := "../_example/application"
	if err := genappevent.Run(
		targetDir,
		genappevent.GoAdapterPubsubberBaseTmpl.RendererFunc(
			os.Stdout,
			"github.com/hori-ryota/go-genappevent/_example/application",
			"github.com/hori-ryota/go-genappevent/_example/domain",
			"github.com/hori-ryota/go-genappevent/_example/external/appevent",
		),
	); err != nil {
		log.Fatal(err)
	}
	// Output:
	// // Code generated ; DO NOT EDIT
	//
	// package appevent
	//
	// import (
	// 	"context"
	//
	// 	"github.com/hori-ryota/go-genappevent/_example/application"
	// 	"github.com/hori-ryota/go-genappevent/_example/domain"
	// 	"github.com/hori-ryota/go-genappevent/_example/external/appevent"
	// )
	//
	// type EventPublisher struct {
	// 	ShopOpenedEventPublisher
	// }
	//
	// func NewEventPublisher(
	// 	shopOpenedEventPublisher ShopOpenedEventPublisher,
	// ) EventPublisher {
	// 	return EventPublisher{
	// 		ShopOpenedEventPublisher: shopOpenedEventPublisher,
	// 	}
	// }
	//
	// type ShopOpenedEventPublisher struct {
	// 	publisher appevent.ShopOpenedEventPublisher
	// 	bufSize   int
	// }
	//
	// func NewShopOpenedEventPublisher(
	// 	publisher appevent.ShopOpenedEventPublisher,
	// 	bufSize int,
	// ) ShopOpenedEventPublisher {
	// 	return ShopOpenedEventPublisher{
	// 		publisher: publisher,
	// 		bufSize:   bufSize,
	// 	}
	// }
	//
	// func (p ShopOpenedEventPublisher) PublishShopOpenedEvent(ctx context.Context, event application.ShopOpenedEvent) <-chan domain.Error {
	// 	c := make(chan domain.Error, p.bufSize)
	// 	go func() {
	// 		defer close(c)
	// 		errSrc := p.publisher.PublishShopOpenedEvent(
	// 			ctx,
	// 			appevent.NewShopOpenedEvent(
	// 				event.OccurredOn(),
	// 				string(event.ShopID()),
	// 				event.OperatorID(),
	// 				event.OperatorType(),
	// 			),
	// 		)
	// 		for err := range errSrc {
	// 			c <- domain.ErrorUnknown(err)
	// 		}
	// 	}()
	// 	return c
	// }
	//
	// type ShopOpenedEventSubscriber struct {
	// 	subscriber appevent.ShopOpenedEventSubscriber
	// 	bufSize    int
	// }
	//
	// func NewShopOpenedEventSubscriber(
	// 	subscriber appevent.ShopOpenedEventSubscriber,
	// 	bufSize int,
	// ) ShopOpenedEventSubscriber {
	// 	return ShopOpenedEventSubscriber{
	// 		subscriber: subscriber,
	// 		bufSize:    bufSize,
	// 	}
	// }
	//
	// func (p ShopOpenedEventSubscriber) SubscribeShopOpenedEvent(ctx context.Context) (<-chan application.ShopOpenedEvent, domain.Error) {
	// 	src, err := p.subscriber.SubscribeShopOpenedEvent(ctx)
	// 	if err != nil {
	// 		return nil, domain.ErrorUnknown(err)
	// 	}
	// 	c := make(chan application.ShopOpenedEvent, p.bufSize)
	// 	go func() {
	// 		defer close(c)
	// 		for event := range src {
	// 			c <- application.NewShopOpenedEvent(
	// 				event.OccurredOn(),
	// 				domain.ShopID(event.ShopID()),
	// 				event.OperatorID(),
	// 				event.OperatorType(),
	// 			)
	// 		}
	// 	}()
	// 	return c, nil
	// }
}
