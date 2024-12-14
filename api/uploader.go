package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/aejoy/prisma-sdk-go/responses"
	"github.com/pkg/errors"
)

func (api *API) Upload(typ string, file io.Reader) (responses.Upload, error) {
	upload := responses.Upload{}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	field, err := writer.CreateFormField("type")
	if err != nil {
		return upload, errors.Wrap(err, "FormField")
	}

	if _, err := field.Write([]byte(typ)); err != nil {
		return upload, errors.Wrap(err, "FormField")
	}

	part, err := writer.CreateFormFile("file", "photo.png")
	if err != nil {
		return upload, errors.Wrap(err, "FormFile Banner")
	}

	if _, err := io.Copy(part, file); err != nil {
		return upload, err
	}

	if err := writer.Close(); err != nil {
		return upload, errors.Wrap(err, "Close Write")
	}

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, api.InstanceAddress+"/photos/upload", body)
	if err != nil {
		return upload, errors.Wrap(err, "NewRequest")
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	defer req.Body.Close()

	res, err := api.client.Do(req)
	if err != nil {
		return upload, errors.Wrap(err, "Client Do")
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&upload); err != nil {
		return upload, errors.Wrap(err, "Unmarshal")
	}

	return upload, nil
}
