package helpers

import (
	"fmt"
	"net/url"

	"github.com/zhryn/asadel/entity"
)

func FormatUrl(baseUrl string, e *entity.Url) string {
	u := url.QueryEscape(e.Long)
	d := url.QueryEscape(e.Deeplink)
	a := url.QueryEscape(e.Android)
	i := url.QueryEscape(e.Ios)

	return fmt.Sprintf("%s?u=%s&d=%s&a=%s&i=%s", baseUrl, u, d, a, i)
}
