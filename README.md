# Eisenhower Matrix Todo API

## Information
Eisehower matrix is a time management tool that helps prioritize tasks based on their urgency and importance. This API implements the eisehower matrix to help manage tasks.

## Tech Stack
- Golang (Main Language)
- Httprouter (Router)
- Database: Postgres

## API Endpoints
- GET ```/api/tasks```
Get all todos

Response Body 200
```json
{
    "code": 0,
    "status": "string",
    "data": [
        {
            "id": 0,
            "title": "string",
            "description": "string",
            "type": "string",
            "created_at": "string",
            "updated_at": "string"
        }
    ]
}
```

- GET ```/api/tasks/:id```
Get todos by id

Response Body 200
```json
{
    "code": 0,
    "status": "string",
    "data": {
        "id": 0,
        "title": "string",
        "description": "string",
        "type": "string",
        "created_at": "string",
        "updated_at": "string"
    }
}
```

- POST ```/api/tasks```
Create a new todo

Request Body
```json
{
    "title": "string",
    "description": "string",
    "type": "string"
}
```

Response Body 200
```json
{
    "code": 0,
    "status": "string",
    "data": {
        "id": 0,
        "title": "string",
        "description": "string",
        "type": "string",
        "created_at": "string",
        "updated_at": "string"
    }
}
```

- PATCH ```/api/tasks/:id```
Update todo by id but not all fields

Request Body
```json
{
    "title": "string",
    "description": "string",
    "type": "string"
}
```

Response Body 200
```json
{
    "code": 0,
    "status": "string",
    "data": {
        "id": 0,
        "title": "string",
        "description": "string",
        "type": "string",
        "created_at": "string",
        "updated_at": "string"
    }
}
```


- DELETE ```/api/tasks/:id```
Delete todo by id

Respoonse Body
```json
{
    "code": 0,
    "status": "string"
}
```