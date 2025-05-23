package stats_handler

import (
	"encoding/json"
	"time"
)

type StatsResponse struct {
	Stats []StatItem `json:"stats"`
}

type StatItem struct {
	TS time.Time `json:"ts"`
	V  int       `json:"v"`
}

func (si StatItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		TS string `json:"ts"`
		V  int    `json:"v"`
	}{
		TS: si.TS.Format("2006-01-02T15:04:05"),
		V:  si.V,
	})
}
