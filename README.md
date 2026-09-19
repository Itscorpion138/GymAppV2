# GymAppV2 — Gym Management REST API

A full-stack gym management application built with **Go, Gin, PostgreSQL, and HTML/CSS/JavaScript**.

The project provides a REST API for managing trainees and coaches, together with a simple web frontend for interacting with the backend.

## Features

* RESTful API built with Go and Gin
* PostgreSQL database integration
* CRUD operations for trainee accounts
* CRUD operations for coach accounts
* UUID-based resource identifiers
* Coach profile image uploads
* Local file storage for uploaded profile pictures
* HTML frontend for interacting with the API
* Environment-based database configuration
* Trainee and Coach roles in the frontend

## Architecture

The application consists of three main parts:

```text
┌─────────────────────┐
│    HTML Frontend    │
└──────────┬──────────┘
           │ HTTP
           ▼
┌─────────────────────┐
│     Go REST API     │
│        Gin          │
└──────────┬──────────┘
           │
      ┌────┴─────┐
      ▼          ▼
┌──────────┐ ┌──────────────────┐
│PostgreSQL│ │  File Storage    │
│ Database │ │ uploads/coaches/ │
└──────────┘ └──────────────────┘
```

The frontend communicates with the Go backend through HTTP requests. The backend stores application data in PostgreSQL and coach profile pictures in the local uploads directory.

## Technologies

| Technology          | Purpose                  |
| ------------------- | ------------------------ |
| Go                  | Backend application      |
| Gin                 | HTTP server and REST API |
| PostgreSQL          | Persistent database      |
| UUID                | Resource identifiers     |
| HTML/CSS/JavaScript | Frontend                 |
| Git/GitHub          | Version control          |

## Data Model

### User Account

The trainee/user account contains:

```text
id
name
lastname
age
height
weight
gender
skilllevel
plan
email
```

### Coach Account

The coach account contains:

```text
id
name
lastname
age
gender
email
cost
verified
profile_pic
bio
```

Coach biographies can contain up to 5000 characters.

## Getting Started

### Prerequisites

Install:

* Go 1.18 or newer
* PostgreSQL
* Git

Verify Go:

```bash
go version
```

Verify that PostgreSQL is running before starting the application.

---

## 1. Clone the Repository

```bash
git clone https://github.com/Itscorpion138/GymAppV2.git
cd GymAppV2/RESTful_GymAppV2
```

---

## 2. Configure the Database

Create a `database.env` file in the backend project directory:

```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gym
```

Replace the values with the credentials of your PostgreSQL installation.

> Never commit real database credentials to GitHub.

A `.env.example` file should be included in the repository so that users can see the required configuration without exposing secrets.

---

## 3. Create the Database

Create the PostgreSQL database:

```sql
CREATE DATABASE gym;
```

Connect to the database and create the required tables.

### User Account

```sql
CREATE TABLE user_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    lastname VARCHAR(100),
    age INT,
    height INT,
    weight FLOAT,
    gender VARCHAR(10),
    skilllevel VARCHAR(20),
    plan VARCHAR(30),
    email VARCHAR(100)
);
```

### Coach Account

```sql
CREATE TABLE coach_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    lastname VARCHAR(100),
    age INT,
    gender VARCHAR(10),
    email VARCHAR(100),
    cost FLOAT,
    verified BOOLEAN DEFAULT FALSE,
    profile_pic VARCHAR(255),
    bio VARCHAR(5000)
);
```

> Database migrations are planned as a future improvement so the database can be initialized without manually executing SQL.

---

## 4. Install Dependencies

From the backend directory:

```bash
go mod tidy
```

This downloads the project's Go dependencies.

---

## 5. Create the Upload Directory

Coach profile pictures are stored locally.

Create the directory:

```bash
mkdir -p uploads/coaches
```

On Windows, create the same directory manually if necessary:

```text
uploads/
└── coaches/
```

Uploaded images use coach IDs as part of their filenames.

---

## 6. Run the Backend

Start the Go server:

```bash
go run main.go
```

By default, the API runs at:

```text
http://localhost:8080
```

---

## 7. Run the Frontend

Open the `html` directory and launch:

```text
html/index.html
```

The frontend allows you to switch between:

