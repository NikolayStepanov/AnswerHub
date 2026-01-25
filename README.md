# AnswerHub
## Q&A API Service

A simple REST API service for managing questions and answers.

## Description

This service provides an API to:
- create and retrieve questions;
- add answers to questions;
- retrieve a question with all its answers;
- delete questions and answers;
- automatically delete all answers when a question is removed.

## Data Models

### Question
| Field       | Type      | Description |
|------------|-----------|-------------|
| id         | int       | Question ID |
| text       | string    | Question text |
| created_at | datetime  | Creation timestamp |

### Answer
| Field        | Type      | Description |
|-------------|-----------|-------------|
| id          | int       | Answer ID |
| question_id | int       | Related question ID |
| user_id     | string    | User identifier (e.g. UUID) |
| text        | string    | Answer text |
| created_at  | datetime  | Creation timestamp |

##  API Endpoints

### Questions

- **Get all questions** — `GET /questions/`
- **Create a new question** — `POST /questions/`
- **Get a question with answers** — `GET /questions/{id}`
- **Delete a question (with answers)** — `DELETE /questions/{id}`

### Answers

- **Add an answer to a question** — `POST /questions/{id}/answers/`
- **Get an answer by ID** — `GET /answers/{id}`
- **Delete an answer** — `DELETE /answers/{id}`

## Business Logic

- An answer cannot be created for a non-existing question
- A user can submit multiple answers to the same question
- When a question is deleted, all related answers are deleted automatically (cascade)
- Timestamps are generated automatically

## Technologies

- Go
- net/http
- PostgreSQL
- GORM
- goose (database migrations)
- Docker
- Docker Compose

## Running the Application
The application is managed using `make` and `docker-compose`.

### Build and run everything
Build the Go service and start all containers:
- `make run-all`

### Start containers only
Start services in detached mode:
- `make compose-up`

### Stop containers
Stop all running containers:
- `make compose-down`

### Remove containers
Stop and remove all containers:
- `make compose-rm`

### Restart containers
Fully restart the environment:
- `make compose-rs`

After startup, the API will be available once all containers are running.
