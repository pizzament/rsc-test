package stats_handler

import (
	"time"
)

type StatsResponse struct {
	Stats []StatItem `json:"stats"`
}

type StatItem struct {
	TS time.Time `json:"ts"`
	V  int       `json:"v"`
}
