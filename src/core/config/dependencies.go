
package config

import (
	"api/src/core/apiconfig"
	"api/src/core/config/dataclient"
	"api/src/libs"
	"api/src/providers"
	"github.com/sarulabs/di"
	"os"
)

func Container() di.Container {

	builder, _ := di.NewBuilder()

	_ = builder.Add(di.Def{
		Name: "http-client",
		Build: func(ctn di.Container) (i interface{}, e error) {
			return libs.HttpClient{}, nil
		},
	})

	_ = builder.Add(di.Def{
		Name: "db-postgres",
		Build: func(ctn di.Container) (i interface{}, e error) {
			return dataclient.Postgres{
				DbConnection: dataclient.DbConnection{
					Username: os.Getenv("POSTGRES_USERNAME"),
					Password: os.Getenv("POSTGRES_PASSWORD"),
					Host:     os.Getenv("POSTGRES_HOST"),
					Port:     os.Getenv("POSTGRES_PORT"),
					Database: os.Getenv("POSTGRES_DATABASE"),
				},
			}, nil
		},
	})

	_ = builder.Add(di.Def{
		Name: "data-client",
		Build: func(ctn di.Container) (i interface{}, e error) {
			return dataclient.DataClient{
				Postgres: ctn.Get("db-postgres").(dataclient.ISqlClient),
			}, nil
		},
	})

	_ = builder.Add(di.Def{
		Name: "api-config-service",
		Build: func(ctn di.Container) (i interface{}, e error) {
			return apiconfig.Service{
				Url:         os.Getenv("API_CONFIG_URL"),
				HttpClient:  ctn.Get("http-client").(libs.HttpClient),
			}, nil
		},
	})
	_ = builder.Add(di.Def{
		Name: "app-provider",
		Build: func(ctn di.Container) (i interface{}, e error) {
			provider := providers.New(ctn.Get("api-config-service").(apiconfig.Service))
			return provider, nil
		},
	})

	return builder.Build()
}
