package main

import (
	"context"
	"eh-backend-api/conf/drivers"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	userDriver, err := drivers.InitializeDriver(ctx)
	if err != nil {
		log.Printf("failed to create UserDriver: %s\n", err)
		os.Exit(2)
	}

	userDriver.Start(ctx)
}
