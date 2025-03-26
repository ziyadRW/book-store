# Go Book Store

A simple **Go** application demonstrating a CRUD-based bookstore API using:

- **Gorilla Mux** for routing
- **GORM** for database interactions (MySQL driver)
- **godotenv** for environment variable management

---

## Project Overview

This project exposes a set of **RESTful** endpoints to create, read, update, and delete books stored in a MySQL database. It follows a layered architecture, organizing code into **internal** packages such as `handlers`, `models`, `routes`, and `utils`.
