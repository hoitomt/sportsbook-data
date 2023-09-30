package sportsbook

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/publicsuffix"
)

var client http.Client
var jar http.CookieJar

func InitJar() error {
	var err error
	jar, err = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}
	client = http.Client{
		Jar: jar,
	}

	return nil
}

func PrintJar(url *url.URL) {
	for _, cookie := range jar.Cookies(url) {
		log.Debug().Msgf("Login Cookie: %s => %s", cookie.Name, cookie.Value)
	}
}
