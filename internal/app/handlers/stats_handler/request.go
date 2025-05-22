package stats_handler

import "time"

type StatsRequest struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
