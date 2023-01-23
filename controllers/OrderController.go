package controllers

import (
	"bioskop/database"
	"bioskop/model"
	"bioskop/repository"
	"bioskop/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderTicket(c *gin.Context) {
	var order model.TicketInput
	var ticket model.Ticket
	err := c.ShouldBindJSON(&order)
	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket.UserID = user.ID
	err = repository.SaveSeat(database.DbConnection, order.NumberSeat)
	if err != nil {
		panic(err)
	}
	ticket.SeatID, err = repository.FindSeat(database.DbConnection, order.NumberSeat)
	if err != nil {
		panic(err)
	}
	film, err := repository.GetFilmByName(database.DbConnection, order.FilmName)
	ticket.FilmCode = film.FilmCode
	if err != nil {
		panic(err)
	}
	sch, err := repository.GetScheduleByFilmCode(database.DbConnection, film.FilmCode)
	var date, starttime time.Time
	date, _ = time.Parse(layoutDate, order.Date)
	starttime, _ = time.Parse(layoutTime, order.StartTime)
	for i := 0; i < len(sch); i++ {
		fmt.Println(sch[i].ScheduleDate.Format(time.RFC822Z))
		fmt.Println(date.Format(time.RFC822Z))
		fmt.Println(sch[i].FilmCode)
		fmt.Println(ticket.FilmCode)
		if sch[i].ScheduleDate.Format(time.RFC822Z) == date.Format(time.RFC822Z) && sch[i].StartTime == starttime && sch[i].FilmCode == film.FilmCode {
			ticket.ScheduleID = sch[i].ScheduleID
			ticket.StudioID = sch[i].StudioID
			err = repository.SaveTicket(database.DbConnection, ticket)
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{"result": "Success Order Ticket"})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "Tidak ditemukan jadwalnya"})
		}
	}
}

func ShowTicket(c *gin.Context) {
	var Invoice model.TicketOutput
	user, err := utils.CurrentUser(c)
	if err != nil {
		panic(err)
	}
	ticket, err := repository.GetTicketById(database.DbConnection, user.CustomerID)
	if err != nil {
		panic(err)
	}
	var schedule []model.Schedule
	Invoice.Username = user.Username
	Invoice.FilmOrder, err = repository.GetFilmByCode(database.DbConnection, ticket.FilmCode)
	schedule, err = repository.GetScheduleByFilmCode(database.DbConnection, ticket.FilmCode)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(schedule), len(Invoice.FilmOrder.Schedule))
	for i := 0; i < len(schedule); i++ {
		if ticket.ScheduleID == schedule[i].ScheduleID {
			Invoice.FilmOrder.Schedule = append(Invoice.FilmOrder.Schedule, schedule[i])
		}
	}
	if err != nil {
		panic(err)
	}
	Invoice.SeatOrder, err = repository.GetSeatById(database.DbConnection, ticket.SeatID)
	if err != nil {
		panic(err)
	}
	Invoice.StudioOrder, err = repository.GetStudioById(database.DbConnection, ticket.StudioID)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": Invoice})
}
