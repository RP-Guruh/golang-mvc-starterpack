# GOLANG MVC STARTER PACK

Hi gopher 🦫 

This is a simple repository for anyone who wants to use **Golang** for **Web API development**. The folder structure is inspired by my experience as a Laravel programmer, so it should feel familiar if you come.


## Package yang digunakan

 - [Godotenv](https://github.com/joho/godotenv)
 - [Gin Framework](https://github.com/go-playground/validator)
 - [MySQL Driver](https://github.com/go-sql-driver/mysql)
 - [Database Migration](https://github.com/golang-migrate/migrate)


## Table Schema

### table name : people

| Field | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `int (auto increment)` | Primary key |
| `first_name` | `varchar (255)` | |
| `last_name` | `varchar (255)` | |
| `place_of_birth` | `varchar (150)` ||
| `date_of_birth` | `date` | |
| `address` | `text` | |



## How to make migration file

- install golang-migrate cli : **go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest**
- running migrate file : **create -ext sql -dir database/migrations -seq create_people_table** 
- update file **.up.sql & .down.sql**
- running : **migrate -path database/migrations -database "mysql://user:password@tcp(localhost:3306)/dbname" up**


