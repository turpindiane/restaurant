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
