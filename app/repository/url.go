package repository

import (
	"github.com/zhryn/asadel/entity"
)

type UrlRepository interface {
	Find(short string) (*entity.Url, error)
	Save(url *entity.Url) error
}
