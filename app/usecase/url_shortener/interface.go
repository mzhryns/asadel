package url_shortener

import "github.com/zhryn/asadel/entity"

type UseCase interface {
	Shorten(longUrl, deeplink, storeAndroid, storeIOS string) (string, error)
	Redirect(shortUrl string) (*entity.Url, error)
}
