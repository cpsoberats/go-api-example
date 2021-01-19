package providers

import (
	"api/src/core/apiconfig"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"time"
)

type AppProvider struct {
	ApiConfigService apiconfig.Service
}

func New(ApiConfigService apiconfig.Service) *AppProvider {
	app := AppProvider{ApiConfigService}
	app.ValidatorProvider()
	return &app
}

func (a AppProvider) ValidatorProvider() {
	// simple example
	govalidator.AddCustomRule("datetime", func(field string, rule string, message string, value interface{}) error {
		_, err := time.Parse("2006-01-02 15:04:05.999", value.(string))
		if err != nil {
			return fmt.Errorf("The %s field must be in yyyy-mm-dd hh:ii:ss format", field)
		}

		return nil
	})

}
