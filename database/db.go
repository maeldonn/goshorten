package database

import (
	"errors"
	"fmt"

	"github.com/maeldonn/url-shortener/types"
)

var links map[string]string

func init() {
	links = make(map[string]string)
}

func GetURL(slug string) (string, error) {
	if link, ok := links[slug]; ok {
		return link, nil
	}
	return "", errors.New("No URL for this slug")
}

func SaveSlug(shortened types.ShortenedURL) error {
	if _, ok := links[shortened.Slug]; ok {
		return fmt.Errorf("Slug [%s] is already used", shortened.Slug)
	}

	links[shortened.Slug] = shortened.URL
	return nil
}

func RemoveSlug(slug string) {
	delete(links, slug)
}
