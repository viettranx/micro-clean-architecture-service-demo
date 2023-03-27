# Clean Architecture in Microservices Demo

## Introduction

> If you're very new to Clean Architecture, I recommend you check [more simple demo](https://github.com/viettranx/simple-clean-architecture-demo).

This repo was built up with three services: **Task**, **User** and **Auth**. Each of them conforms **Clean Architecture**.

![Clean Architecture Overview](./statics/img/overview-clean-architecture.jpg)
![Clean Architecture Flow](./statics/img/clean-architecture-flow.jpg)

## Folder structure
Let's focus on directories in [`/services`](./services):

```
services
├── auth
│   ├── business
│   ├── entity
│   ├── repository
│   │   ├── mysql
│   │   └── rpc
│   └── transport
│       ├── api
│       └── rpc
├── task
│   ├── business
│   ├── entity
│   ├── repository
│   │   ├── mysql
│   │   └── rpc
│   └── transport
│       └── api
└── user
    ├── business
    ├── entity
    ├── repository
    │   └── mysql
    └── transport
        ├── api
        └── rpc
```

Protobuf and generated files at [proto folder](./proto):

```
├── auth.proto
├── pb
│   ├── auth.pb.go
│   ├── auth_grpc.pb.go
│   ├── user.pb.go
│   └── user_grpc.pb.go
└── user.proto
```

## Microservices in this demo

Instead of separating to 3 repos on GitHub, I merged them to a single repo to give more transparent and convenience to build up the demo.
But the services are isolated with each others.

Service stacks: 
- GIN (for HTTP service)
- gRPC (internal rpc)
- MySQL (with GORM)
- JWT

### Task Service
- REST API: CRUD Task (require authorization for all apis)
- gRPC: None

### User Service
- REST API:
  - Get user profile (require authorization)
- gRPC:
  - Get public user info by ID
  - Get many public user infos by IDs

### Auth Service
- REST API:
  - Register new user account
  - Login (issue JWT)
- gRPC:
  - Introspect Access Token

## The Diagrams
Some of main diagrams to demonstration how they work together:

### Create Task
![Create Task Diagram](./statics/img/create-tasks-diagram.png)

### Fetch Tasks
![Fetch Task Diagram](./statics/img/get-tasks-diagram.png)

## How to run this demo

### 1. Clone and start the service with Docker-Compose
Open your terminal/console:
```shell
git clone https://github.com/viettranx/microservices-clean-architecture-demo
cd microservices-clean-architecture-demo

docker compose up --force-recreate --detach --build app
```

If everything is ok, the service will be running at on `localhost:3000`

Print all ENV variables:
```bash
docker compose exec app ./demo_app outenv
```

### 2. Make some requests
 - Register new account

```shell
curl --location 'http://localhost:3000/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "demo@demo.com",
    "password": "12345678",
    "last_name": "Microservices",
    "first_name": "Demo "
}'
```
```shell
{"data": true}
```

- Login

```shell
curl --location 'http://localhost:3000/v1/authenticate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "demo@demo.com",
    "password": "12345678"
}'
```
The access token will return like this
```shell
{
  "data": {
    "access_token": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJlNTMycW9zOGpqTTIiLCJleHAiOjE2ODAyNTI0MDgsIm5iZiI6MTY3OTY0NzYwOCwiaWF0IjoxNjc5NjQ3NjA4LCJqdGkiOiI3OTEzYzhjYy05NmI0LTQ3ZmUtOWIzZi01MTUwZTk5NTM3MGUifQ.51d6zVuGtcAbw2poEWV4TffhEqJG8uxMOcGq7Mt8sZA",
      "expire_in": 604800
    }
  }
}
```

- Create a new task
```shell
curl --location 'http://localhost:3000/v1/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJlNTMycW9zOGpqTTIiLCJleHAiOjE2ODAyNTI0MDgsIm5iZiI6MTY3OTY0NzYwOCwiaWF0IjoxNjc5NjQ3NjA4LCJqdGkiOiI3OTEzYzhjYy05NmI0LTQ3ZmUtOWIzZi01MTUwZTk5NTM3MGUifQ.51d6zVuGtcAbw2poEWV4TffhEqJG8uxMOcGq7Mt8sZA' \
--header 'Content-Type: application/json' \
--data '{
    "title": "This is a new task",
    "description": "Build a simple demo clean architecture"
}'
```

New Task id will be returned
```shell
{"data":"e532sJ4XpCi8"}
```

- Fetch list tasks
```shell
curl --location 'http://localhost:3000/v1/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJnR3pUQlVScWhhakciLCJleHAiOjE2ODAwNjQzNzcsIm5iZiI6MTY3OTQ1OTU3NywiaWF0IjoxNjc5NDU5NTc3LCJqdGkiOiI5Y2U4ZjUwMC1hYTczLTQ1MTYtYmJhYy04OTBiNDBjNDVkMDAifQ.VsqU8_Gn_EMwKAYW-En3BG58nWTREG6s3zCVXu7sZgQ'
```
```json
{
  "data": [
    {
      "id": "e532sJ4XpCi8",
      "created_at": "2023-03-24T08:50:07Z",
      "updated_at": "2023-03-24T08:50:07Z",
      "title": "This is a new task",
      "description": "Build a simple demo clean architecture",
      "status": "doing",
      "user": {
        "id": "e532qos8jjM2",
        "last_name": "Microservices",
        "first_name": "Demo",
        "avatar": null
      }
    }
  ],
  "paging": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "cursor": "",
    "next_cursor": ""
  },
  "extra": {}
}
```

## Conclusion

Building services with Clean Architecture is difficult, especially in Microservices.
I hope this repo help you. Enjoy it and feel free to create PRs or Issues.
