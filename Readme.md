# GOGO

A framework in golang for handaling big monolithic projects By using GoFiber and GORM.

## Features

- Boilerplate Generating
- Migration Handling

## Instalation
First of all download the project
```
git clone https://github.com/Akikzamil/gogo.git
```
add path of the file into the env

## Dev build
```
go build -o gogo.exe
```

## Starting
Run the command:
```
gogo init projectName.
```

There will be a new project created.

copy the `.env.example` file to `.env` file.

set the Database connection with proper values.

## Run Migration
```
gogo migrate
```
## Migration Rollback
```
gogo migrate:rollback
```
## Rollback Migration
```
gogo migrate:rollback
```
## Make Model
```
gogo make model cat
```
## Make Model With migration
```
gogo make model cat --migration
```