package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

// validates that an entry exists in the dishes table with the provided id and course
func validateCourse(db *sql.DB, id string, c string) (dish, int, error) {
	// validate that id is a valid integer
	i, err := strconv.Atoi(id)
	if err != nil {
		return dish{}, http.StatusBadRequest, fmt.Errorf("The id for %s must be a valid integer", c)
	}

	// retrieve dish from the db
	d, err := queryDishById(db, i)
	if err == sql.ErrNoRows {
		return dish{}, http.StatusBadRequest, fmt.Errorf("There is no dish associated with id %s", id)
	} else if err != nil {
		return dish{}, http.StatusInternalServerError, err
	}
	if d.Course != c {
		return dish{}, http.StatusBadRequest, fmt.Errorf("There is no %s associated with id %s", c, id)
	}
	return d, http.StatusOK, nil
}
