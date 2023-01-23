package model

type Film struct {
	FilmID     int
	FilmName   string
	FilmStatus string
	FilmCode   string
	Schedule   []Schedule
}

func (film *Film) InsertSchedule(sch []Schedule) []Schedule {
	film.Schedule = sch
	return film.Schedule
}
