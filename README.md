# Appliance Management Backend

This is the backend part of the Appliance Management application. It provides a RESTful API for managing and searching appliances.

## Table of Contents

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Technologies Used](#technologies-used)

## Getting Started

To get started with the backend of the Appliance Management application, follow the instructions below.

## Prerequisites

- Go: Make sure you have Go installed on your system.

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/ashimk/inventory-backend.git
   ```

2. Navigate to the project directory:

   ```bash
   cd inventory-backend
   ```

3. Install the required dependencies:

   ```bash
   go mod download
   ```

## Configuration

Before running the backend, make sure to set the necessary environment variables. Create a `.env` file in the project root directory and set the following variables:

```
PORT=8080
```

## API Endpoints

The backend provides the following API endpoints for managing appliances:

- `GET /appliances`: Get a list of all recorded appliances.
- `POST /appliance`: Add a new appliance.
- `PUT /appliance`: Update an existing appliance.
- `DELETE /appliance`: Delete an appliance based on the provided query parameters.

Please refer to the source code for more details on the request and response formats for each endpoint.

## Usage

1. Start the development server:

   ```bash
   go run main.go
   ```

   This will start the backend server and the API will be accessible at `http://localhost:8080`.

2. Use the API endpoints to perform CRUD operations on the appliances.

## Technologies Used

- Go: Backend programming language for building the API.
- Gorilla Mux: Router and dispatcher library for building RESTful APIs in Go.