* Trainee
* Coach

and interact with the backend to create, edit, and delete users and coaches.

Coach profile pictures can also be uploaded through the application.

---

## Database Reset

To reset the application database, remove and recreate the relevant tables.

For example:

```sql
DROP TABLE IF EXISTS coach_account;
DROP TABLE IF EXISTS user_account;
```

Then recreate the tables using the SQL definitions above.

> Database migrations will make this process easier in a future version.

---

## Project Structure

The current project contains the REST API and frontend in the same repository.

```text
GymAppV2/
│
├── RESTful_GymAppV2/
│   ├── main.go
│   ├── go.mod
│   ├── database.env
│   ├── uploads/
│   │   └── coaches/
│   └── ...
│
└── html/
    └── ...
```

The backend structure is planned to be further modularized as the project is refactored.

---

## Future Improvements

The current version provides the core gym management functionality. Planned improvements include:

* [x] Database migrations
* [ ] Request validation
* [ ] Consistent API error responses
* [ ] Authentication
* [ ] Authorization
* [ ] Password hashing
* [ ] Improved project architecture
* [ ] Unit tests
* [ ] Integration tests
* [ ] Docker support
* [ ] OpenAPI / Swagger documentation
* [ ] Pagination and filtering
* [ ] Improved database constraints
* [ ] Improved file-upload validation
* [ ] `.env.example` and secure configuration

## What This Project Demonstrates

This project was built to practice backend development and client-server communication.

It demonstrates experience with:

* Designing and implementing REST APIs
* Developing HTTP services with Go
* Working with PostgreSQL
* Performing CRUD operations
* Handling UUID-based resources
* Connecting a frontend application to a backend API
* Handling file uploads
* Managing persistent application data
* Structuring a small full-stack application

## Project Status

**Active development/refactoring**

The initial version established the REST API, PostgreSQL integration, frontend communication, and coach image upload functionality.

The next stage is focused on improving the backend architecture, validation, testing, database management, security, and deployment workflow.

## Author

**Erfan Modirian**

Junior Backend Developer
Mathematics undergraduate at Ferdowsi University of Mashhad

GitHub: https://github.com/Itscorpion138# GymAppV2 — Gym Management REST API

A full-stack gym management application built with **Go, Gin, PostgreSQL, and HTML/CSS/JavaScript**.

The project provides a REST API for managing trainees and coaches, together with a simple web frontend for interacting with the backend.

## Features

* RESTful API built with Go and Gin
* PostgreSQL database integration
* CRUD operations for trainee accounts
* CRUD operations for coach accounts
* UUID-based resource identifiers
* Coach profile image uploads
* Local file storage for uploaded profile pictures
* HTML frontend for interacting with the API
* Environment-based database configuration
* Trainee and Coach roles in the frontend

## Architecture

The application consists of three main parts:

```text
┌─────────────────────┐
│    HTML Frontend    │
└──────────┬──────────┘
           │ HTTP
           ▼
┌─────────────────────┐
│     Go REST API     │
│        Gin          │
└──────────┬──────────┘
           │
      ┌────┴─────┐
      ▼          ▼
┌──────────┐ ┌──────────────────┐
│PostgreSQL│ │  File Storage    │
│ Database │ │ uploads/coaches/ │
└──────────┘ └──────────────────┘
```

The frontend communicates with the Go backend through HTTP requests. The backend stores application data in PostgreSQL and coach profile pictures in the local uploads directory.

## Technologies

| Technology          | Purpose                  |
| ------------------- | ------------------------ |
| Go                  | Backend application      |
| Gin                 | HTTP server and REST API |
| PostgreSQL          | Persistent database      |
| UUID                | Resource identifiers     |
| HTML/CSS/JavaScript | Frontend                 |
| Git/GitHub          | Version control          |

## Data Model

### User Account

The trainee/user account contains:

```text
id
name
lastname
age
height
weight
gender
skilllevel
plan
email
```

### Coach Account

The coach account contains:

```text
id
name
lastname
age
gender
email
cost
verified
profile_pic
bio
```

Coach biographies can contain up to 5000 characters.

## Getting Started

### Prerequisites

Install:

* Go 1.18 or newer
* PostgreSQL
* Git

Verify Go:

