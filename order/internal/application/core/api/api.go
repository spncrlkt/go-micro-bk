package api

import (
	"github.com/spncrlkt/go-micro-bk/order/internal/application/core/domain"
	"github.com/spncrlkt/go-micro-bk/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
    paymentErr := a.payment.Charge(&order)
    if paymentError != nil {
        return domain.Order{}, paymentErr
    }
	return order, nil
}
