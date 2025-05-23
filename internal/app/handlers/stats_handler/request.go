package stats_handler

import (
	"fmt"
	"time"
)

type StatsRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (sr StatsRequest) ParseTimes() (from, to time.Time, err error) {
	layout := "2006-01-02T15:04:05"

	if from, err = time.Parse(layout, sr.From); err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid 'from' time: %w", err)
	}

	if to, err = time.Parse(layout, sr.To); err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid 'to' time: %w", err)
	}

	return from, to, nil
}
