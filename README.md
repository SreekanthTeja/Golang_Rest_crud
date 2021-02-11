# Golang Rest Api Crud with Mysql
Overview:
---
    The project deals with rest api of crud with mysql database using Golang.

Running the project:
---
```json
$ git clone https://github.com/SreekanthTeja/Golang_Rest_crud.git
Go to project root
$ cd <project name>

Install dependicy libraries for the project
$ go get "github.com/go-sql-driver/mysql"
$ go get "github.com/gorilla/mux" 
Run the server
$ go run g1.go 
Testing the project
1.localhost:8080/
2.localhost:8080/p/1
3.localhost:8080/new
4.localhost:8080/update/2["can give 1,3 ..as per availability"]
5.localhost:8080/delete/3
```