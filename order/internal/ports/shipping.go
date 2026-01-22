package ports

import (
	"context"


	
	"github.com/cassiaagomes/microservices/order/internal/application/core/domain"
)

type ShippingPort interface {
	Create(ctx context.Context, order *domain.Order) (int32, error) // retorna DeliveryDays
}