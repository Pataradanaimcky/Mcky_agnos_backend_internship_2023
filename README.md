# Strong Password Recommendation API

This project provides a backend server that exposes an API for recommending steps to make a password strong. The API takes a password as input and calculates the minimum number of actions required to make it strong based on specific criteria.

## API Endpoint

The API endpoint for recommending steps to make a password strong is:

**POST** `/api/strong_password_steps`

### Request Payload

The API expects a JSON payload in the following format:

```json
{
  "init_password": "aA1"
}
```
Response Payload
The API responds with a JSON payload containing the minimum number of steps required to make the password strong:

```json
{
  "num_of_steps": 3
}
```

## Criteria for a Strong Password

A password is considered strong if it meets the following criteria:

Password length is between 6 and 20 characters (inclusive).
Contains at least 1 lowercase letter, 1 uppercase letter, and 1 digit.
Does not contain 3 repeating characters in a row (e.g., "11123").

## Tech Stack

The backend server is built using Go and the Gin web framework. It uses PostgreSQL as the database to store log information.

## Docker Compose

The project includes a docker-compose.yml file that defines the database service using PostgreSQL. To set up the database, run the following command:

```
docker-compose up -d
```

## Environment Variables

The project uses environment variables to configure the database connection. Create a .env file in the project root directory with the following content:

```
DATABASE_URL=postgres://your_postgres_username:your_postgres_password@localhost:5432/your_database_name?sslmode=disable
```
Replace your_postgres_username, your_postgres_password, and your_database_name with your PostgreSQL credentials and database name.

## Database Initialization

The main.go file initializes the PostgreSQL database, creates the required logs table if it does not exist, and starts the server.

## Middleware

The project includes a request logger middleware that logs incoming requests and their processing time. The logged information is stored in the PostgreSQL database.

## Unit Tests

The password_util_test.go file contains unit tests for the utils.CountActionsToMakePasswordStrong function. To run the tests, use the following command:

```
go test ./tests
```

## How to Run

To run the project, execute the following command:

```
go run main.go
```
The server will start listening on port 8080.

## Bonus Features

The project includes unit tests to ensure the correctness of the password strength calculation logic.
HTTP status codes are used to indicate the success or failure of API requests.
The project follows a clear code structure for easy maintenance and readability.
Conclusion

This project provides a backend API for determining the minimum number of actions required to make a password strong based on specific criteria. It uses Go and the Gin web framework for building the server, and PostgreSQL as the database for logging request information. The project includes unit tests and a request logger middleware to enhance the reliability and maintainability of the codebase.
