package url_repository_v1

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	irepository "github.com/zhryn/asadel/app/repository"
	"github.com/zhryn/asadel/entity"
)

type repository struct {
	collection *mongo.Collection
}

func New(client *mongo.Client, dbName, dbCollection string) irepository.UrlRepository {
	return &repository{
		collection: client.Database(dbName).Collection(dbCollection),
	}
}

func (r *repository) Find(short string) (*entity.Url, error) {
	var urlData *Url

	err := r.collection.FindOne(
		context.Background(),
		bson.M{"short": short},
	).Decode(&urlData)

	out := urlData.ToUrlEntity()
	return out, err
}

func (r *repository) Save(url *entity.Url) error {
	urlData, err := new(Url).FromUrlEntity(url)
	if err != nil {
		return err
	}

	urlData.DateCreated = time.Now()

	_, err = r.collection.InsertOne(
		context.Background(),
		urlData,
	)

	return err
}
