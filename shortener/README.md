## Link shortener (part 1)

```go
type Shortener interface {
    Shorten(url string) string
    Resolve(url string) string
}
```

`Shorten` returns **_short link_** (example - otus.ru/some-long-link -> otus.ru/jhg34) and saves mapping short link to **_source link_**. 

`Resolve` returns **_source link_** or empty string if **_source link_** was not found.

Tips: 
 - Use memory without external DB (we will use it later)
