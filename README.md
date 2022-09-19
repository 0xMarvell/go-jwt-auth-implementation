# Go-JWT-Auth-Implementation

Implementation of a simple authentication system using [JWT (JSON Web Tokens)](https://jwt.io/). The API generates a token (which will be used for authentication) for the user. The token is then stored in an [httpOnly cookie](https://www.cookiepro.com/knowledge/httponly-cookie/) for security purposes, as opposed to storing the token in local storage.

The API was built using the following tools:

- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [The Golang-JWT package](https://github.com/golang-jwt/jwt)
- [The Bcrypt package](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## DOCS

Read the documentation [here](https://documenter.getpostman.com/view/15381378/2s7Ymrkmuz)

## Environment Variables

This project makes use of the [godotenv](github.com/joho/godotenv) package to store environment variables. To run this project on your local machine, you will need to add your Postgres database details as environment variables. This [article](https://dev.to/schadokar/use-environment-variable-in-your-next-golang-project-2o6c) explains how to use the package.
> An `.env.example` file has been provided to help with setting up your environment variables.

## Run API locally

- Clone Repo

    ```bash
    $ git clone https://github.com/0xMarvell/go-jwt-auth-implementation.git
    ```

- Make sure to have [Go](https://go.dev/) installed on your local machine
- Open the code base directory in terminal
- Launch API server:

    ```go
    $ go build -o auth cmd/web/main.go
    $ ./auth
    ```

> Test API using Postman, Insomnia or any other API testing client of your choice
