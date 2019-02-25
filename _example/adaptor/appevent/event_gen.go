// Code generated ; DO NOT EDIT

package appevent

import (
	"context"

	"github.com/hori-ryota/go-genappevent/_example/application"
	"github.com/hori-ryota/go-genappevent/_example/domain"
	"github.com/hori-ryota/go-genappevent/_example/external/appevent"
)

type EventPublisher struct {
	ShopOpenedEventPublisher
}

func NewEventPublisher(
	shopOpenedEventPublisher ShopOpenedEventPublisher,
) EventPublisher {
	return EventPublisher{
		ShopOpenedEventPublisher: shopOpenedEventPublisher,
	}
}

type ShopOpenedEventPublisher struct {
	publisher appevent.ShopOpenedEventPublisher
	bufSize   int
}

func NewShopOpenedEventPublisher(
	publisher appevent.ShopOpenedEventPublisher,
	bufSize int,
) ShopOpenedEventPublisher {
	return ShopOpenedEventPublisher{
		publisher: publisher,
		bufSize:   bufSize,
	}
}

func (p ShopOpenedEventPublisher) PublishShopOpenedEvent(ctx context.Context, event application.ShopOpenedEvent) <-chan domain.Error {
	c := make(chan domain.Error, p.bufSize)
	go func() {
		defer close(c)
		errSrc := p.publisher.PublishShopOpenedEvent(
			ctx,
			appevent.NewShopOpenedEvent(
				event.OccurredOn(),
				string(event.ShopID()),
				event.OperatorID(),
				event.OperatorType(),
			),
		)
		for err := range errSrc {
			c <- domain.ErrorUnknown(err)
		}
	}()
	return c
}

type ShopOpenedEventSubscriber struct {
	subscriber appevent.ShopOpenedEventSubscriber
	bufSize    int
}

func NewShopOpenedEventSubscriber(
	subscriber appevent.ShopOpenedEventSubscriber,
	bufSize int,
) ShopOpenedEventSubscriber {
	return ShopOpenedEventSubscriber{
		subscriber: subscriber,
		bufSize:    bufSize,
	}
}

func (p ShopOpenedEventSubscriber) SubscribeShopOpenedEvent(ctx context.Context) (<-chan application.ShopOpenedEvent, domain.Error) {
	src, err := p.subscriber.SubscribeShopOpenedEvent(ctx)
	if err != nil {
		return nil, domain.ErrorUnknown(err)
	}
	c := make(chan application.ShopOpenedEvent, p.bufSize)
	go func() {
		defer close(c)
		for event := range src {
			c <- application.NewShopOpenedEvent(
				event.OccurredOn(),
				domain.ShopID(event.ShopID()),
				event.OperatorID(),
				event.OperatorType(),
			)
		}
	}()
	return c, nil
}
