package api

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	prisma "github.com/aejoy/prisma-sdk-go/internal/web"
	"github.com/aejoy/prisma-sdk-go/internal/web/dto"
	"github.com/aejoy/prisma-sdk-go/models"
	"github.com/aejoy/prisma-sdk-go/pkg/errors"
)

type API struct {
	client prisma.PrismaGraphQLClient
}

func New(baseURL string) API {
	return API{prisma.NewClient(http.DefaultClient, baseURL, nil)}
}

func (a API) GetByHash(context context.Context, sha256 string) (photos []models.Photo, err error) {
	res, err := a.client.GetByHash(context, &sha256)
	if err != nil {
		return photos, err
	}

	for _, photo := range res.Photos {
		if photo != nil {
			photo := *photo

			photos = append(photos, models.Photo{
				ID:        *photo.ID,
				HasAvatar: *photo.Has.Avatar,
				HasBanner: *photo.Has.Banner,
			})
		}
	}

	return photos, nil
}

func (a API) Upload(context context.Context, typ models.PhotoType, photo *graphql.Upload) (string, error) {
	var photoType dto.PhotoType

	switch typ {
	case models.PhotoTypePhoto:
		photoType = dto.PhotoTypePhoto
	case models.PhotoTypeAvatar:
		photoType = dto.PhotoTypeAvatar
	case models.PhotoTypeBanner:
		photoType = dto.PhotoTypeBanner
	}

	res, err := a.client.UploadPhoto(context, &photoType, photo)
	if err != nil {
		return "", err
	}

	photoID := res.GetUploadPhoto()
	if photoID == nil {
		return "", errors.ErrPhotoNil
	}

	return *photoID, err
}
