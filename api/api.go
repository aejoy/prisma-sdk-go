package api

import (
	"context"
	"net/http"

	prisma "github.com/aejoy/prisma-sdk-go/internal/web"
	"github.com/aejoy/prisma-sdk-go/models"
	"github.com/aejoy/prisma-sdk-go/pkg/consts"
)

type API struct {
	client prisma.PrismaGraphQLClient
}

func New(baseURL string) API {
	return API{prisma.NewClient(http.DefaultClient, baseURL, nil)}
}

func (a API) CheckBySHA256(hash string) (photo models.Photo, err error) {
	timeout, cancel := context.WithTimeout(context.Background(), consts.DefaultTimeout)
	defer cancel()

	res, err := a.client.CheckBySha256(timeout, &hash)
	if err != nil {
		return photo, err
	}

	for _, existPhoto := range res.Photos {
		photo.ID = *existPhoto.ID
	}

	return photo, nil
}
