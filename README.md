# RestAPI Tasks

## Description
A simple REST API for managing tasks.
Each task contains an id, title, and completion status.

## Endpoints
- `POST /tasks` — create a task
- `GET /tasks/{id}` — get a task by id
- `DELETE /tasks/{id}` - delete a task by id
- `PUT /tasks/{id}` - update a task by id

## How to run
Change directory to cmd and run:
cd cmd
go run main.go

## Example
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"id": 1, "title": "test", "done": false}'