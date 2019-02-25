package application

import (
	"context"
	"time"

	"github.com/hori-ryota/go-genappevent/_example/domain"
	"github.com/hori-ryota/zaperr"
	"go.uber.org/zap"
)

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

//genconstructor
type OperatorDescriptor struct {
	id  string `required:"" getter:""`
	typ string `required:"" getter:""`
}

type ShopRepository interface {
	Create(ctx context.Context, shop domain.Shop) domain.Error
	Update(ctx context.Context, shop domain.Shop, version string) domain.Error
	FindByID(ctx context.Context, id domain.ShopID) (shop domain.Shop, version string, derr domain.Error)
	Delete(ctx context.Context, shop domain.Shop) domain.Error
}
