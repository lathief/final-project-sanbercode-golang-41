package repository

import (
	"bioskop/model"
	"database/sql"
)

func SaveTicket(db *sql.DB, ticket model.Ticket) (err error) {
	sql := `
	INSERT INTO ticket(
	customerid, scheduleid, seatid, studioid, filmcode)
	VALUES ($1, $2, $3, $4, $5);
	`
	errs := db.QueryRow(sql, ticket.UserID, ticket.ScheduleID, ticket.SeatID, ticket.StudioID, ticket.FilmCode)
	return errs.Err()
}

func GetTicketById(db *sql.DB, id int) (result model.Ticket, err error) {
	sql := "SELECT * FROM ticket WHERE customerid = $1"
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.TicketID, &result.UserID, &result.ScheduleID, &result.SeatID, &result.StudioID, &result.FilmCode)
		if err != nil {
			panic(err)
		}
	}
	return
}
