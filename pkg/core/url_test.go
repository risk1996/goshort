package core

import "testing"

type Case struct {
	Input    string
	Expected string
}

// REF: https://github.com/sekimura/go-normalize-url/blob/master/normalizeurl_test.go
func TestNormalizeURL(t *testing.T) {
	tests := []Case{
		/* [ 0] */ {"www.sekimura.org", "http://www.sekimura.org"},
		/* [ 1] */ {"sekimura.org", "http://sekimura.org"},
		/* [ 2] */ {"HTTP://sekimura.org", "http://sekimura.org"},
		/* [ 3] */ {"//sekimura.org", "http://sekimura.org"},
		/* [ 4] */ {"http://sekimura.org", "http://sekimura.org"},
		/* [ 5] */ {"http://sekimura.org:80", "http://sekimura.org"},
		/* [ 6] */ {"https://sekimura.org:443", "https://sekimura.org"},
		/* [ 7] */ {"ftp://sekimura.org:21", "ftp://sekimura.org"},
		/* [ 8] */ {"http://www.sekimura.org", "http://www.sekimura.org"},
		/* [ 9] */ {"www.sekimura.org", "http://www.sekimura.org"},
		/* [10] */ {"http://sekimura.org/foo/", "http://sekimura.org/foo"},
		/* [11] */ {"sekimura.org/?foo=bar%20baz%3F", "http://sekimura.org/?foo=bar+baz%3F"},
		/* [12] */ {"http://sekimura.org/?", "http://sekimura.org"},
		/* [13] */ {"http://xn--xample-hva.com", "http://%C3%AAxample.com"},
		/* [14] */ {"http://sekimura.org/?b=bar&a=foo", "http://sekimura.org/?a=foo&b=bar"},
	}

	for i := range tests {
		actual, err := NormalizeURL(tests[i].Input)
		if err != nil {
			t.Errorf("[%2d] NormalizeURL(\"%v\") failed: %v", i, tests[i].Input, err)
		}
		if actual != tests[i].Expected {
			t.Errorf("[%2d] NormalizeURL(\"%v\") produces %v whereas %v is expected", i, tests[i].Input, actual, tests[i].Expected)
		}
	}
}
