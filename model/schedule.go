package model

import (
	"time"
)

type Schedule struct {
	ScheduleID   int
	FilmCode     string
	ScheduleDate time.Time
	StartTime    time.Time
	EndTime      time.Time
	Price        string
	Duration     int
	StudioID     int
}

type ScheduleInput struct {
	FilmCode  string
	Date      string
	StartTime string
	Duration  int
	Price     string
	StudioID  int
}
