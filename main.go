package gorseclient

import (
	"github.com/LouisAldorio/go-gorse-client/gorseclient"
)

func NewAPI(baseUrl string, token string, option ...gorseclient.Option) *gorseclient.API {
	return gorseclient.NewAPI(baseUrl, token, option...)
}
