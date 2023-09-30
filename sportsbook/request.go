package sportsbook

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var loginPostAddress = "https://www.sportsbook.ag/cca/customerauthn/pl/login"
var wagerGetAddress = "https://www.sportsbook.ag/sbk/sportsbook4/history.sbk"

func Login() bool {
	loginPostUrl, _ := url.Parse(loginPostAddress)

	username := os.Getenv("SPORTSBOOK_USERNAME")
	password := os.Getenv("SPORTSBOOK_PASSWORD")
	blackbox := os.Getenv("SPORTSBOOK_BLACKBOX")

	postData := url.Values{}
	postData.Set("username", username)
	postData.Set("password", password)
	postData.Set("blackbox", blackbox)
	postData.Set("service", "/sbk/sportsbook4/home.sbk")
	postData.Set("login_fail", "/ctr/acctmgt/pl/rui/login.ctr?service=/sbk/sportsbook4/home.sbk")
	postData.Set("sp_casinoid", "7000")
	postData.Set("sp_siteID", "7000")
	postData.Set("lp_casinoid", "7000")

	log.Debug().Msgf("postData: %#v", postData)

	loginRequest, _ := http.NewRequest(http.MethodPost, loginPostUrl.String(), strings.NewReader(postData.Encode()))
	loginRequest.Header.Set("path", "/cca/customerauthn/pl/login")
	loginRequest.Header.Set("authority", "www.sportsbook.ag")
	loginRequest.Header.Set("Content-type", "application/x-www-form-urlencoded")

	loginResp, err := client.Do(loginRequest)
	if err != nil {
		log.Fatal().Msgf("Error logging in. %s", err)
	}

	log.Info().Msgf("Login response code: %d", loginResp.StatusCode)

	redirectUrlRaw := loginResp.Header.Get("Location")
	redirectUrl, _ := url.Parse(redirectUrlRaw)
	log.Debug().Msgf("Redirect URL: %s", redirectUrl)
	for _, cookie := range jar.Cookies(redirectUrl) {
		log.Debug().Msgf("Login Cookie: %s => %s", cookie.Name, cookie.Value)
	}

	body, err := io.ReadAll(loginResp.Body)
	if err != nil {
		log.Fatal().Msgf("Error reading the body. %s", err)
	}
	// log.Debug().Msgf("Login Body: %s", string(body))

	// The presence of "My Bets" indicates that the login was successful
	return strings.Contains(string(body), "My Bets")
}

func GetRecentWagerData() (string, error) {
	return GetWagerDataByPage(1, "")
}

// date format: mm/dd/yyyy
func GetWagerDataByPage(page int, startDate string) (string, error) {
	wagerDataUrl, _ := url.Parse(wagerGetAddress)

	if page <= 0 {
		page = 0
	}

	postData := url.Values{}
	postData.Set("betState", "0")
	postData.Set("searchByDateType", "1")
	postData.Set("dateRangeMode", "CUSTOM")
	postData.Set("customDateRangeDays", "30")
	postData.Set("customDateRangeDirection", "2")
	postData.Set("page", strconv.Itoa(page))

	if len(startDate) == 0 {
		startDate = time.Now().Format("01/02/2006")
	}
	log.Info().Msgf("GetWagerData: %s - 30 days", startDate)
	postData.Set("customDateRangeDate", startDate)

	log.Debug().Msgf("GetWagerData: %#v", postData)

	wagerDataRequest, _ := http.NewRequest(http.MethodPost, wagerDataUrl.String(), strings.NewReader(postData.Encode()))
	wagerDataRequest.Header.Set("Content-type", "application/x-www-form-urlencoded")

	wagerDataResp, err := client.Do(wagerDataRequest)
	if err != nil {
		log.Error().Msgf("Error logging in. %s", err)
		return "", nil
	}

	body, err := io.ReadAll(wagerDataResp.Body)
	if err != nil {
		log.Error().Msgf("Error reading the body. %s", err)
		return "", nil
	}
	return string(body), nil
}
