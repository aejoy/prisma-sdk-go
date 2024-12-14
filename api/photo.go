package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/aejoy/prisma-sdk-go/consts"

	"github.com/aejoy/go-pkg/utils"
	"github.com/aejoy/prisma-sdk-go/responses"
	"github.com/aejoy/prisma-sdk-go/types"
	"github.com/pkg/errors"
)

func (api *API) GetPhotos(params types.GetPhotos) (responses.Photos, error) {
	if params.Count == nil {
		params.Count = utils.PointerValue[int](consts.DefaultCount)
	}

	if params.Offset == nil {
		params.Offset = utils.PointerValue[int](consts.DefaultOffset)
	}

	queryParams := &url.Values{
		"count":  {fmt.Sprintf("%d", *params.Count)},
		"offset": {fmt.Sprintf("%d", *params.Offset)},
	}

	photos := responses.Photos{}

	if len(params.PhotoIDs) > 0 {
		queryParams.Add("ids", strings.Join(params.PhotoIDs, ","))
	}

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, fmt.Sprintf("%s/photos?%s", api.InstanceAddress, queryParams.Encode()), nil)
	if err != nil {
		return photos, errors.Wrap(err, "NewRequest")
	}

	res, err := api.client.Do(req)
	if err != nil {
		return photos, errors.Wrap(err, "Client Do")
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&photos); err != nil {
		return photos, errors.Wrap(err, "Unmarshal")
	}

	return photos, nil
}
