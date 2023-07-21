package meta_data_reader

import (
	"github.com/zhryn/asadel/app/repository"
	"github.com/zhryn/asadel/entity"
)

type usecase struct {
	meta_data_repo repository.MetaDataRepository
}

func NewUseCase(meta_data_repo repository.MetaDataRepository) UseCase {
	return &usecase{meta_data_repo: meta_data_repo}
}

func (uc *usecase) GetOpenGraph(url string) (*entity.MetaData, error) {
	meta, err := uc.meta_data_repo.GetOpenGraph(url)
	return meta, err
}
