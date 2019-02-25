// Code generated ; DO NOT EDIT

package application

import (
	"context"
	"time"

	"github.com/hori-ryota/go-genappevent/_example/domain"
)

type ShopOpenedEvent struct {
	occurredOn   time.Time
	shopID       domain.ShopID
	operatorID   string
	operatorType string
}

func NewShopOpenedEvent(
	occurredOn time.Time,
	shopID domain.ShopID,
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

func (e ShopOpenedEvent) ShopID() domain.ShopID {
	return e.shopID
}
func (e ShopOpenedEvent) OperatorID() string {
	return e.operatorID
}
func (e ShopOpenedEvent) OperatorType() string {
	return e.operatorType
}

type ShopOpenedEventPublisher interface {
	PublishShopOpenedEvent(context.Context, ShopOpenedEvent) <-chan domain.Error
}

type ShopOpenedEventSubscriber interface {
	SubscribeShopOpenedEvent(context.Context) (<-chan ShopOpenedEvent, domain.Error)
}
