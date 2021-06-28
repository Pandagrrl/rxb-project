package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var (
	db         *sql.DB
	filmSelect = "SELECT f.film_id, f.title, f.description, f.release_year, l.name, f.rental_duration, f.rental_rate, f.length, f.replacement_cost, f.rating , f.last_update, f.special_features FROM film f INNER JOIN language l ON f.language_id = l.language_id"
)

type (
	FilmCategory struct {
		FilmID     uint32    `json:"film_id"`
		CategoryID uint32    `json:"category_id"`
		LastUpdate time.Time `json:"last_updated"`
	}

	Film struct {
		FilmID          uint32    `json:"film_id"`
		Title           string    `json:"title"`
		Description     string    `json:"description"`
		ReleaseYear     string    `json:"release_year"`
		Language        string    `json:"name"`
		RentalDuration  uint32    `json:"rental_duration"`
		RentalRate      float32   `json:"rental_rate"`
		Length          uint32    `json:"length"`
		ReplacementCost float32   `json:"replacement_cost"`
		Rating          string    `json:"rating"`
		LastUpdate      time.Time `json:"last_update"`
		SpecialFeatures []uint8   `json:"special_features"`
	}
)

func InitDB(dataSourceName string) error {
	var err error

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	return db.Ping()
}

func AllFilmCategories() ([]FilmCategory, error) {
	// This now uses the unexported global variable.
	rows, err := db.Query("SELECT * FROM film_category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []FilmCategory

	for rows.Next() {
		var category FilmCategory

		err := rows.Scan(&category.FilmID, &category.CategoryID, &category.LastUpdate)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func AllFilms() ([]Film, error) {
	rows, err := db.Query("SELECT f.film_id, f.title, f.description, f.release_year, l.name, f.rental_duration, f.rental_rate, f.length, f.replacement_cost, f.rating , f.last_update, f.special_features FROM film f INNER JOIN language l ON f.language_id = l.language_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var films []Film

	for rows.Next() {
		var film Film

		err := rows.Scan(&film.FilmID, &film.Title, &film.Description, &film.ReleaseYear, &film.Language, &film.RentalDuration, &film.RentalRate, &film.Length, &film.ReplacementCost, &film.Rating, &film.LastUpdate, &film.SpecialFeatures)
		if err != nil {
			return nil, err
		}

		films = append(films, film)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}

func SearchFilms(title string) ([]Film, error) {
	selectSql := filmSelect + " where title like '%' || $1 || '%'"
	rows, err := db.Query(selectSql, title)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var films []Film

	for rows.Next() {
		var film Film

		err := rows.Scan(&film.FilmID, &film.Title, &film.Description, &film.ReleaseYear, &film.Language, &film.RentalDuration, &film.RentalRate, &film.Length, &film.ReplacementCost, &film.Rating, &film.LastUpdate, &film.SpecialFeatures)
		if err != nil {
			return nil, err
		}

		films = append(films, film)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}
