package url_repository_v1

import (
	"time"

	"github.com/zhryn/asadel/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Short       string             `bson:"short,omitempty"`
	Long        string             `bson:"long,omitempty"`
	Deeplink    string             `bson:"deeplink"`
	Android     string             `bson:"android"`
	Ios         string             `bson:"ios"`
	DateCreated time.Time          `bson:"date_created,omitempty"`
}

func (u *Url) ToUrlEntity() *entity.Url {
	return &entity.Url{
		Id:          u.Id.String()[10:34],
		Short:       u.Short,
		Long:        u.Long,
		Deeplink:    u.Deeplink,
		Android:     u.Android,
		Ios:         u.Ios,
		DateCreated: u.DateCreated,
	}
}

func (Url) FromUrlEntity(e *entity.Url) (*Url, error) {
	if e.Id == "" {
		return &Url{
			Short:       e.Short,
			Long:        e.Long,
			Deeplink:    e.Deeplink,
			Android:     e.Android,
			Ios:         e.Ios,
			DateCreated: e.DateCreated,
		}, nil
	}

	id, err := primitive.ObjectIDFromHex(e.Id)
	if err != nil {
		return nil, err
	}

	return &Url{
		Id:          id,
		Short:       e.Short,
		Long:        e.Long,
		Deeplink:    e.Deeplink,
		Android:     e.Android,
		Ios:         e.Ios,
		DateCreated: e.DateCreated,
	}, err
}
