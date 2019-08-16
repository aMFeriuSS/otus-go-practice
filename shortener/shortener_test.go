package shortener_test

import (
	"testing"

	"github.com/amferiuss/otus-go-practice/shortener"
)

func Test_ShortenerShorten(t *testing.T) {
	shortnr := shortener.NewShortener()

	want := "bing.com/36d798"
	if got := shortnr.Shorten("http://bing.com/search?q=dotnet"); got != want {
		t.Errorf("shortener.Shorten() = %q, want %q", got, want)
	}
	want = "bing.com/46dbfa"
	if got := shortnr.Shorten("http://bing.com/search?q=java"); got != want {
		t.Errorf("shortener.Shorten() = %q, want %q", got, want)
	}
	want = "bing.com/0a17b5"
	if got := shortnr.Shorten("http://bing.com/search?q=golang"); got != want {
		t.Errorf("shortener.Shorten() = %q, want %q", got, want)
	}
}

func Test_ShortenerResolve(t *testing.T) {
	shortnr := shortener.NewShortener()

	want := ""
	if got := shortnr.Resolve("bing.com/sdfds"); got != want {
		t.Errorf("shortener.Resolve() = %q, want %q", got, want)
	}

	shortnr.Shorten("http://bing.com/search?q=dotnet")
	shortnr.Shorten("http://bing.com/search?q=java")
	shortnr.Shorten("http://bing.com/search?q=golang")

	want = "http://bing.com/search?q=dotnet"
	if got := shortnr.Resolve("bing.com/36d798"); got != want {
		t.Errorf("shortener.Resolve() = %q, want %q", got, want)
	}

	want = "http://bing.com/search?q=java"
	if got := shortnr.Resolve("bing.com/46dbfa"); got != want {
		t.Errorf("shortener.Resolve() = %q, want %q", got, want)
	}

	want = "http://bing.com/search?q=golang"
	if got := shortnr.Resolve("bing.com/0a17b5"); got != want {
		t.Errorf("shortener.Resolve() = %q, want %q", got, want)
	}
}
