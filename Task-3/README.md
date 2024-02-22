# Jokes API

The Jokes API allows you to create, retrieve, update, and delete jokes through a RESTful interface.

## Getting Started

These instructions will get you through the bootstrap phase of creating and deploying instances of jokes resources.

### Prerequisites

- The Go web server running locally on port `8080`.
- `curl` installed on your machine to make HTTP requests from the command line.

### API Endpoints and `curl` Commands

#### Create a New Joke

To add a new joke:

```bash
curl -X POST http://localhost:8080/jokes \
     -H 'Content-Type: application/json' \
     -d '{ 
           "setup": "Why dont scientists trust atoms?",
           "punchline": "Because they make up everything"
         }'
```
```bash
curl -X POST http://localhost:8080/jokes \
     -H 'Content-Type: application/json' \
     -d '{ 
           "setup": "Why was the math book sad?",
           "punchline": "Because it had too many problems"
         }'
```
```bash
curl -X POST http://localhost:8080/jokes \
     -H 'Content-Type: application/json' \
     -d '{ 
           "setup": "Why did the developer break up with Kubernetes?",
           "punchline": "Because they needed more space; Kubernetes was always hogging the cluster"
         }'
```

#### Get All Jokes

To retrieve all jokes:

```bash
curl -X GET http://localhost:8080/jokes
```

#### Get a Single Joke by ID

To retrieve a joke by its ID (replace `1` with the actual ID):

```bash
curl -X GET http://localhost:8080/jokes/1
```

#### Update a Joke

To update an existing joke by ID (replace `1` with the actual ID):

```bash
curl -X PATCH http://localhost:8080/jokes/3 \
     -H 'Content-Type: application/json' \
     -d '{
           "setup": "Updated setup",
           "punchline": "Updated punchline"
         }'
```

#### Delete a Joke by ID

To delete a joke by its ID (replace `1` with the actual ID):

```bash
curl -X DELETE http://localhost:8080/jokes/1
```