```bash
go version
```

Verify that PostgreSQL is running before starting the application.

---

## 1. Clone the Repository

```bash
git clone https://github.com/Itscorpion138/GymAppV2.git
cd GymAppV2/RESTful_GymAppV2
```

---

## 2. Configure the Database

Create a `database.env` file in the backend project directory:

```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gym
```

Replace the values with the credentials of your PostgreSQL installation.

> Never commit real database credentials to GitHub.

A `.env.example` file should be included in the repository so that users can see the required configuration without exposing secrets.

---

## 3. Create the Database

Create the PostgreSQL database:

```sql
CREATE DATABASE gym;
```

Connect to the database and create the required tables.

### User Account

```sql
CREATE TABLE user_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    lastname VARCHAR(100),
    age INT,
    height INT,
    weight FLOAT,
    gender VARCHAR(10),
    skilllevel VARCHAR(20),
    plan VARCHAR(30),
    email VARCHAR(100)
);
```

### Coach Account

```sql
CREATE TABLE coach_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    lastname VARCHAR(100),
    age INT,
    gender VARCHAR(10),
    email VARCHAR(100),
    cost FLOAT,
    verified BOOLEAN DEFAULT FALSE,
    profile_pic VARCHAR(255),
    bio VARCHAR(5000)
);
```

> Database migrations are planned as a future improvement so the database can be initialized without manually executing SQL.

---

## 4. Install Dependencies

From the backend directory:

```bash
go mod tidy
```

This downloads the project's Go dependencies.

---

## 5. Create the Upload Directory

Coach profile pictures are stored locally.

Create the directory:

```bash
mkdir -p uploads/coaches
```

On Windows, create the same directory manually if necessary:

```text
uploads/
└── coaches/
```

Uploaded images use coach IDs as part of their filenames.

---

## 6. Run the Backend

Start the Go server:

```bash
go run main.go
```

By default, the API runs at:

```text
http://localhost:8080
```

---

## 7. Run the Frontend

Open the `html` directory and launch:

```text
html/index.html
```

The frontend allows you to switch between:

* Trainee
* Coach

and interact with the backend to create, edit, and delete users and coaches.

Coach profile pictures can also be uploaded through the application.

---

## Database Reset

To reset the application database, remove and recreate the relevant tables.

For example:

```sql
DROP TABLE IF EXISTS coach_account;
DROP TABLE IF EXISTS user_account;
```

Then recreate the tables using the SQL definitions above.

> Database migrations will make this process easier in a future version.

---

## Project Structure

The current project contains the REST API and frontend in the same repository.

```text
GymAppV2/
│
├── RESTful_GymAppV2/
│   ├── main.go
│   ├── go.mod
│   ├── database.env
│   ├── uploads/
│   │   └── coaches/
│   └── ...
│
└── html/
    └── ...
```

The backend structure is planned to be further modularized as the project is refactored.

---

## Future Improvements

The current version provides the core gym management functionality. Planned improvements include:

* [x] Database migrations
* [ ] Request validation
* [ ] Consistent API error responses
* [ ] Authentication
* [ ] Authorization
* [ ] Password hashing
* [ ] Improved project architecture
* [ ] Unit tests
* [ ] Integration tests
* [ ] Docker support
* [ ] OpenAPI / Swagger documentation
* [ ] Pagination and filtering
* [ ] Improved database constraints
* [ ] Improved file-upload validation
* [ ] `.env.example` and secure configuration

## What This Project Demonstrates

This project was built to practice backend development and client-server communication.

It demonstrates experience with:

* Designing and implementing REST APIs
* Developing HTTP services with Go
* Working with PostgreSQL
* Performing CRUD operations
* Handling UUID-based resources
* Connecting a frontend application to a backend API
* Handling file uploads
* Managing persistent application data
* Structuring a small full-stack application

## Project Status

**Active development/refactoring**

The initial version established the REST API, PostgreSQL integration, frontend communication, and coach image upload functionality.

The next stage is focused on improving the backend architecture, validation, testing, database management, security, and deployment workflow.

## Author

**Erfan Modirian**

Junior Backend Developer
Mathematics undergraduate at Ferdowsi University of Mashhad

GitHub: https://github.com/Itscorpion138
