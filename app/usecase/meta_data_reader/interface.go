package meta_data_reader

import "github.com/zhryn/asadel/entity"

type UseCase interface {
	GetOpenGraph(url string) (*entity.MetaData, error)
}
