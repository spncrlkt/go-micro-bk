package main

import (
	"github.com/spncrlkt/go-micro-bk/order/config"
	"github.com/spncrlkt/go-micro-bk/order/internal/adapters/db"
	"github.com/spncrlkt/go-micro-bk/order/internal/adapters/grpc"
	"github.com/spncrlkt/go-micro-bk/order/internal/adapters/payment"
	"github.com/spncrlkt/go-micro-bk/order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to db. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
