# Go-Fiber-Mongo-HRMS

A simple Human Resource Management System (HRMS) API built using **Go**, **Fiber**, and **MongoDB**. This project provides a RESTful API to perform CRUD operations on employee data.

## Features
- Create, Read, Update, and Delete (CRUD) employee records.
- Uses Fiber for lightweight and fast web server handling.
- Uses MongoDB as the database to store employee information.
- Implements structured error handling.
- Provides JSON-based API responses.

## Technologies Used
- **Golang** (Go) - Backend language
- **Fiber** - Express-inspired web framework for Go
- **MongoDB** - NoSQL database for storing employee records
- **Go Mongo Driver** - MongoDB driver for Go

## Installation and Setup
### Prerequisites
Ensure you have the following installed on your system:
- [Go](https://golang.org/doc/install)
- [MongoDB](https://www.mongodb.com/try/download/community)

### Steps to Run the Project
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/Go-Fiber-Mongo-HRMS.git
   cd Go-Fiber-Mongo-HRMS
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Start MongoDB (if not already running):
   ```sh
   mongod --dbpath=/path/to/your/db
   ```

4. Run the application:
   ```sh
   go run main.go
   ```

5. The server should now be running on `http://localhost:3000`

## API Endpoints
### 1. Get All Employees
- **Endpoint:** `GET /employee`
- **Response:** Returns a list of all employees

### 2. Create a New Employee
- **Endpoint:** `POST /employee`
- **Request Body:** (JSON)
  ```json
  {
    "name": "John Doe",
    "salary": 50000,
    "age": 30
  }
  ```
- **Response:** Returns the newly created employee record

### 3. Update an Employee
- **Endpoint:** `PUT /employee/:id`
- **Request Body:** (JSON)
  ```json
  {
    "name": "John Doe Updated",
    "salary": 55000,
    "age": 31
  }
  ```
- **Response:** Returns the updated employee record

### 4. Delete an Employee
- **Endpoint:** `DELETE /employee/:id`
- **Response:** Returns a success message if deletion is successful

## Project Structure
```
Go-Fiber-Mongo-HRMS/
│-- main.go            # Main application file
│-- go.mod             # Go module file
│-- go.sum             # Dependency file
│-- README.md          # Project documentation
```

## Error Handling
- **400 Bad Request:** When invalid input is provided.
- **404 Not Found:** When an employee with a given ID is not found.
- **500 Internal Server Error:** When there is a server-side issue.

## Contribution
Feel free to contribute! Fork the repo, create a new branch, and submit a pull request.



