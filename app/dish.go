package main

type dish struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Course string  `json:"course"`
	Price  float64 `json:"price"`
}

const (
	Appetizer = "appetizer"
	Entree    = "entree"
	Dessert   = "dessert"
	Beverage  = "beverage"
)
