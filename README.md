# 🚀 Golang + React CRUD App with PostgreSQL

## 📌 Overview
This project is a **full-stack CRUD application** built using:
- **Backend:** Golang (Gin framework) with PostgreSQL
- **Frontend:** React (Vite + Material UI)
- **Database:** PostgreSQL

This application allows users to:
- **View** all users (GET `/users`)
- **Add** a new user (POST `/users`)
- **Update** user details (PUT `/users/:id`)

## 🛠️ Tech Stack
### Backend:
- **Golang**
- **Gin Framework**
- **PostgreSQL**
- **Gin-CORS Middleware**

### Frontend:
- **React** (Vite)
- **Axios** (for API requests)
- **Material UI** (for UI components)

## 📂 Project Structure
```
crud_golang/
│── server/               # Backend (Golang)
│   ├── main.go           # Entry point for the Go server
│   ├── db.go             # Database connection
│   ├── handlers.go       # API route handlers
│   ├── models.go         # User model
│   ├── go.mod            # Go module dependencies
│
│── front-end/            # Frontend (React)
│   ├── src/
│   │   ├── App.jsx       # Main React component
│   │   ├── api.js        # API calls using Axios
│   │   ├── index.css     # Styles
│   │   ├── main.jsx      # React app entry point
│   ├── package.json      # React dependencies
│   ├── vite.config.js    # Vite server configuration
```

## ⚙️ Setup & Installation
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/ahmed-elmarrouni/crud-golang-react.git
cd crud-golang-react
```

### 2️⃣ Setup the Backend (Golang + PostgreSQL)
#### Install Go Dependencies
```sh
cd server
go mod tidy
```
#### Configure Database
1. Create a PostgreSQL database:
```sql
CREATE DATABASE crud_golang_db;
```
2. Create a `users` table:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Run the Backend Server
```sh
go run .
```
The backend will start at: **`http://localhost:8080`**

### 3️⃣ Setup the Frontend (React + Vite)
```sh
cd ../front-end
npm install
npm run dev
```
The frontend will be available at: **`http://localhost:5173`**

## 📌 API Endpoints
**GET** `/users`

**POST** `/users`

