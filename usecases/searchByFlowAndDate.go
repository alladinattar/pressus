package usecases

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (s *service) SearchByFlowAndDate(flow, from, until string) ([]string, error) {
	dateFrom, err := time.Parse("2006-01-02", from)
	if err != nil {
		log.Error("Failed parse date:", err.Error())
	}

	dateUntil, err := time.Parse("2006-01-02", until)
	if err != nil {
		log.Error("Failed parse date:", err.Error())
	}

	log.Debug("From: ", dateFrom.String(), "Until: ", dateUntil.String())
	return s.searchEngine.FindByFlowAndDate(flow, dateFrom, dateUntil)

}
