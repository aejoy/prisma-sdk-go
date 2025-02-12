package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aejoy/prisma-sdk-go/api"
)

func main() {
	api := api.New("http://localhost:4000")

	context, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	photo, err := api.CheckBySHA256(context, "5b3b01789bfa80a2872ae6fd1e1acd47f84d54de76e2ee96bff987360d85440ae1220c9256549024d84b75fb8af11b605077ab4501d348dc9dd462de78a2c13e")
	if err != nil {
		panic(err)
	}

	fmt.Println("Exist photo:", photo.ID)
}
