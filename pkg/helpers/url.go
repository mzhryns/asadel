package helpers

import (
	"fmt"
	"net/url"
)

func FormatUrl(baseUrl, longUrl string) string {
	encodedUrl := url.QueryEscape(longUrl)
	return fmt.Sprintf("%s?url=%s", baseUrl, encodedUrl)
}
