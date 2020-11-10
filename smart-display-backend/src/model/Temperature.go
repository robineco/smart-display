package model

import "time"

type Temperature struct {
	ID       int
	Temp     float64
	Time     time.Time
	Location string
}
