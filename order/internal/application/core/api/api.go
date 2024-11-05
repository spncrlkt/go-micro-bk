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
    if paymentErr != nil {
        st, _ := status.FromError(paymentErr)
        fieldErr := &errdetails.BadRequest_FieldViolation{
            Field: "payment"
            Description: st.Message(),
        }
        badReq := &errdetails.BadRequest{}
        badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
        orderStatus := status.New(codes.InvalidArgument, "order creation failed")
        statusWithDetails, _ := orderStatus.WithDetails(badReq)
        return domain.Order{}, statusWithDetails.Err()
    }
	return order, nil
}
