package models

import "time"

type HeartBeat struct {
	JobId int
	time  time.Time
}
