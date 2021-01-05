# Tasker

<img align="right" width="50%" src="./images/gopher.png">

Tasker is a simple Trello-like project management service.

## Description

Back-end (REST API) part of the task management application (like Trello) with no authentication and authorization required.

The main entity is a Project (or Board) that always has its name and contains multiple Columns.

A Column has name, contains Tasks and represents their status.
When a Project created, “default” Column is created also. Columns can be moved left or right.

A Task can be created only inside the Column and can be moved within the Column (change priority) or across the Columns (change status).

A Task can have Comments that could contain questions or Task clarification information.

API docs is Postman collection in `api` folder.

## Run instructions

1) Spin up `postgres` container.
```bash
$ docker-compose up postgres
```

2) Create database.
```bash
$ docker-compose exec -it postgres bash
$ psql -U postgres
$ CREATE DATABASE tasker;
```

3) Spin up `tasker` container.
```bash
$ docker-compose up tasker
```
