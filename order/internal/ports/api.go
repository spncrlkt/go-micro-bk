package ports

import "github.com/spncrlkt/go-micro-bk/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
