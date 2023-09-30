package sportsbook

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/sportsbook-data/lib"
)

func init() {
	lib.SetLogLevelTest()
}

func TestWagerData(t *testing.T) {
	htmlData, err := os.ReadFile("../test_data/recent_wager_data.html")
	if err != nil {
		log.Fatal().Msgf("Error opening file. %s", err)
	}
	log.Info().Msg("Test log message")
	err = ParseWagerData(string(htmlData))
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
