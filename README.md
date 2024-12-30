
# Simple Inventory Management System

---

##### Author: Vincent E. Dogan Dursun
##### Framework: Beego
##### ORM: Beego ORM
##### Language: Go/Golang
##### Database: Postgres SQL & SQLite

---

## Description

- The framework used for this application is Beego & Beego ORM.
- The database used for this application is Postgres SQL for active object management and SQLite for testing database.
- The application is a simple inventory management system that allows users to add, update, delete, and view products,
    warehouses, users with different roles (admin or user), and stocks.
- The application has a simple authentication system that allows users to sign up, sign in, and sign out.
- The application has a simple authorization system that allows users to access different parts of the application based on 
their roles. (e.g. admin can create, modify or delete items, but user can only view items)

---

## Installation

- Clone the repository
- Install the dependencies
- Run the application

```bash
git clone
cd inventory-management-system
go mod download
go run main.go
```

---
