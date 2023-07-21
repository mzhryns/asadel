package meta_data_repository_v1

import (
	irepository "github.com/zhryn/asadel/app/repository"
	"github.com/zhryn/asadel/entity"
	"github.com/zhryn/asadel/pkg/helpers"
)

type repository struct{}

func New() irepository.MetaDataRepository {
	return &repository{}
}

func (r *repository) GetOpenGraph(url string) (*entity.MetaData, error) {
	resp, err := helpers.FetchPage(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	metaData := new(MetaData).FromResponseBody(resp.Body)

	out := metaData.ToMetaDataEntity()
	return out, nil
}
