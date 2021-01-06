# Tasker

<img align="right" width="50%" src="./images/gopher.png">

Tasker is a simple Trello-like project management service.

## Description

Back-end (REST API) part of the task management application (like Trello) with no authentication and authorization required.

The main entity is a Project that always has its name and contains multiple Columns.

A Column has name, contains Tasks and represents their status.
When a Project created, “default” Column is created also. Columns can be moved left or right.

A Task can be created only inside the Column and can be moved within the Column (change priority) or across the Columns (change status).

A Task can have Comments that could contain questions or Task clarification information.

API docs is Postman collection in `api` folder.

Deployed version: <http://167.99.253.9:8080/api/v1>

## Run instructions

1) Create `.env` file for server configuration. For example:
```bash
SERVER_ADDR=:8080
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=123
POSTGRES_DBNAME=tasker
POSTGRES_SSLMODE=disable
```

2) Spin up `postgres` container.
```bash
$ docker-compose up postgres
```

3) Create database.
```bash
$ docker exec -it postgres bash
$ psql -U postgres
$ CREATE DATABASE tasker;
```

4) Spin up `tasker` container.
```bash
$ docker-compose up tasker
```
