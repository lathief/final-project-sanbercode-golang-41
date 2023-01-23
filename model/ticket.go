package model

type Ticket struct {
	TicketID   int
	UserID     int
	ScheduleID int
	SeatID     int
	StudioID   int
	FilmCode   string
}

type TicketInput struct {
	Username     string
	Email        string
	MobileNumber string
	NumberSeat   string
	StartTime    string
	FilmName     string
	Date         string
}

type TicketOutput struct {
	Username    string
	SeatOrder   Seat
	FilmOrder   Film
	StudioOrder Studio
}
