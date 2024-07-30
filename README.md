# Auth api
 Authentication api using Go/Gin/Docker/PostgreSQL

## Run
 Run with docker
 ```cmd
 cd deployments/docker-compose
 docker-compose up -d
 ```

## Endpoints
 ### POST `/auth/login/`
 Request example
 ```json
 {
     "login": "alexdenkk",
     "password": "12345678",
 }
 ```
 Response example
 ```json
 {
     "token": "<token>"
 }
 ```
