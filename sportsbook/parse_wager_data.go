package sportsbook

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
)

type Ticket struct {
	ID          string
	Description string
	WagerDate   time.Time
	EventDate   time.Time
	BetAmt      float32
	BetWinAmt   float32
	BetResult   string
}

func ParseWagerData(html string) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Error().Err(err)
		return err
	}
	// ticketsRaw := doc.Find("#searchResult").Text()
	// log.Info().Msgf("Tickets: %s", ticketsRaw)

	doc.Find("#searchResult .panel").Each(func(i int, s *goquery.Selection) {
		betIdRaw := s.Find(".betId").Text()
		betDescRaw := s.Find("#betDesc").Text()
		betDateRaw := s.Find("#betDate").Text()
		betAmtRaw := s.Find("#betAmt").Text()
		betWinAmtRaw := s.Find("#betWinAmt").Text()
		betResultRaw := s.Find("#betResult").Text()
		team1 := s.Find("#team1").Text()
		fnScore1 := s.Find("#fnScore1").Text()
		team2 := s.Find("#team2").Text()
		fnScore2 := s.Find("#fnScore2").Text()
		eventTimeRaw := s.Find("#eventTime").Text()
		betDetailsRaw := s.Find("#market").Text()

		r, _ := regexp.Compile(`\d{2}/\d{2}/\d{2} \d{2}:\d{2}`)
		betDate := r.FindString(betDateRaw)
		eventTime := r.FindString(eventTimeRaw)
		betDetails := strings.TrimSpace(betDetailsRaw)

		log.Info().Msg("New Ticket")
		log.Info().Msgf("Ticket ID: %s", betIdRaw)
		log.Info().Msgf("Ticket Desciption: %s", betDescRaw)
		log.Info().Msgf("Ticket Date: %s", betDate)
		log.Info().Msgf("Ticket Bet Amount: %s", betAmtRaw)
		log.Info().Msgf("Ticket Bet Win Amount: %s", betWinAmtRaw)
		log.Info().Msgf("Ticket Result: %s", betResultRaw)
		log.Info().Msgf("Ticket team1: %s", team1)
		log.Info().Msgf("Ticket fnScore1: %s", fnScore1)
		log.Info().Msgf("Ticket team2: %s", team2)
		log.Info().Msgf("Ticket fnScore2: %s", fnScore2)
		log.Info().Msgf("Ticket eventTime: %s", eventTime)
		log.Info().Msgf("Ticket betDetails: %s", betDetails)
	})
	return nil
}
