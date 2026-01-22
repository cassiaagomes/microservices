package main

import (
	"log"

	"github.com/cassiaagomes/microservices/order/config"
	"github.com/cassiaagomes/microservices/order/internal/adapters/db"
	"github.com/cassiaagomes/microservices/order/internal/adapters/grpc"
	payment_adapter "github.com/cassiaagomes/microservices/order/internal/adapters/payment"
	shipping_adapter "github.com/cassiaagomes/microservices/order/internal/adapters/shipping"

	"github.com/cassiaagomes/microservices/order/internal/application/core/api"
)


func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURl()) 
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Errpr: %v", err)
	}
	shippingAdapter, err := shipping_adapter.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize shipping stub. Errpr: %v", err)

	}
	application := api.NewApplication(dbAdapter,paymentAdapter,shippingAdapter, dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}

// This application have 2 ttk