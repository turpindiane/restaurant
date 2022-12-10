package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	DB *sql.DB
}

// getDishes responds with all dishes
// GET /dishes
func (handler *handler) getDishes(c *gin.Context) {
	dishes, err := queryDishes(handler.DB)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, dishes)
	}
}

// getAppetizers responds with all appetizers
// GET /appetizers
func (handler *handler) getAppetizers(c *gin.Context) {
	dishes, err := queryDishesByCourse(handler.DB, Appetizer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, dishes)
	}
}

// getEntrees responds with all entrees
// GET /entrees
func (handler *handler) getEntrees(c *gin.Context) {
	dishes, err := queryDishesByCourse(handler.DB, Entree)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, dishes)
	}
}

// getDesserts responds with all desserts
// GET /desserts
func (handler *handler) getDesserts(c *gin.Context) {
	dishes, err := queryDishesByCourse(handler.DB, Dessert)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, dishes)
	}
}

// getBeverages responds with all beverages
// GET /beverages
func (handler *handler) getBeverages(c *gin.Context) {
	dishes, err := queryDishesByCourse(handler.DB, Beverage)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, dishes)
	}
}

// calculate returns the total of a bill based on the menu items selected
// a bill must contain an entree and can optionally include one beverage, one appetizer, and one dessert
// GET /calculate
func (handler *handler) calculate(c *gin.Context) {
	e := c.Query(Entree)
	a := c.Query(Appetizer)
	d := c.Query(Dessert)
	b := c.Query(Beverage)

	var total float64

	// validate and add entree
	entree, status, err := validateCourse(handler.DB, e, Entree)
	if err != nil {
		c.IndentedJSON(status, err.Error())
		return
	}
	total = total + entree.Price

	// validate and add appetizer if present
	if a != "" {
		appetizer, status, err := validateCourse(handler.DB, a, Appetizer)
		if err != nil {
			c.IndentedJSON(status, err.Error())
			return
		}
		total = total + appetizer.Price
	}

	// validate and add dessert if present
	if d != "" {
		dessert, status, err := validateCourse(handler.DB, d, Dessert)
		if err != nil {
			c.IndentedJSON(status, err.Error())
			return
		}
		total = total + dessert.Price
	}

	// validate and add beverage if present
	if b != "" {
		beverage, status, err := validateCourse(handler.DB, b, Beverage)
		if err != nil {
			c.IndentedJSON(status, err.Error())
			return
		}
		total = total + beverage.Price
	}

	c.IndentedJSON(http.StatusOK, total)
}
