package config

import (
	"api/src/core/config"
	"api/src/core/config/dataclient"
	"api/src/social/artist/http/controllers"
	"api/src/social/artist/repositories"
	"api/src/social/artist/services"
	"github.com/sarulabs/di"
)

func Container() di.Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add(di.Def{
		Name: "artist_repository",
		Build: func(ctn di.Container) (interface{}, error) {
			return repositories.ArtistRepository{
				DataClient: config.Container().Get("data-client").(dataclient.IDataClient),
			}, nil
		},
	})

	_ = builder.Add(di.Def{
		Name: "artist_service",
		Build: func(ctn di.Container) (interface{}, error) {
			return services.ArtistService{
				ArtistRepository: ctn.Get("artist_repository").(repositories.ArtistRepository),
			}, nil
		},
	})

	_ = builder.Add(di.Def{
		Name: "artist_controller",
		Build: func(ctn di.Container) (interface{}, error) {
			return controllers.ArtistController{
				ArtistService: ctn.Get("artist_service").(services.IArtistService),
			}, nil
		},
	})

	return builder.Build()
}