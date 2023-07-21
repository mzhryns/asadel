package repository

import "github.com/zhryn/asadel/entity"

type MetaDataRepository interface {
	GetOpenGraph(url string) (*entity.MetaData, error)
}
