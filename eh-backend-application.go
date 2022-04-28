package main

import (
	"context"
	"eh-backend-api/conf/drivers"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	userDriver, err := drivers.InitializeUserDriver(ctx)
	if err != nil {
		fmt.Printf("failed to create UserDriver: %s\n", err)
		os.Exit(2)
	}
	userDriver.Start(ctx)
}
