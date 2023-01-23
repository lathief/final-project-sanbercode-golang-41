package repository

import (
	"bioskop/model"
	"database/sql"
)

func SaveSeat(db *sql.DB, seat string) (err error) {
	sql := `
	INSERT INTO seat(numberseat)
	VALUES ($1);
	`
	errs := db.QueryRow(sql, seat)
	return errs.Err()
}
func FindSeat(db *sql.DB, numberseat string) (seatID int, err error) {
	var seat model.Seat
	sql := `SELECT * FROM seat WHERE numberseat=$1`
	rows, err := db.Query(sql, numberseat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&seat.SeatID, &seat.NumberSeat)
		if err != nil {
			panic(err)
		}
		seatID = seat.SeatID
	}
	return
}

func GetSeatById(db *sql.DB, id int) (seat model.Seat, err error) {
	sql := `SELECT * FROM seat WHERE seatid=$1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&seat.SeatID, &seat.NumberSeat)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetStudioById(db *sql.DB, id int) (seat model.Studio, err error) {
	sql := `SELECT * FROM studio WHERE studioid=$1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&seat.StudioID, &seat.StudioName)
		if err != nil {
			panic(err)
		}
	}
	return
}
