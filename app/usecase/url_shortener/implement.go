package url_shortener

import (
	"github.com/google/uuid"
	"github.com/zhryn/asadel/app/repository"
	"github.com/zhryn/asadel/entity"
)

type usecase struct {
	url_repo repository.UrlRepository
}

func NewUseCase(url_repo repository.UrlRepository) UseCase {
	return &usecase{url_repo: url_repo}
}

func (uc *usecase) Shorten(longUrl string) (string, error) {
	shortUrl := uuid.New().String()[:8]
	url := &entity.Url{Short: shortUrl, Long: longUrl}
	err := uc.url_repo.Save(url)

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (uc *usecase) Redirect(shortUrl string) (string, error) {
	url, err := uc.url_repo.Find(shortUrl)
	if err != nil {
		return "", err
	}

	return url.Long, nil
}
