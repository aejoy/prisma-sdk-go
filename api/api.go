package api

import "net/http"

type API struct {
	client          *http.Client
	InstanceAddress string
}

func New(instanceAddress string) *API {
	return &API{http.DefaultClient, instanceAddress}
}
