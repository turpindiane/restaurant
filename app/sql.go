package main

import (
	"database/sql"
)

const (
	sqlGetDishes         = `select id, title, price, dish_course from public.dishes;`
	sqlGetDishesByCourse = `select id, title, price, dish_course from public.dishes where dish_course = $1;`
	sqlGetDishByID       = `select id, title, price, dish_course from public.dishes where id = $1;`
)

// retrieve all dishes from the dishes table
func queryDishes(db *sql.DB) ([]dish, error) {
	rows, err := db.Query(sqlGetDishes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes := []dish{}
	for rows.Next() {
		var d dish
		if err := rows.Scan(&d.ID, &d.Title, &d.Price, &d.Course); err != nil {
			return nil, err
		}
		dishes = append(dishes, d)
	}
	return dishes, nil
}

// retrieve all dishes associated with the given course from the dishes table
func queryDishesByCourse(db *sql.DB, c string) ([]dish, error) {
	rows, err := db.Query(sqlGetDishesByCourse, c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes := []dish{}
	for rows.Next() {
		var d dish
		if err := rows.Scan(&d.ID, &d.Title, &d.Price, &d.Course); err != nil {
			return nil, err
		}
		dishes = append(dishes, d)
	}
	return dishes, nil
}

// retrieves a single dish by id
func queryDishById(db *sql.DB, id int) (dish, error) {
	d := dish{}
	row := db.QueryRow(sqlGetDishByID, id)
	err := row.Scan(&d.ID, &d.Title, &d.Price, &d.Course)

	return d, err
}
