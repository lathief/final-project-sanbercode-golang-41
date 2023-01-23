package repository

import (
	"bioskop/model"
	"database/sql"
)

func GetAllFilm(db *sql.DB) (result []model.Film, err error) {
	sql := `SELECT * FROM film`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var film = model.Film{}

		err = rows.Scan(&film.FilmID, &film.FilmName, &film.FilmStatus, &film.FilmCode)
		if err != nil {
			panic(err)
		}

		result = append(result, film)
	}
	return
}

func InsertFilm(db *sql.DB, film model.Film) (err error) {
	sql := `
	INSERT INTO film(
	filmname, filmstatus, filmcode)
	VALUES ($1, $2, $3);
	`
	errs := db.QueryRow(sql, film.FilmName, film.FilmStatus, film.FilmCode)
	return errs.Err()
}

func UpdateFilm(db *sql.DB, film model.Film, id string) (err error) {
	sql := `
	UPDATE film
	SET filmname=$2, filmstatus=$3, filmcode=$4
	WHERE filmid=$1;
	`
	errs := db.QueryRow(sql, id, film.FilmName, film.FilmStatus, film.FilmCode)
	return errs.Err()
}

func DeleteFilm(db *sql.DB, id int) (err error) {
	sql := `DELETE FROM film WHERE filmid = $1`
	errs := db.QueryRow(sql, id)
	return errs.Err()
}

func GetFilmById(db *sql.DB, id int) (result model.Film, err error) {
	sql := `SELECT * FROM film WHERE filmid = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.FilmID, &result.FilmName, &result.FilmStatus, &result.FilmCode)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetFilmByCode(db *sql.DB, code string) (result model.Film, err error) {
	sql := `SELECT * FROM film WHERE filmcode = $1`
	rows, err := db.Query(sql, code)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.FilmID, &result.FilmName, &result.FilmStatus, &result.FilmCode)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetFilmByName(db *sql.DB, name string) (result model.Film, err error) {
	sql := `SELECT * FROM film WHERE filmname = $1`
	rows, err := db.Query(sql, name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.FilmID, &result.FilmName, &result.FilmStatus, &result.FilmCode)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetFilmByStatus(db *sql.DB, status string) (result []model.Film, err error) {
	sql := "SELECT * FROM film WHERE filmstatus=$1"
	rows, err := db.Query(sql, status)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var film model.Film
	for rows.Next() {
		err = rows.Scan(&film.FilmID, &film.FilmName, &film.FilmStatus, &film.FilmCode)
		if err != nil {
			panic(err)
		}
		result = append(result, film)
	}
	return
}

func GetFilmByKeyword(db *sql.DB, name string, status string) (result []model.Film, err error) {
	sql := "SELECT * FROM film WHERE filmname LIKE '" + name + "%'"
	sql = sql + "AND filmstatus='" + status + "'"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var film model.Film
	for rows.Next() {
		err = rows.Scan(&film.FilmID, &film.FilmName, &film.FilmStatus, &film.FilmCode)
		if err != nil {
			panic(err)
		}
		result = append(result, film)
	}
	return
}
