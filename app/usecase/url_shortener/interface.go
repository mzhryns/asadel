package url_shortener

type UseCase interface {
	Shorten(longUrl string) (string, error)
	Redirect(shortUrl string) (string, error)
}
