// Code generated ; DO NOT EDIT

package appevent

import (
	"context"
	"sync"
	"time"
)

type ShopOpenedEvent struct {
	occurredOn   time.Time
	shopID       string
	operatorID   string
	operatorType string
}

func NewShopOpenedEvent(
	occurredOn time.Time,
	shopID string,
	operatorID string,
	operatorType string,
) ShopOpenedEvent {
	return ShopOpenedEvent{
		occurredOn:   occurredOn,
		shopID:       shopID,
		operatorID:   operatorID,
		operatorType: operatorType,
	}
}

func (e ShopOpenedEvent) EventName() string {
	return "ShopOpened"
}

func (e ShopOpenedEvent) OccurredOn() time.Time {
	return e.occurredOn
}

func (e ShopOpenedEvent) ShopID() string {
	return e.shopID
}
func (e ShopOpenedEvent) OperatorID() string {
	return e.operatorID
}
func (e ShopOpenedEvent) OperatorType() string {
	return e.operatorType
}

type ShopOpenedEventPublisher interface {
	PublishShopOpenedEvent(context.Context, ShopOpenedEvent) <-chan error
}

type ShopOpenedEventSubscriber interface {
	SubscribeShopOpenedEvent(context.Context) (<-chan ShopOpenedEvent, error)
}

type ShopOpenedEventPublishers struct {
	publishers []ShopOpenedEventPublisher
	bufSize    int
}

func NewShopOpenedEventPublishers(
	publishers []ShopOpenedEventPublisher,
	bufSize int,
) ShopOpenedEventPublishers {
	return ShopOpenedEventPublishers{
		publishers: publishers,
		bufSize:    bufSize,
	}
}

func (ps ShopOpenedEventPublishers) PublishShopOpenedEvent(ctx context.Context, event ShopOpenedEvent) <-chan error {
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
			for err := range ps.publishers[i].PublishShopOpenedEvent(ctx, event) {
				c <- err
			}
		}()
	}
	return c
}
