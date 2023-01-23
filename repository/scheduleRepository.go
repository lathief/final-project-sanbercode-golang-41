package repository

import (
	"bioskop/model"
	"database/sql"
)

func GetAllSchedule(db *sql.DB) (result []model.Schedule, err error) {
	sql := `SELECT * FROM schedule`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var schedule = model.Schedule{}

		err = rows.Scan(&schedule.ScheduleID, &schedule.FilmCode, &schedule.ScheduleDate, &schedule.StartTime, &schedule.EndTime, &schedule.Price, &schedule.Duration, &schedule.StudioID)
		if err != nil {
			panic(err)
		}

		result = append(result, schedule)
	}
	return
}

func InsertSchedule(db *sql.DB, schedule model.Schedule) (err error) {
	sql := `
	INSERT INTO schedule(
	filmcode, scheduledate, starttime, endtime, price, duration, studioid)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	errs := db.QueryRow(sql, schedule.FilmCode, schedule.ScheduleDate, schedule.StartTime, schedule.EndTime, schedule.Price, schedule.Duration, schedule.StudioID)
	return errs.Err()
}

func UpdateSchedule(db *sql.DB, schedule model.Schedule, id string) (err error) {
	sql := `
	UPDATE schedule
	SET filmcode=$2, scheduledate=$3, starttime=$4, endtime=$5, price=$6, duration=$7, studioid=$8
	WHERE scheduleid=$1;
	`
	errs := db.QueryRow(sql, id, schedule.FilmCode, schedule.ScheduleDate, schedule.StartTime, schedule.EndTime, schedule.Price, schedule.Duration, schedule.StudioID)
	return errs.Err()
}

func DeleteSchedule(db *sql.DB, id int) (err error) {
	sql := `DELETE FROM schedule WHERE scheduleid = $1`
	errs := db.QueryRow(sql, id)
	return errs.Err()
}

func GetScheduleByFilmCode(db *sql.DB, code string) (result []model.Schedule, err error) {
	sql := `SELECT * FROM schedule WHERE filmcode = $1`
	rows, err := db.Query(sql, code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var schedule model.Schedule
	for rows.Next() {
		err = rows.Scan(&schedule.ScheduleID, &schedule.FilmCode, &schedule.ScheduleDate, &schedule.StartTime, &schedule.EndTime, &schedule.Price, &schedule.Duration, &schedule.StudioID)
		if err != nil {
			panic(err)
		}
		result = append(result, schedule)
	}
	return
}

func GetScheduleById(db *sql.DB, id int) (result model.Schedule, err error) {
	sql := `SELECT * FROM schedule WHERE scheduleid = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.ScheduleID, &result.FilmCode, &result.ScheduleDate, &result.StartTime, &result.EndTime, &result.Price, &result.Duration, &result.StudioID)
		if err != nil {
			panic(err)
		}
	}
	return
}
