package model

import "time"

type Stat struct {
	Timestamp time.Time
	Count     int
}
