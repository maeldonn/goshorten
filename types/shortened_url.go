package types

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type ShortenedURL struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

func ExtractShortenedURL(body io.ReadCloser) ShortenedURL {
	var shortened ShortenedURL
	data, _ := ioutil.ReadAll(body)
	defer body.Close()

	json.Unmarshal(data, &shortened)
	return shortened
}
