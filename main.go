package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sportsbook-data/lib"
	"github.com/sportsbook-data/sportsbook"
)

func main() {
	htmlBody, err := sportsbook.GetRecentWagerData()
	if err != nil {
		log.Fatal().Msgf("ERROR requesting wager data: %s", err)
	}
	log.Debug().Msg(htmlBody)
	log.Info().Msg("Write recent wager data to test_data/recent_wager_data.html")
	os.WriteFile("test_data/recent_wager_data.html", []byte(htmlBody), 0644)
}

func init() {
	lib.SetLogLevel()

	err := sportsbook.InitJar()
	if err != nil {
		log.Fatal().Msgf("Error instantiating the cookejar. %s", err)
	}
	loginSuccess := sportsbook.Login()

	if loginSuccess {
		log.Info().Msg("Login Successful")
	} else {
		log.Fatal().Msg("Login failed")
	}
}
