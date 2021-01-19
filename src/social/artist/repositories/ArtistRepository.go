package repositories

import (
	"api/src/core/config/dataclient"
	"api/src/social/artist/entities"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"log"
)

type IArtistRepository interface {
	GetArtistList(from, size uint64) []entities.Artist
	GetDataClient() dataclient.IDataClient
}

type ArtistRepository struct {
	DataClient dataclient.IDataClient
}
func (r ArtistRepository) GetDataClient() dataclient.IDataClient {
	return r.DataClient
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func (r ArtistRepository) GetArtistList(from, size uint64) []entities.Artist {

   dbConnection := r.GetDataClient().PostgresClient().Client()
	defer func() {
		dbConnection.Close()
		if r := recover(); r != nil {
			log.Print("error in GetArtistList")
			err, _ := r.(stackTracer)
			log.Print(err.StackTrace())
		}
	}()

	var artists []entities.Artist

   	rows, err := squirrel.Select(
		"id, name").
		From(`"Artist"`).
   		OrderBy("id asc").
		Limit(size).
		Offset(from).
		RunWith (dbConnection).Query()

	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		artist := entities.Artist{}

		_ = rows.Scan(&artist.Id, &artist.Name)

		artists = append(artists, artist)
	}

	return artists
}
