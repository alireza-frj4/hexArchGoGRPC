package main

import (
	"log"
	"os"

	"github.com/alireza-frj4/app1/internal/adapters/app/api"
	"github.com/alireza-frj4/app1/internal/adapters/core/arithmetic"
	"github.com/alireza-frj4/app1/internal/adapters/framwork/right/db"
	"github.com/alireza-frj4/app1/internal/ports"

	gRPC "github.com/alireza-frj4/app1/internal/adapters/framwork/left/grpc"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPorts
	var gRPCAdapter ports.GEPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnecton()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}
