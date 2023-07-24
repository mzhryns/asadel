package helpers

import (
	"fmt"
	"net/url"

	"github.com/zhryn/asadel/entity"
)

func FormatUrl(baseUrl string, e *entity.Url) string {
	queryParams := url.Values{
		"u": []string{e.Long},
		"d": []string{e.Deeplink},
		"a": []string{e.Android},
		"i": []string{e.Ios},
	}

	return fmt.Sprintf("%s?%s", baseUrl, queryParams.Encode())
}
