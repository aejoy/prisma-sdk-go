package main

import (
	"context"
	"fmt"

	"github.com/aejoy/prisma-sdk-go/api"
	"github.com/aejoy/prisma-sdk-go/pkg/consts"
)

func main() {
	api := api.New("http://localhost:4000")

	context, cancel := context.WithTimeout(context.TODO(), consts.DefaultTimeout)
	defer cancel()

	photos, err := api.GetByHash(context, "ff526ae11096704558d55d0a431aa3f1775e4d36e2af0765f081a873c5bc4e5c6d4d85b23562ab0ddb603062f29ee29c8006dd0da5a7a7a1e1146ac3afe92e4f")
	if err != nil {
		panic(err)
	}

	fmt.Println(photos)
}
