## Introduction

this repository for educational purpose only 

## feature 
* using gin
* using gorm
* using mysql
* have feature auto migrate from sql query
* have feature to use function in sql
  please check the instalation for detail


## Installation

* make sure you edit the env example file or you just can create new database with same name like mine
* go mod tidy 
* please migrate the table and all function using
  - "go run main.go -migrateDatabaseFunction="true" -migrateTable="true" (after you migrate please just use go run main.go, otherwise its will bring error"

