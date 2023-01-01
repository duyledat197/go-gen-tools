package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.coffee..coffee", false},
		{"foo@bar.bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		actual := IsEmail(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsEmail(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.ORG", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"ftp.foo.bar", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"http://foobar.com/a-", true},
		{"http://foobar.پاکستان/", true},
		{"http://foobar.c_o_m", false},
		{"http://_foobar.com", false},
		{"http://foo_bar.com", true},
		{"http://user:pass@foo_bar_bar.bar_foo.com", true},
		{"", false},
		{"xyz://foobar.com", false},
		// {"invalid.", false}, is it false like "localhost."?
		{".com", false},
		{"rtmp://foobar.com", false},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", false},
		{"http://www.foo---bar.com/", false},
		{"http://r6---snnvoxuioq6.googlevideo.com", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", false},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
		{"http://foo^bar.org", false},
		{"http://foo&*bar.org", false},
		{"http://foo&bar.org", false},
		{"http://foo bar.org", false},
		{"http://foo.bar.org", true},
		{"http://www.foo.bar.org", true},
		{"http://www.foo.co.uk", true},
		{"foo", false},
		{"http://.foo.com", false},
		{"http://,foo.com", false},
		{",foo.com", false},
		{"http://myservice.:9093/", true},
		// according to issues #62 #66
		{"https://pbs.twimg.com/profile_images/560826135676588032/j8fWrmYY_normal.jpeg", true},
		// according to #125
		{"http://prometheus-alertmanager.service.q:9093", true},
		{"aio1_alertmanager_container-63376c45:9093", true},
		{"https://www.logn-123-123.url.with.sigle.letter.d:12345/url/path/foo?bar=zzz#user", true},
		{"http://me.example.com", true},
		{"http://www.me.example.com", true},
		{"https://farm6.static.flickr.com", true},
		{"https://zh.wikipedia.org/wiki/Wikipedia:%E9%A6%96%E9%A1%B5", true},
		{"google", false},
		// According to #87
		{"http://hyphenated-host-name.example.co.in", true},
		{"http://cant-end-with-hyphen-.example.com", false},
		{"http://-cant-start-with-hyphen.example.com", false},
		{"http://www.domain-can-have-dashes.com", true},
		{"http://m.abcd.com/test.html", true},
		{"http://m.abcd.com/a/b/c/d/test.html?args=a&b=c", true},
		{"http://[::1]:9093", true},
		{"http://[::1]:909388", false},
		{"1200::AB00:1234::2552:7777:1313", false},
		{"http://[2001:db8:a0b:12f0::1]/index.html", true},
		{"http://[1200:0000:AB00:1234:0000:2552:7777:1313]", true},
		{"http://user:pass@[::1]:9093/a/b/c/?a=v#abc", true},
		{"https://127.0.0.1/a/b/c?a=v&c=11d", true},
		{"https://foo_bar.example.com", true},
		{"http://foo_bar.example.com", true},
		{"http://foo_bar_fizz_buzz.example.com", true},
		{"http://_cant_start_with_underescore", false},
		{"http://cant_end_with_underescore_", false},
		{"foo_bar.example.com", true},
		{"foo_bar_fizz_buzz.example.com", true},
		{"http://hello_world.example.com", true},
		// According to #212
		{"foo_bar-fizz-buzz:1313", true},
		{"foo_bar-fizz-buzz:13:13", false},
		{"foo_bar-fizz-buzz://1313", false},
	}
	for _, test := range tests {
		actual := IsURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsSHA256(t *testing.T) {
	assert.False(t, IsSHA256("this is not a sha256"))
	assert.False(t, IsSHA256("ABCDEF2345789abc"))
	assert.False(t, IsSHA256("aaa bbb ccc ddd eee fff ggg hhh iii jjj kkk lll mmm nnn ooo ppp "))

	h := sha256.Sum256([]byte{8, 8, 8, 8, 8, 8, 8, 8, 8})
	assert.True(t, IsSHA256(hex.EncodeToString(h[:])))
}
