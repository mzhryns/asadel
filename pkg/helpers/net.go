package helpers

import "net/http"

func FetchPage(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}
