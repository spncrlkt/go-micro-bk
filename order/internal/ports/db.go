package ports

import "github.com/spncrlkt/go-micro-bk/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
