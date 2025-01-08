package core

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/idna"
)

// REF: https://github.com/sekimura/go-normalize-url/blob/master/normalizeurl.go
var (
	DefaultPorts = map[string]int{
		"http":  80,
		"https": 443,
		"ftp":   21,
	}
)

// REF: https://github.com/sekimura/go-normalize-url/blob/master/normalizeurl.go
// Normalize url strings
// http://en.wikipedia.org/wiki/URL_normalization
func NormalizeURL(s string) (string, error) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "//") {
		s = "http:" + s
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	if u.Scheme == "" {
		// Ugh...
		u, err = url.Parse("http://" + s)
		if err != nil {
			return s, err
		}
	}

	p, ok := DefaultPorts[u.Scheme]
	if ok {
		u.Host = strings.TrimSuffix(u.Host, fmt.Sprintf(":%d", p))
	}

	got, err := idna.ToUnicode(u.Host)
	if err != nil {
		return got, err
	} else {
		u.Host = got
	}

	v := u.Query()
	u.RawQuery = v.Encode()

	h := u.String()
	h = strings.TrimSuffix(h, "?")
	h = strings.TrimSuffix(h, "/")

	return h, nil
}
