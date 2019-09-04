package practice

import (
	"testing"

	"github.com/amferiuss/otus-go-practice/ntp"
	"github.com/amferiuss/otus-go-practice/shortener"
)

// Week-1: ntp.
func Test_NtpTime(t *testing.T) {
	ntp.Time()
	ntp.NtpTime()
}

// Week-1: shortener.
func Test_Shortener(t *testing.T) {
	shortnr := shortener.NewShortener()

	link := "http://bing.com/search?q=golang"
	want := "bing.com/0a17b5"

	shortLink := shortnr.Shorten(link)
	if shortLink != want {
		t.Errorf("shortener.Shorten() = %q, want %q", shortLink, want)
	}
}