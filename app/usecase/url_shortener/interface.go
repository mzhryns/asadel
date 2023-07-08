package url_shortener

type UseCase interface {
	Shorten(longUrl string) (string, error)
	Resolve(shortUrl string) (string, error)
}
