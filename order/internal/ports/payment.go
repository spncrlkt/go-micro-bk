package ports

import "github.com/spncrlkt/go-micro-bk/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
