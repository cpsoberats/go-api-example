package apiconfig

import (
	"api/src/libs"
)


type IApiConfigService interface {
}

type Service struct {
	Url         string
	HttpClient  libs.IHttpClient
}

