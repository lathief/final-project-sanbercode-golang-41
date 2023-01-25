package controllers

import (
	"bioskop/database"
	"bioskop/model"
	"bioskop/repository"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFilm(c *gin.Context) {
	var schedule []model.Schedule
	var (
		result gin.H
	)
	film, err := repository.GetAllFilm(database.DbConnection)

	for i := 0; i < len(film); i++ {
		schedule, _ = repository.GetScheduleByFilmCode(database.DbConnection, film[i].FilmCode)
		film[i].Schedule = schedule
	}
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": film,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertFilm(c *gin.Context) {
	var film model.Film
	err := c.ShouldBindJSON(&film)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	err = repository.InsertFilm(database.DbConnection, film)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Film"})
}

func UpdateFilm(c *gin.Context) {
	var film model.Film
	var filmTmp model.Film
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&film)
	if err != nil {
		panic(err)
	}
	filmTmp, err = repository.GetFilmById(database.DbConnection, id)
	if err != nil {
		panic(err)
	}
	if filmTmp.FilmID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Not Found FilmID " + c.Param("id"),
		})
	} else {
		err = repository.UpdateFilm(database.DbConnection, film, c.Param("id"))
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": "Success Update FilmID " + c.Param("id"),
		})
	}

}

func DeleteFilm(c *gin.Context) {
	var filmTmp model.Film
	id, err := strconv.Atoi(c.Param("id"))
	filmTmp, err = repository.GetFilmById(database.DbConnection, id)
	if err != nil {
		panic(err)
	}
	if filmTmp.FilmID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Not Found FilmID " + c.Param("id"),
		})
	} else {
		err = repository.DeleteFilm(database.DbConnection, id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": "Success Update FilmID " + c.Param("id"),
		})
	}
}

func GetFilmById(c *gin.Context) {
	var schedule []model.Schedule
	var (
		result gin.H
	)
	id, err := strconv.Atoi(c.Param("id"))
	film, err := repository.GetFilmById(database.DbConnection, id)
	schedule, _ = repository.GetScheduleByFilmCode(database.DbConnection, film.FilmCode)
	film.Schedule = schedule
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": film,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetFilmByNameStatus(c *gin.Context) {
	var schedule []model.Schedule
	var (
		result gin.H
	)
	name := c.Query("name")
	status := c.DefaultQuery("status", "OnAir")
	film, err := repository.GetFilmByKeyword(database.DbConnection, name, status)
	for i := 0; i < len(film); i++ {
		schedule, _ = repository.GetScheduleByFilmCode(database.DbConnection, film[i].FilmCode)
		film[i].Schedule = schedule
	}
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": film,
		}
	}

	c.JSON(http.StatusOK, result)
}
