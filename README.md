[![CircleCI](https://circleci.com/gh/victorsteven/Go-JWT-Postgres-Mysql-Restful-API.svg?style=svg)](https://circleci.com/gh/victorsteven/Go-JWT-Postgres-Mysql-Restful-API)

# Go-JWT-Postgres-Mysql-Restful-API
This is an application built with golang, jwt, gorm, postgresql, mysql.

You can follow the guide here:
https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121

### Dockerizing the API
The dockerized API can be found here:
https://github.com/victorsteven/Dockerized-Golang-Postgres-Mysql-API

We will be using third party packages in this application. If you have never installed them before, you can run the following commands:

go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/mysql" //If using mysql 
go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1