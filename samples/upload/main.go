package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aejoy/prisma-sdk-go/models"
	"github.com/aejoy/prisma-sdk-go/pkg/consts"

	"github.com/aejoy/prisma-sdk-go/api"
)

func main() {
	api := api.New("http://localhost:4000")

	context, cancel := context.WithTimeout(context.TODO(), consts.DefaultTimeout)
	defer cancel()

	file, err := os.ReadFile("./99c0d458e948ac62b09213f165d8dafb.jpg")
	if err != nil {
		panic(err)
	}

	photoID, err := api.Upload(context, models.PhotoTypePhoto, &graphql.Upload{
		File: bytes.NewReader(file),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded photo ID:", photoID)
}
