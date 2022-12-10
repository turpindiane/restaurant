<h1 align="center">
  Restaurant API Project in Go
</h1>

## Site purpose

This site queries a PostgreSQL database containing restaurant menu items and calculates a bill.

## Building Blocks

This site was created in GoLang with a PostgreSQL database. This site uses Gin for creating a router.

## Running locally

### Installations

In order to run locally, your machine must have Go and PostgreSQL installed and running.

### Database Configuration

To set your database credentials, export the following variables:

```shell
export APP_DB_USERNAME=
export APP_DB_NAME=
export APP_DB_PASSWORD=
```

To create and populate the table and enum required by this project, run the up [migrations](https://github.com/turpindiane/restaurant/tree/main/migrations) in this repository.

### Starting the Server

```shell
go run .
```

### Running in Postman

For easily calling the endpoint in Postman, import the [collection](https://github.com/turpindiane/restaurant/blob/main/restaurant.postman_collection.json) in this repository.

## Database Schema

table public.dishes
| Column | Type |
| ----------- | ----------- |
| id | serial |
| title | varchar(255) |
| dish_course | course |
| price | decimal |

enum courses
| Values |
| ----------- |
| 'appetizer' |
| 'entree' |
| 'dessert' |
| 'beverage' |

## Endpoints

| URL             | Description                                                                                                                                                  |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| GET /dishes     | retrieves all dishes in the public.dishes table                                                                                                              |
| GET /appetizers | retrieves all appetizers in the public.dishes table                                                                                                          |
| GET /entrees    | retrieves all entrees in the public.dishes table                                                                                                             |
| GET /desserts   | retrieves all desserts in the public.dishes table                                                                                                            |
| GET /beverages  | retrieves all beverages in the public.dishes table                                                                                                           |
| GET /calculate  | calculates a bill total based on the provided dishes by id. Required query parameter is entree, option query parameters are appetizer, dessert, and beverage |
