package fasthttp

import (
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

type (
	// Cookie implements `engine.Cookie`.
	Cookie struct {
		*fasthttp.Cookie
	}
)

// Name implements `engine.Cookie#Name` function.
func (c *Cookie) Name() string {
	return string(c.Cookie.Key())
}

// Value implements `engine.Cookie#Value` function.
func (c *Cookie) Value() string {
	return string(c.Cookie.Value())
}

// Path implements `engine.Cookie#Path` function.
func (c *Cookie) Path() string {
	return string(c.Cookie.Path())
}

// Domain implements `engine.Cookie#Domain` function.
func (c *Cookie) Domain() string {
	return string(c.Cookie.Domain())
}

// Expires implements `engine.Cookie#Expires` function.
func (c *Cookie) Expires() time.Time {
	return c.Cookie.Expire()
}

// Secure implements `engine.Cookie#Secure` function.
func (c *Cookie) Secure() bool {
	return c.Cookie.Secure()
}

// HTTPOnly implements `engine.Cookie#HTTPOnly` function.
func (c *Cookie) HTTPOnly() bool {
	return c.Cookie.HTTPOnly()
}

// SameSite implements `engine.Cookie#SameSite` function.
func (c *Cookie) SameSite() http.SameSite {
	return sameSiteFastHttp2Http(c.Cookie.SameSite())
}

var (
	sameSiteHttp2FastHttpMap = map[http.SameSite]fasthttp.CookieSameSite{
		http.SameSiteNoneMode:    fasthttp.CookieSameSiteDisabled,
		http.SameSiteLaxMode:     fasthttp.CookieSameSiteLaxMode,
		http.SameSiteStrictMode:  fasthttp.CookieSameSiteStrictMode,
		http.SameSiteDefaultMode: fasthttp.CookieSameSiteDefaultMode,
	}
	sameSiteFastHttp2HttpMap map[fasthttp.CookieSameSite]http.SameSite
)

func init() {
	sameSiteFastHttp2HttpMap = make(map[fasthttp.CookieSameSite]http.SameSite)
	for httpSameSite, fasthttpSameSite := range sameSiteHttp2FastHttpMap {
		sameSiteFastHttp2HttpMap[fasthttpSameSite] = httpSameSite
	}
}

func sameSiteHttp2FastHttp(httpSameSite http.SameSite) fasthttp.CookieSameSite {
	if fasthttpSameSite, ok := sameSiteHttp2FastHttpMap[httpSameSite]; ok {
		return fasthttpSameSite
	}
	return fasthttp.CookieSameSiteDefaultMode
}

func sameSiteFastHttp2Http(fasthttpSameSite fasthttp.CookieSameSite) http.SameSite {
	if httpSameSite, ok := sameSiteFastHttp2HttpMap[fasthttpSameSite]; ok {
		return httpSameSite
	}
	return http.SameSiteDefaultMode
}
