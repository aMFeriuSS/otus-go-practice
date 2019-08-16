package shortener

import (
	"crypto/md5"
	"fmt"
    "log"
	"net/url"
)

type Shortener interface {
	// Shorten returns short link (example - otus.ru/some-long-link -> otus.ru/jhg34).
	// Also saves mapping short link to source link.
	Shorten(url string) string
	// Resolve returns source link or empty string if source link was not found.
    Resolve(url string) string
}

type shortener struct {
	store map[string]string
}

func NewShortener() Shortener {
	return &shortener{
		store: make(map[string]string),
	}
}

func (s *shortener) Shorten(link string) string {
	if short, ok := s.store[link]; ok {
		return short
	}

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	
	host := u.Host
	data := []byte(link)
	code := fmt.Sprintf("%x", md5.Sum(data))
	shortUrl := host + "/" + code[:6]
	s.store[shortUrl] = link

	return shortUrl
}

func (s *shortener) Resolve(shortUrl string) string {
	return s.store[shortUrl]
}