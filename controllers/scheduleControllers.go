package controllers

import (
	"bioskop/database"
	"bioskop/model"
	"bioskop/repository"
	"fmt"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

var layoutDate string = "2006-01-02"
var layoutTime string = "15:04:05"

func GetAllSchedule(c *gin.Context) {
	var (
		result gin.H
	)
	schedule, err := repository.GetAllSchedule(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"schedule": schedule,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertSchedule(c *gin.Context) {
	var schedule model.ScheduleInput
	var sch model.Schedule
	err := c.ShouldBindJSON(&schedule)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	sch.ScheduleDate, _ = time.Parse(layoutDate, schedule.Date)
	sch.Duration = schedule.Duration
	sch.StartTime, _ = time.Parse(layoutTime, schedule.StartTime)
	fmt.Println(sch.StartTime, schedule.StartTime)
	sch.EndTime = sch.StartTime.Add(time.Minute * time.Duration(sch.Duration))
	sch.FilmCode = schedule.FilmCode
	sch.Price = schedule.Price
	sch.StudioID = schedule.StudioID
	err = repository.InsertSchedule(database.DbConnection, sch)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"schedule": "Success Insert Schedule"})
}

func UpdateSchedule(c *gin.Context) {
	var schedule model.ScheduleInput
	var sch model.Schedule
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&schedule)
	if err != nil {
		panic(err)
	}
	sch, err = repository.GetScheduleById(database.DbConnection, id)
	if sch.ScheduleID != 0 {
		sch.ScheduleDate, _ = time.Parse(layoutDate, schedule.Date)
		sch.Duration = schedule.Duration
		sch.StartTime, _ = time.Parse(layoutTime, schedule.StartTime)
		sch.EndTime = sch.StartTime.Add(time.Minute * time.Duration(sch.Duration))
		sch.FilmCode = schedule.FilmCode
		sch.Price = schedule.Price
		sch.StudioID = schedule.StudioID

		err = repository.UpdateSchedule(database.DbConnection, sch, c.Param("id"))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"schedule": "Success Update ScheduleID " + c.Param("id"),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"schedule": "Not Found ScheduleID " + c.Param("id"),
		})
	}
}

func DeleteSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var sch model.Schedule
	sch, err = repository.GetScheduleById(database.DbConnection, id)
	if sch.ScheduleID != 0 {
		err = repository.DeleteSchedule(database.DbConnection, id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"schedule": "Success Delete ScheduleID " + c.Param("id"),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"schedule": "Not Found ScheduleID " + c.Param("id"),
		})
	}

}

func GetScheduleByFilmCode(c *gin.Context) {
	var (
		result gin.H
	)
	code := c.Param("code")

	schedule, err := repository.GetScheduleByFilmCode(database.DbConnection, code)
	film, err := repository.GetFilmByCode(database.DbConnection, code)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"film":     film,
			"schedule": schedule,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetScheduleById(c *gin.Context) {
	var (
		result gin.H
	)
	id, _ := strconv.Atoi(c.Param("id"))
	schedule, err := repository.GetScheduleById(database.DbConnection, id)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"schedule": schedule,
		}
	}
	c.JSON(http.StatusOK, result)
}
